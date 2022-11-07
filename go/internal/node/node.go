package node

import (
	"math"
	"math/rand"
	"time"

	"github.com/kumarabd/CS6378-Project2/go/config"
	application_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/application"
	"github.com/kumarabd/CS6378-Project2/go/pkg/lamport"
	mutex_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/mutex"
	"github.com/realnighthawk/bucky/logger"
)

type Node struct {
	log             logger.Handler
	id              string
	delay           float64
	eTime           float64
	numReq          int
	neighbours      map[string]*application_pkg.Neighbour
	application     application_pkg.Application
	mutex           *mutex_pkg.Mutex
	Channel         *Channel
	avgResponseTime float64
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
	delay := rand.ExpFloat64() * float64(cfg.IR)
	csTime := rand.ExpFloat64() * float64(cfg.CT)

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
	}

	node.delay = math.Floor(time.Duration(node.delay*1000000000).Seconds()*1000000) / 1000000

	return &node, nil
}

func (n *Node) Start() error {
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

	// Start requests
	n.log.Info("delay: ", n.delay)
	startClock := time.Now()
	prevClock := time.Now()
	for n.numReq > 0 {
		diff := math.Floor(time.Since(prevClock).Seconds()*1000000) / 1000000
		if diff == n.delay {
			req_clock := math.Floor(time.Since(startClock).Seconds()*1000000) / 1000000
			n.log.Info("requesting cs at ", req_clock)
			n.application.CS_Enter(req_clock)

			exec_clock := math.Floor(time.Since(startClock).Seconds()*1000000) / 1000000
			n.log.Info("executing cs at ", exec_clock)
			n.mutex.Execute_CS()

			done_clock := math.Floor(time.Since(startClock).Seconds()*1000000) / 1000000
			n.log.Info("leaving cs at ", done_clock)
			n.application.CS_Leave()

			// Calculate response time
			n.avgResponseTime = (n.avgResponseTime + (done_clock - req_clock)) / 2

			n.numReq--
			prevClock = time.Now()
		}
	}

	//// Send exit to neighbours
	curr_clock := math.Floor(time.Since(startClock).Seconds()*1000000) / 1000000
	n.log.Info("finished at ", curr_clock)
	n.log.Info("average response time:  ", math.Floor(time.Duration(n.avgResponseTime*1000000000).Seconds()*1000000)/1000000)
	<-stopChan
	return nil
}

func (n *Node) listen(ch chan struct{}) {
	for {
		connection, err := n.Channel.Listen()
		if err != nil {
			n.log.Error(err)
			continue
		}
		go n.application.ProcessMessage(connection, n.numReq, ch)
	}
}

func (n *Node) Get_ID() string {
	return n.id
}
