package profile

import (
	"fmt"
	"net"
)

type Connection struct {
	conn net.Conn
}

func NewConnection(addr string) (Connection, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return Connection{}, err
	}

	c := Connection{conn: conn}

	setName(c)

	return c, err
}

func (c Connection) Send(msg string) {
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func setName(c Connection) {
	var name string
	fmt.Println("What is your profile name?")
	fmt.Scanln(&name)

	c.Send(name)
}
