package server

import (
	"communicate"
	"server/internal/listener"
	"server/internal/objects"
	"server/internal/thread"

	"github.com/google/uuid"
)

type Server struct {
	address    string
	listener   chan *communicate.Connection
	incoming   chan objects.ServerMessage
	clients    map[uuid.UUID]*objects.Client
	newClients map[uuid.UUID]*objects.Client
}

func NewServer(addr string) *Server {
	l := make(chan *communicate.Connection, 3)
	i := make(chan objects.ServerMessage, 10)
	c := make(map[uuid.UUID]*objects.Client)
	nc := make(map[uuid.UUID]*objects.Client)

	return &Server{
		address:    addr,
		listener:   l,
		incoming:   i,
		clients:    c,
		newClients: nc,
	}
}

func (s *Server) StartListener() {
	go listener.NewListener(s.address, s.listener)
}

func (s *Server) IncomingConnection() {
	select {
	case c := <-s.listener:
		// Make client
		client := objects.NewClient(uuid.New(), "", c)

		// Add to waiting list
		s.addToWaiting(client)

		// Ask for name
		c.Send(communicate.Message{
			MsgType: communicate.AskName,
			Msg:     []byte(client.ID().String()),
		})

		// Start session
		go thread.NewSession(client.ID(), c, s.incoming)
	default:
		return
	}
}

func (s *Server) IncomingMessage() {
	select {
	case msg := <-s.incoming:
		s.handleMsg(msg)
	default:
		return
	}
}
