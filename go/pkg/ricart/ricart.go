package ricart

import (
	"math"
	"net"
	"time"

	"github.com/kumarabd/CS6378-Project2/go/pkg/application"
	application_pkg "github.com/kumarabd/CS6378-Project2/go/pkg/application"
	"github.com/realnighthawk/bucky/logger"
)

type Ricart struct {
	id           string
	neighbours   map[string]*application_pkg.Neighbour
	ReplyPending map[string]struct{}
	enterCS      chan struct{}
	inCS         bool
	log          logger.Handler
}

func New(id string, neighbours map[string]*application_pkg.Neighbour, log logger.Handler) (*Ricart, error) {
	ricart := Ricart{
		id:         id,
		neighbours: neighbours,
		log:        log,
		enterCS:    make(chan struct{}),
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

func (l *Ricart) CS_Enter(time float64) {}

func (l *Ricart) CS_Leave() {}

func (l *Ricart) ProcessMessage(connection net.Conn) {}

func (l *Ricart) SendExit() {
	msg := application.Message{
		Source: l.id,
		Type:   application.EXIT,
		Time:   math.Floor(time.Since(time.Now()).Seconds()*1000000) / 1000000,
	}
	for _, target := range l.neighbours {
		err := l.send(target.ID, msg)
		if err != nil {
			l.log.Error(err)
			continue
		}
	}
}
