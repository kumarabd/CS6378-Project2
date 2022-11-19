package node

import (
	"encoding/json"
	"math/rand"
	"net"
	"regexp"
	"time"

	"github.com/kumarabd/CS6378-Project2/go/config"
	"github.com/kumarabd/CS6378-Project2/go/logger"
	application_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/application"
	"github.com/kumarabd/CS6378-Project2/go/pkg/lamport"
	mutex_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/mutex"
)

type Node struct {
	log             logger.Handler
	id              string
	delay           int64
	eTime           int64
	numReq          int
	neighbours      map[string]*application_pkg.Neighbour
	application     application_pkg.Application
	mutex           *mutex_pkg.Mutex
	Channel         *Channel
	avgResponseTime float64
	syncDelay       float64
	vectorTimeStamp []int64
}

func New(id string, cfg config.Config, log logger.Handler) (*Node, error) {
	// Create a server
	channel, err := NewChannel(cfg.Address[id].Host, cfg.Address[id].Port)
	if err != nil {
		return nil, err
	}

	neighbours := make(map[string]*application_pkg.Neighbour, 0)
	for idx, mem := range cfg.Address {
		if idx != id {
			neighbours[idx] = &application_pkg.Neighbour{
				ID: idx,
				HostPort: &config.HostPort{
					Host: mem.Host,
					Port: mem.Port,
				},
			}
		}
	}

	app, err := lamport.New(id, neighbours, log)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	delay := time.Duration(int64(rand.ExpFloat64()*float64(cfg.IR)) * 1000000).Milliseconds()
	csTime := time.Duration(int64(rand.ExpFloat64()*float64(cfg.CT)) * 1000000 + 1).Milliseconds()

	mutex, err := mutex_pkg.New(id, csTime)
	if err != nil {
		return nil, err
	}

	node := Node{
		log:             log,
		id:              id,
		delay:           delay,
		eTime:           csTime,
		numReq:          cfg.R,
		application:     app,
		mutex:           mutex,
		Channel:         channel,
		neighbours:      neighbours,
		avgResponseTime: 0.0,
		vectorTimeStamp: make([]int64, 0),
	}

	return &node, nil
}

func (n *Node) Start() error {
	// Init parameters
	n.log.Info("delay: ", n.delay)
	n.log.Info("execution time: ", n.eTime)
	startClock := time.Now()
	n.application.SetClock(startClock)

	// Start server
	stopChan := make(chan struct{})
	go n.listen(stopChan)

	// Check if all neighbours connected
	connected_list := 0
	for _, mem := range n.neighbours {
		connected := false
		for !connected {
			connection, err := ConnectHost(mem.HostPort.Host, mem.HostPort.Port)
			if err != nil {
				//retry
				continue
			}
			mem.Connection = connection
			connected = true
		}
		connected_list++
	}
	n.log.Info("connected")

	// Start requests
	nr := n.numReq
	prevClock := time.Now()
	startClock = time.Now()
	n.application.SetClock(startClock)
	n.log.WithField("clock", time.Since(startClock).Milliseconds()).Info("started")
	for n.numReq > 0 {
		diff := time.Since(prevClock).Milliseconds()
		if diff >= n.delay {
			req_clock := time.Since(startClock).Milliseconds()
			n.log.WithField("clock", req_clock).Info("requesting cs")
			n.application.CS_Enter()
			n.syncDelay = n.syncDelay - float64(req_clock)

			exec_clock := time.Since(startClock).Milliseconds()
			// n.log.WithField("clock", time.Since(startClock).Milliseconds()).Info("executing cs")
			// n.log.WithField("clock", l.scalarClock).Info("executing cs")
			n.vectorTimeStamp = append(n.vectorTimeStamp, exec_clock)
			n.mutex.Execute_CS()

			done_clock := time.Since(startClock).Milliseconds()
			// n.log.WithField("clock", time.Since(startClock).Milliseconds()).Info("leaving cs")
			// n.log.WithField("clock", l.scalarClock).Info("leaving cs")
			n.application.CS_Leave()

			// Calculate response time
			n.avgResponseTime = n.avgResponseTime + float64(done_clock-req_clock)
			n.syncDelay = n.syncDelay + float64(done_clock)

			n.numReq--
			n.log.WithField("remaining cs requests", n.numReq).Info("completed a recent cs")
			prevClock = time.Now()
		}
	}

	curr_clock := time.Since(startClock).Milliseconds()
	n.log.WithField("clock", curr_clock).Info("finished")
	n.log.WithField("clock", curr_clock).Info("average response time:  ", n.avgResponseTime/float64(nr))
	n.log.WithField("clock", curr_clock).Info("sync delay:  ", n.syncDelay/float64(nr))
	n.log.WithField("clock", curr_clock).Info("throughput:  ", 1/((n.syncDelay/float64(nr))+float64(n.eTime)))
	n.log.WithField("clock", curr_clock).Info("vector time:  ", n.vectorTimeStamp)
	<-stopChan
	return nil
}

func (n *Node) listen(ch chan struct{}) {
	nrequest := n.numReq
	for {
		connection, err := n.Channel.Listen()
		if err != nil {
			n.log.Error(err)
			continue
		}
		go func(conn net.Conn, nreq int) {
			for {
				buffer := make([]byte, 2048)
				mLen, err := conn.Read(buffer)
				if err != nil {
					n.log.Error(err)
					return
				}

				go func(d string, nr int) {
					re := regexp.MustCompile(`\{.*?\}`)
					segs2 := re.FindAllStringSubmatch(d, -1)
					msgs := make([]*application_pkg.Message, 0)
					for _, el := range segs2[0] {
						msg := application_pkg.Message{}
						if err = json.Unmarshal([]byte(el), &msg); err != nil {
							n.log.Error(err)
							continue
						}
						msgs = append(msgs, &msg)
						go n.application.ProcessMessage(msgs, nr, ch)
					}
				}(string(buffer[:mLen]), nreq)
			}
		}(connection, nrequest)
	}
}

func (n *Node) Get_ID() string {
	return n.id
}

