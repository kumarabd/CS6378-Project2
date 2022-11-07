package lamport

import (
	"container/heap"
	"encoding/json"
	"math"
	"net"
	"time"

	application_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/application"
	"github.com/realnighthawk/bucky/logger"
)

type Lamport struct {
	id              string
	neighbours      map[string]*application_pkg.Neighbour
	ReplyPending    map[string]struct{}
	priorityQueue   PriorityQueue
	enterCS         chan struct{}
	inCS            bool
	log             logger.Handler
	releaseReceived int
}

func New(id string, neighbours map[string]*application_pkg.Neighbour, log logger.Handler) (*Lamport, error) {
	lamport := Lamport{
		id:              id,
		neighbours:      neighbours,
		log:             log,
		enterCS:         make(chan struct{}),
		releaseReceived: 0,
	}

	lamport.priorityQueue = make(PriorityQueue, 0)
	heap.Init(&lamport.priorityQueue)
	lamport.ReplyPending = make(map[string]struct{})
	lamport.inCS = false

	return &lamport, nil
}

func (l *Lamport) send(target string, msg application_pkg.Message) error {
	conn := l.neighbours[target].Connection
	_, err := conn.Write(msg.Encode())
	if err != nil {
		return err
	}
	return nil
}

func (l *Lamport) CS_Enter(time float64) {
	msg := application_pkg.Message{
		Source: l.id,
		Type:   application_pkg.REQUEST,
		Time:   time,
	}

	for _, target := range l.neighbours {
		// Send targets
		err := l.send(target.ID, msg)
		if err != nil {
			l.log.Error(err)
			continue
		}
		l.ReplyPending[target.ID] = struct{}{}
	}

	// Add pair to the priority queue
	item := Item{
		ID:   l.id,
		Time: time,
	}
	heap.Push(&l.priorityQueue, &item)
	l.log.Info("waiting for replies")
	<-l.enterCS
	l.inCS = true
}

func (l *Lamport) CS_Leave() {
	l.priorityQueue.Pop()

	msg := application_pkg.Message{
		Source: l.id,
		Type:   application_pkg.RELEASE,
		Time:   math.Floor(time.Since(time.Now()).Seconds()*1000000) / 1000000,
	}

	for _, target := range l.neighbours {
		// Send targets
		l.log.Info("sending ", msg.Type, " to ", target.ID, " at ", msg.Time)
		err := l.send(target.ID, msg)
		if err != nil {
			l.log.Error(err)
			continue
		}
	}

	l.inCS = false
}

func (l *Lamport) ProcessMessage(connection net.Conn, nr int, stopCh chan struct{}) {
	for {
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			l.log.Error(err)
			return
		}
		msg := application_pkg.Message{}
		if err = json.Unmarshal(buffer[:mLen], &msg); err != nil {
			l.log.Error(err)
			continue
		}
		//if msg.Decode(buffer[:mLen]); err != nil {
		//	l.log.Error(err)
		//	continue
		//}

		l.log.Info("received ", msg.Type, " from ", msg.Source)
		delete(l.ReplyPending, msg.Source)
		switch msg.Type {
		case application_pkg.REQUEST:
			{
				item := Item{
					ID:   msg.Source,
					Time: msg.Time,
				}
				heap.Push(&l.priorityQueue, &item)
				// If not in cs then send reply
				if !l.inCS {
					obj := application_pkg.Message{
						Source: l.id,
						Type:   application_pkg.REPLY,
						Time:   msg.Time,
					}
					l.log.Info("sending ", obj.Type, " to ", msg.Source, " at ", msg.Time)
					err = l.send(msg.Source, obj)
					if err != nil {
						l.log.Error(err)
						continue
					}
				}
				break
			}
		case application_pkg.RELEASE:
			{
				l.priorityQueue.Pop()
				l.releaseReceived++
				l.log.Info("l.releaseReceived: ", l.releaseReceived)
				l.log.Info("match: ", len(l.neighbours) == l.releaseReceived)
				break
			}
		default:
			break
		}
		if len(l.ReplyPending) == 0 && l.priorityQueue.Len() > 0 && l.priorityQueue.Top().(*Item).ID == l.id {
			// Enter CS
			l.enterCS <- struct{}{}
		}
		if l.releaseReceived == nr*len(l.neighbours) {
			stopCh <- struct{}{}
			return
		}
	}
}
