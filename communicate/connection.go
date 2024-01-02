package communicate

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Connection struct {
	c       net.Conn
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func StartConn(addr string, name string) (Connection, error) {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return Connection{}, err
	}

	c := NewConn(conn)
	sendName(c, name)

	return c, nil
}

func NewConn(conn net.Conn) Connection {
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	return Connection{conn, enc, dec}
}

func (c Connection) Read() (Message, error) {
	var receivedMessage Message
	err := c.decoder.Decode(&receivedMessage)

	if err != nil {
		return Message{}, err
	}

	return receivedMessage, nil
}

func (c Connection) Send(msg Message) {
	err := c.encoder.Encode(msg)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (c Connection) Close() {
	err := c.c.Close()
	if err != nil {
		println(err.Error())
	}
}

func (c Connection) GetAddr() string {
	return c.c.RemoteAddr().String()
}
