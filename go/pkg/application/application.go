package application

import (
	"net"

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
	ProcessMessage(conn net.Conn)
	SendExit()
}
