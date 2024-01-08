package client

import (
	"communicate"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type Client struct {
	connection *communicate.Connection
	username   string
	id         uuid.UUID
}

func NewClient(addr string) *Client {
	conn, err := communicate.StartConn(addr)
	if err != nil {
		fmt.Println("Error starting connection")
		os.Exit(1)
	}

	name, id := name(conn)

	return &Client{connection: conn, username: name, id: id}
}

func (c *Client) Connection() *communicate.Connection {
	return c.connection
}

func (c *Client) Username() string {
	return c.username
}

func (c *Client) ID() uuid.UUID {
	return c.id
}

func (c *Client) Send(msg communicate.Message) {
	c.connection.Send(msg)
}

func (c *Client) Read() (communicate.Message, error) {
	return c.connection.Read()
}
