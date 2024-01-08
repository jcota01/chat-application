package server

import (
	"communicate"
	"fmt"
	"server/internal/objects"
)

// / Handle incoming messages from clients
func (s *Server) handleMsg(msg objects.ServerMessage) {
	/// Get the client from the client list
	c := s.getClient(msg.ID)

	switch msg.T {
	case objects.ConnEnd:
		/// Remove the client from the client list
		s.remove(msg.ID)

	case objects.Msg:
		if c != nil {
			/// Add the name of the client to the message
			msg.Data.Msg = []byte(fmt.Sprintf("%s: %s", c.Name(), string(msg.Data.Msg)))
		}

		/// Send the message to all clients
		s.sendToClients(msg.Data)

	case objects.Name:
		nc := s.newClients[msg.ID]
		if nc != nil {
			/// Set the name of the client
			nc.SetName(string(msg.Data.Msg))

			/// Move the client from the waiting list to the client list
			s.moveFromWaiting(msg.ID)

			/// Send a ready message to the client
			client := s.getClient(msg.ID)
			client.Send(communicate.Message{
				MsgType: communicate.Ready,
				Msg:     []byte{},
			})
		}
	}
}
