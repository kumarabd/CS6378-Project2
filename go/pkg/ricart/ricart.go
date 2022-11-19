package ricart

import (
	"time"
	"sync"
	"github.com/kumarabd/CS6378-Project2/go/logger"
	application_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/application"
)

var wg sync.WaitGroup
var mutex = &sync.Mutex{}

type Ricart struct {
	id           string
	requestTime  int64
	neighbours   map[string]*application_pkg.Neighbour
	deferred     map[string]struct{}
	ReplyPending map[string]struct{}
	scalarClock  int64
	replySent    int
	enterCS      chan struct{}
	inCS         bool
	log          logger.Handler
	startTime    time.Time
}

func New(id string, neighbours map[string]*application_pkg.Neighbour, log logger.Handler) (*Ricart, error) {
	ricart := Ricart{
		id:           id,
		neighbours:   neighbours,
		log:          log,
		enterCS:      make(chan struct{}),
		replySent:    0,
		scalarClock:  0,
		requestTime: 10000000,
		deferred:     make(map[string]struct{}),
		ReplyPending: make(map[string]struct{}),
		inCS:         false,
	}
	return &ricart, nil
}

func (l *Ricart) send(target string, msg *application_pkg.Message) error {
	conn := l.neighbours[target].Connection
	for conn == nil {
	}
	_, err := conn.Write(msg.Encode())
	if err != nil {
		return err
	}
	return nil
}

func (l *Ricart) CS_Enter() {
	// currClock := time.Since(l.startTime).Milliseconds()
	l.scalarClock++
	msg := application_pkg.Message{
		Source: l.id,
		Type:   application_pkg.REQUEST,
		Time:   l.scalarClock,
	}
	// should it be initialized before l.scalarClock++?
	l.requestTime = l.scalarClock
	for _, target := range l.neighbours {
		// Send targets
		l.log.WithField("clock", l.scalarClock).Info("sending ", msg.Type, " to ", target.ID)
		err := l.send(target.ID, &msg)
		if err != nil {
			l.log.Error(err)
			continue
		}
		l.ReplyPending[target.ID] = struct{}{}
	}
	time.Sleep(200 * time.Millisecond)
	
	for {
		for _, target := range l.neighbours {
			// Send targets
			_, present := l.ReplyPending[target.ID]
			if present {
				l.log.WithField("clock", l.scalarClock).Info("sending reattempt ", msg.Type, " to ", target.ID)
				err := l.send(target.ID, &msg)
				if err != nil {
					l.log.Error(err)
					continue
				}
			}
		}
		time.Sleep(100 * time.Millisecond)
		if len(l.ReplyPending) == 0 {
			break
		}
		l.log.WithField("clock", l.scalarClock).Info("waiting for replies")
		l.log.WithField("clock", l.scalarClock).Info("l.ReplyPending: ", l.ReplyPending)
		l.log.WithField("clock", l.scalarClock).Info("l.deferred: ", l.deferred)
	}

	l.log.WithField("clock", l.scalarClock).Info("executing cs")
	<-l.enterCS
}

func (l *Ricart) CS_Leave() {
	// currClock := time.Since(l.startTime).Milliseconds()
	l.log.WithField("clock", l.scalarClock).Info("leaving cs")
	l.scalarClock++
	msg := application_pkg.Message{
		Source: l.id,
		Type:   application_pkg.REPLY,
		Time:   l.scalarClock,
	}

	l.requestTime = -1
	for target, _ := range l.deferred {
		// Send targets
		l.log.WithField("clock", l.scalarClock).Info("sending ", msg.Type, " to ", target)
		err := l.send(target, &msg)
		if err != nil {
			l.log.Error(err)
			continue
		}
		l.replySent++
	}
	mutex.Lock()
	l.deferred = make(map[string]struct{})
	mutex.Unlock()
	l.log.WithField("clock", l.scalarClock).Info(" l.replySent: ", l.replySent)
	l.log.WithField("clock", l.scalarClock).Info(" l.ReplyPending: ", l.ReplyPending)
	l.log.WithField("clock", l.scalarClock).Info("l.deferred: ", l.deferred)
	l.inCS = false
}

func (l *Ricart) ProcessMessage(msgs []*application_pkg.Message, nr int, stopCh chan struct{}) {
	for _, msg := range msgs {
		// currClock := time.Since(l.startTime).Milliseconds()
		l.log.WithField("clock", l.scalarClock).Info("received ", msg.Type, " from ", msg.Source)

		// piggyback on incoming timestamp
		if l.scalarClock < msg.Time {
			l.scalarClock = msg.Time + 1
		}

		switch msg.Type {
		case application_pkg.REQUEST:
			{
				time.Sleep(10 * time.Millisecond)
				// currClock := time.Since(l.startTime).Milliseconds()
				l.log.WithField("clock", l.scalarClock).Info(" msg.Time: ", msg.Time)
				l.log.WithField("clock", l.scalarClock).Info(" l.requestTime: ", l.requestTime)
				
				_, waitDependency := l.ReplyPending[msg.Source]
				// break timestamp ties with the last condition in the below if statement
				if l.requestTime < 0 || (l.requestTime > 0 && msg.Time < l.requestTime) || (l.requestTime > 0 && (msg.Time == l.requestTime) && (msg.Source < l.id) && waitDependency) {
					obj := application_pkg.Message{
						Source: l.id,
						Type:   application_pkg.REPLY,
						Time:   l.scalarClock,
					}
					l.log.WithField("clock", msg.Time).Info("sending ", obj.Type, " to ", msg.Source)
					err := l.send(msg.Source, &obj)
					if err != nil {
						l.log.Error(err)
						break
					}
					l.replySent++
				} else {
					mutex.Lock()
					l.deferred[msg.Source] = struct{}{}
					mutex.Unlock()
				}
				break
			}
		case application_pkg.REPLY:
			{
				delete(l.ReplyPending, msg.Source)
				break
			}
		default:
			break
		}
		l.log.WithField("clock", l.scalarClock).Info(" l.replySent: ", l.replySent)
		l.log.WithField("clock", l.scalarClock).Info(" l.ReplyPending: ", l.ReplyPending)
		l.log.WithField("clock", l.scalarClock).Info("l.deferred: ", l.deferred)
		l.log.WithField("clock", l.scalarClock).Info("match: ", nr*len(l.neighbours) == l.replySent)
		if l.replySent == nr*len(l.neighbours) {
			stopCh <- struct{}{}
			return
		}
		if len(l.ReplyPending) == 0 && !l.inCS {
			// Enter CS
			l.enterCS <- struct{}{}
			l.inCS = true
		}
	}
}

func (l *Ricart) SetClock(clock time.Time) {
	l.startTime = clock
}


