package communicate

import (
	"encoding/gob"
	"net"
)

type Connection struct {
	C       net.Conn
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func StartConn(addr string) (*Connection, error) {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return nil, err
	}

	c := NewConn(conn)
	return c, nil
}

func NewConn(conn net.Conn) *Connection {
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	return &Connection{C: conn, encoder: enc, decoder: dec}
}

func (c *Connection) Read() (Message, error) {
	var receivedMessage Message
	err := c.decoder.Decode(&receivedMessage)

	if err != nil {
		return Message{}, err
	}

	return receivedMessage, nil
}

func (c *Connection) Send(msg Message) {
	_ = c.encoder.Encode(msg)
}

func (c *Connection) Close() {
	err := c.C.Close()
	if err != nil {
		println(err.Error())
	}
}

func (c *Connection) GetAddr() string {
	return c.C.RemoteAddr().String()
}
