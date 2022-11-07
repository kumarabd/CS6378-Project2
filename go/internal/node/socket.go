package node

import (
	"net"
)

type Channel struct {
	listener net.Listener
}

func ConnectHost(host, port string) (net.Conn, error) {
	return net.Dial("tcp", host+":"+port)
}

func NewChannel(host, port string) (*Channel, error) {
	srv, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}
	return &Channel{
		listener: srv,
	}, nil
}

func (c *Channel) Listen() (net.Conn, error) {
	connection, err := c.listener.Accept()
	if err != nil {
		return nil, err
	}
	return connection, nil
}
