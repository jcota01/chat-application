package thread

import (
	"communicate"
	"errors"
	"github.com/google/uuid"
	"net"
	"os"
	"server/internal/objects"
	"syscall"
)

type Session struct {
	id          uuid.UUID
	Conn        *communicate.Connection
	sendChannel chan<- objects.ServerMessage
}

func NewSession(id uuid.UUID, conn *communicate.Connection, sc chan<- objects.ServerMessage) {
	s := Session{
		id:          id,
		Conn:        conn,
		sendChannel: sc,
	}

	s.Handle()
}

func (s Session) End() {
	s.Conn.Close()

	sm := objects.ServerMessage{
		T:    objects.ConnEnd,
		ID:   s.id,
		Data: communicate.Message{},
	}

	s.sendChannel <- sm
}

func (s Session) Handle() {
	defer s.End()

	for {
		// Read from connection
		// Returns false if the connection is broken
		if !s.readConnection() {
			break
		}
	}
}

func (s Session) readConnection() bool {
	msg, err := s.Conn.Read()

	if err != nil {
		// Convert to net error
		var netErr *net.OpError
		if errors.As(err, &netErr) {
			// Try to get sys call error from net error
			var sysErr *os.SyscallError
			if errors.As(netErr.Err, &sysErr) {
				// If the connection is ended
				if errors.Is(sysErr.Err, syscall.WSAECONNRESET) {
					return false
				}
			}
		}

		return true
	}

	s.HandleMessage(msg)

	return true
}
