package objects

import (
	"communicate"
	"github.com/google/uuid"
)

type Client struct {
	id         uuid.UUID
	name       string
	connection *communicate.Connection
}

func NewClient(id uuid.UUID, name string, conn *communicate.Connection) *Client {
	return &Client{
		id:         id,
		name:       name,
		connection: conn,
	}
}

func (c *Client) Send(msg communicate.Message) {
	c.connection.Send(msg)
}

func (c *Client) ID() uuid.UUID       { return c.id }
func (c *Client) Name() string        { return c.name }
func (c *Client) Addr() string        { return c.connection.GetAddr() }
func (c *Client) SetName(name string) { c.name = name }
