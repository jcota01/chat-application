package server

import (
	"net"
	"os"
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
