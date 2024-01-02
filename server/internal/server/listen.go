package server

import (
	"net"
)

type Connection struct {
	conn net.Conn
}

func (s Server) Accept() (Connection, error) {
	conn, err := s.listener.Accept()
	if err != nil {
		return Connection{}, err
	}

	return Connection{conn}, nil
}

func (c Connection) GetAddr() string {
	return c.conn.RemoteAddr().String()
}

func (c Connection) Read(buffer []byte) (string, error) {
	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}

func (c Connection) Close() {
	err := c.conn.Close()
	if err != nil {
		println(err.Error())
	}
}
