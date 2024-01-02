package server

import (
	"communicate"
	"fmt"
	"net"
	"os"
	"server/internal/users"
)

type Server struct {
	address  string
	listener net.Listener
}

func NewServer(addr string) (Server, error) {
	// Create TCP listener
	listener, err := net.Listen("tcp", addr)

	// Check for err is TCP listener
	if err != nil {
		return Server{}, err
	}

	return Server{address: addr, listener: listener}, nil
}

func (s Server) Close() {
	err := s.listener.Close()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func (s Server) Accept() (communicate.Connection, error) {
	conn, err := s.listener.Accept()
	if err != nil {
		return communicate.Connection{}, err
	}

	return communicate.NewConn(conn), nil
}

func HandleConnection(conn communicate.Connection) {
	defer conn.Close()

	user := users.User{Conn: conn}
	for {
		msg, e := user.Conn.Read()
		if e != nil {
			fmt.Println(e.Error())
		}

		switch msg.MsgType {
		case communicate.Name:
			user.Name = msg.Msg
			fmt.Printf("User %s has joined.\n", user.Name)
		case communicate.MsgCont:
			fmt.Printf("%s: %s\n", user.Name, msg.Msg)
		case communicate.MsgEnd:
			fmt.Printf("%s: %s\n", user.Name, msg.Msg)
		}
	}
}
