package ricart

import (
	"time"

	"github.com/kumarabd/CS6378-Project2/go/logger"
	application_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/application"
)

type Ricart struct {
	id           string
	requestTime  int64
	neighbours   map[string]*application_pkg.Neighbour
	deferred     map[string]struct{}
	ReplyPending map[string]struct{}
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
	currClock := time.Since(l.startTime).Milliseconds()
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
			l.log.Error(err)
			continue
		}
		l.ReplyPending[target.ID] = struct{}{}
	}
	l.requestTime = currClock

	l.log.WithField("clock", currClock).Info("waiting for replies")
	l.log.WithField("clock", currClock).Info("l.ReplyPending: ", l.ReplyPending)
	l.log.WithField("clock", currClock).Info("l.deferred: ", l.deferred)
	<-l.enterCS
}

func (l *Ricart) CS_Leave() {
	currClock := time.Since(l.startTime).Milliseconds()
	msg := application_pkg.Message{
		Source: l.id,
		Type:   application_pkg.REPLY,
		Time:   currClock,
	}

	for target, _ := range l.deferred {
		// Send targets
		l.log.WithField("clock", currClock).Info("sending ", msg.Type, " to ", target)
		err := l.send(target, &msg)
		if err != nil {
			l.log.Error(err)
			continue
		}
		l.replySent++
	}
	l.deferred = make(map[string]struct{})
	l.log.WithField("clock", currClock).Info(" l.replySent: ", l.replySent)
	l.log.WithField("clock", currClock).Info(" l.ReplyPending: ", l.ReplyPending)
	l.log.WithField("clock", currClock).Info("l.deferred: ", l.deferred)
	l.requestTime = -1
	l.inCS = false
}

func (l *Ricart) ProcessMessage(msgs []*application_pkg.Message, nr int, stopCh chan struct{}) {
	for _, msg := range msgs {
		currClock := time.Since(l.startTime).Milliseconds()
		l.log.WithField("clock", currClock).Info("received ", msg.Type, " from ", msg.Source)
		switch msg.Type {
		case application_pkg.REQUEST:
			{
				l.log.WithField("clock", currClock).Info(" msg.Time: ", msg.Time)
				l.log.WithField("clock", currClock).Info(" l.requestTime: ", l.requestTime)
				if l.requestTime < 0 || (l.requestTime > 0 && msg.Time < l.requestTime) {
					obj := application_pkg.Message{
						Source: l.id,
						Type:   application_pkg.REPLY,
						Time:   currClock,
					}
					l.log.WithField("clock", msg.Time).Info("sending ", obj.Type, " to ", msg.Source)
					err := l.send(msg.Source, &obj)
					if err != nil {
						l.log.Error(err)
						break
					}
					l.replySent++
				} else {
					l.deferred[msg.Source] = struct{}{}
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
		l.log.WithField("clock", currClock).Info(" l.replySent: ", l.replySent)
		l.log.WithField("clock", currClock).Info(" l.ReplyPending: ", l.ReplyPending)
		l.log.WithField("clock", currClock).Info("l.deferred: ", l.deferred)
		l.log.WithField("clock", currClock).Info("match: ", nr*len(l.neighbours) == l.replySent)
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
