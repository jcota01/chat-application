package server

import (
	"communicate"
	"github.com/google/uuid"
	"server/internal/objects"
)

func (s *Server) add(c *objects.Client) {
	s.clients[c.ID()] = c
}

func (s *Server) remove(id uuid.UUID) {
	delete(s.clients, id)
}

func (s *Server) addToWaiting(c *objects.Client) {
	s.newClients[c.ID()] = c
}

func (s *Server) moveFromWaiting(id uuid.UUID) {
	c := s.newClients[id]
	if c != nil {
		s.add(c)
		delete(s.newClients, c.ID())
	}
}

func (s *Server) sendToClients(msg communicate.Message) {
	for _, c := range s.clients {
		c.Send(msg)
	}
}

func (s *Server) getClient(id uuid.UUID) *objects.Client {
	return s.clients[id]
}
