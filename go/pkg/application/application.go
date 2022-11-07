package application

import (
	"net"
	"time"

	"github.com/kumarabd/CS6378-Project2/go/config"
)

type Neighbour struct {
	ID         string
	HostPort   *config.HostPort
	Connection net.Conn
}

type Application interface {
	CS_Enter(time float64)
	CS_Leave()
	ProcessMessage(msg *Message, nr int, stopCh chan struct{})
	SetClock(clock time.Time)
}
