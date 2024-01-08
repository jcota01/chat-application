package objects

import (
	"communicate"
	"github.com/google/uuid"
)

type Type uint8

const (
	ConnEnd Type = 0
	Msg     Type = 1
	Name    Type = 2
)

type ServerMessage struct {
	T    Type
	ID   uuid.UUID
	Data communicate.Message
}
