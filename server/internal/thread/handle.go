package thread

import (
	"communicate"
	"server/internal/objects"
)

func (s Session) HandleMessage(message communicate.Message) {

	switch message.MsgType {
	case communicate.Msg:
		sm := objects.ServerMessage{
			T:    objects.Msg,
			ID:   s.id,
			Data: message,
		}
		s.sendChannel <- sm
		return
	case communicate.Name:
		sm := objects.ServerMessage{
			T:    objects.Name,
			ID:   s.id,
			Data: message,
		}
		s.sendChannel <- sm
		return
	}
}
