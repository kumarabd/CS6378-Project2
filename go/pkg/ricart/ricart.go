package ricart

import (
	"time"

	"github.com/kumarabd/CS6378-Project2/go/logger"
	application_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/application"
	cmap "github.com/kumarabd/CS6378-Project2/go/pkg/sync_map"
)

type Ricart struct {
	id           string
	requestTime  int64
	neighbours   map[string]*application_pkg.Neighbour
	deferred     *cmap.SyncMap
	ReplyPending *cmap.SyncMap
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
		requestTime:  10000000,
		deferred:     cmap.New(),
		ReplyPending: cmap.New(),
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
		//l.ReplyPending[target.ID] = struct{}{}
		err = l.ReplyPending.Set(target.ID, struct{}{})
		if err != nil {
			l.log.Error(err)
			continue
		}
	}
	time.Sleep(200 * time.Millisecond)

	for {
		for _, target := range l.neighbours {
			// Send targets
			_, err := l.ReplyPending.Get(target.ID)
			if err == nil {
				l.log.WithField("clock", l.scalarClock).Info("sending reattempt ", msg.Type, " to ", target.ID)
				err := l.send(target.ID, &msg)
				if err != nil {
					l.log.Error(err)
					continue
				}
			}
		}
		time.Sleep(100 * time.Millisecond)
		if l.ReplyPending.Size() == 0 {
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
	for target, _ := range l.deferred.DeepCopy() {
		// Send targets
		l.log.WithField("clock", l.scalarClock).Info("sending ", msg.Type, " to ", target)
		err := l.send(target, &msg)
		if err != nil {
			l.log.Error(err)
			continue
		}
		l.replySent++
	}
	l.deferred = cmap.New()
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

				_, waitDependency := l.ReplyPending.Get(msg.Source)
				// break timestamp ties with the last condition in the below if statement
				if l.requestTime < 0 || (l.requestTime > 0 && msg.Time < l.requestTime) || (l.requestTime > 0 && (msg.Time == l.requestTime) && (msg.Source < l.id) && waitDependency == nil) {
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
					l.deferred.Set(msg.Source, struct{}{})
				}
				break
			}
		case application_pkg.REPLY:
			{
				l.ReplyPending.Delete(msg.Source)
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
		if l.ReplyPending.Size() == 0 && !l.inCS {
			// Enter CS
			l.enterCS <- struct{}{}
			l.inCS = true
		}
	}
}

func (l *Ricart) SetClock(clock time.Time) {
	l.startTime = clock
}
