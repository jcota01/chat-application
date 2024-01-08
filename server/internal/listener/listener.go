package listener

import (
	"communicate"
	"fmt"
	"net"
	"os"
)

type Listener struct {
	listener net.Listener
	channel  chan<- *communicate.Connection
}

func (l Listener) Accept() {
	conn, err := l.listener.Accept()
	if err == nil {
		l.channel <- communicate.NewConn(conn)
	}
}

// NewListener Start a listener on the given address
func NewListener(addr string, channel chan<- *communicate.Connection) {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	l := Listener{listener, channel}
	defer l.Close()

	for {
		l.Accept()
	}
}

// Close the listener
func (l Listener) Close() {
	err := l.listener.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
