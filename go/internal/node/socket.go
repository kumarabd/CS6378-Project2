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

//func Send() error {
//	_, err = connection.Write([]byte("Hello Server! Greetings."))
//}

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

//func send(message Message) {
//	m,err:=json.Marshal(&message)
//	if err!=nil {
//		return err
//	}

//	if _, err = connection.Write(m) {
//		return err
//	}
//}

//func listen() {
//	buffer := make([]byte, 1024)
//    mLen, err := connection.Read(buffer)
//    if err != nil {
//     	return err
//    }

//	m := Message{}
//	err = json.Unmarshal(buffer[:mLen], &m)
//	if err!=nil {
//		return err
//	}
//}
