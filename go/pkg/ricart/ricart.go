package ricart

import (
	"encoding/json"
	"math"
	"net"
	"time"

	application_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/application"
	"github.com/realnighthawk/bucky/logger"
)

type Ricart struct {
	id            string
	requestTime   float64
	neighbours    map[string]*application_pkg.Neighbour
	deferred      map[string]struct{}
	ReplyPending  map[string]struct{}
	replyReceived int
	enterCS       chan struct{}
	inCS          bool
	log           logger.Handler
}

func New(id string, neighbours map[string]*application_pkg.Neighbour, log logger.Handler) (*Ricart, error) {
	ricart := Ricart{
		id:            id,
		neighbours:    neighbours,
		log:           log,
		enterCS:       make(chan struct{}),
		replyReceived: 0,
	}

	ricart.ReplyPending = make(map[string]struct{})
	ricart.inCS = false

	return &ricart, nil
}

func (l *Ricart) send(target string, msg application_pkg.Message) error {
	conn := l.neighbours[target].Connection
	_, err := conn.Write(msg.Encode())
	if err != nil {
		return err
	}
	return nil
}

func (l *Ricart) CS_Enter(time float64) {
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
	l.requestTime = time

	l.log.Info("waiting for replies")
	<-l.enterCS
	l.inCS = true
}

func (l *Ricart) CS_Leave() {
	msg := application_pkg.Message{
		Source: l.id,
		Type:   application_pkg.REPLY,
		Time:   math.Floor(time.Since(time.Now()).Seconds()*1000000) / 1000000,
	}

	for target, _ := range l.deferred {
		// Send targets
		l.log.Info("sending ", msg.Type, " to ", target, " at ", msg.Time)
		err := l.send(target, msg)
		if err != nil {
			l.log.Error(err)
			continue
		}
	}

	l.inCS = false
}

func (l *Ricart) ProcessMessage(connection net.Conn, nr int, stopCh chan struct{}) {
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
		l.log.Info("received ", msg.Type, " from ", msg.Source)
		switch msg.Type {
		case application_pkg.REQUEST:
			{
				if len(l.ReplyPending) == 0 || msg.Time <= l.requestTime {
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
				} else {
					l.deferred[msg.Source] = struct{}{}
				}
				break
			}
		case application_pkg.REPLY:
			{
				delete(l.ReplyPending, msg.Source)
				l.replyReceived++
				break
			}
		default:
			break
		}
		if len(l.ReplyPending) == 0 {
			// Enter CS
			l.enterCS <- struct{}{}
		}
		if l.replyReceived == nr*len(l.neighbours) {
			stopCh <- struct{}{}
			return
		}
	}
}
