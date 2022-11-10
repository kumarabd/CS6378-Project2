package lamport

import (
	"container/heap"
	"fmt"
	"time"

	"github.com/kumarabd/CS6378-Project2/go/logger"
	application_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/application"
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
	startTime       time.Time
	requestTime     int64
}

func New(id string, neighbours map[string]*application_pkg.Neighbour, log logger.Handler) (*Lamport, error) {
	lamport := Lamport{
		id:              id,
		neighbours:      neighbours,
		log:             log,
		enterCS:         make(chan struct{}),
		releaseReceived: 0,
		ReplyPending:    make(map[string]struct{}),
		requestTime:     -1,
		inCS:            false,
	}

	lamport.priorityQueue = make(PriorityQueue, 0)
	heap.Init(&lamport.priorityQueue)

	return &lamport, nil
}

func (l *Lamport) send(target string, msg *application_pkg.Message) error {
	conn := l.neighbours[target].Connection
	for conn == nil {
	}
	_, err := conn.Write(msg.Encode())
	if err != nil {
		return err
	}
	return nil
}

func (l *Lamport) CS_Enter() {
	currClock := time.Since(l.startTime).Milliseconds()
	// Add pair to the priority queue
	item := Item{
		ID:   l.id,
		Time: currClock,
	}
	heap.Push(&l.priorityQueue, &item)

	msg := application_pkg.Message{
		Source: l.id,
		Type:   application_pkg.REQUEST,
		Time:   currClock,
	}

	for _, target := range l.neighbours {
		// Send targets
		l.log.WithField("clock", currClock).Info("sending ", msg.Type, " to ", target.ID)
		err := l.send(target.ID, &msg)
		if err != nil {
			fmt.Println(err)
			l.log.Error(err)
			continue
		}
		l.ReplyPending[target.ID] = struct{}{}
	}
	l.requestTime = currClock

	l.log.WithField("clock", currClock).Info("waiting for replies")
	<-l.enterCS
}

func (l *Lamport) CS_Leave() {
	currClock := time.Since(l.startTime).Milliseconds()
	l.priorityQueue.Pop() // careful

	msg := application_pkg.Message{
		Source: l.id,
		Type:   application_pkg.RELEASE,
		Time:   currClock,
	}

	for _, target := range l.neighbours {
		// Send targets
		l.log.WithField("clock", currClock).Info("sending ", msg.Type, " to ", target.ID)
		err := l.send(target.ID, &msg)
		if err != nil {
			l.log.Error(err)
			continue
		}
	}
	l.requestTime = -1

	l.inCS = false
}

func (l *Lamport) ProcessMessage(msgs []*application_pkg.Message, nr int, stopCh chan struct{}) {
	for _, msg := range msgs {
		currClock := time.Since(l.startTime).Milliseconds()
		l.log.WithField("clock", currClock).Info("received ", msg.Type, " from ", msg.Source)
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
					currClock := time.Since(l.startTime).Milliseconds()
					obj := application_pkg.Message{
						Source: l.id,
						Type:   application_pkg.REPLY,
						Time:   currClock,
					}
					l.log.WithField("clock", currClock).Info("sending ", obj.Type, " to ", msg.Source)
					err := l.send(msg.Source, &obj)
					if err != nil {
						l.log.Error(err)
					}
				}

				if l.requestTime > 0 && msg.Time > l.requestTime {
					delete(l.ReplyPending, msg.Source)
				}
				break
			}
		case application_pkg.REPLY:
			{
				delete(l.ReplyPending, msg.Source)
				break
			}
		case application_pkg.RELEASE:
			{
				if l.requestTime > 0 && msg.Time > l.requestTime {
					delete(l.ReplyPending, msg.Source)
				}
				l.priorityQueue.Pop() // remove that particular process
				l.releaseReceived++
				break
			}
		default:
			break
		}
		l.log.WithField("clock", currClock).Info(l.ReplyPending)
		l.log.WithField("clock", currClock).Info(l.priorityQueue)
		l.log.WithField("clock", currClock).Info("l.releaseReceived: ", l.releaseReceived)
		l.log.WithField("clock", currClock).Info("target: ", nr*len(l.neighbours))
		if l.releaseReceived == nr*len(l.neighbours) {
			l.log.WithField("clock", currClock).Info("sending out")
			stopCh <- struct{}{}
			return
		}
		if len(l.ReplyPending) == 0 && l.priorityQueue.Len() > 0 && l.priorityQueue.Top().(*Item).ID == l.id && !l.inCS {
			// Enter CS
			l.enterCS <- struct{}{}
			l.inCS = true
		}
	}
}

func (l *Lamport) SetClock(clock time.Time) {
	l.startTime = clock
}
