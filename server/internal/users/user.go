package users

import (
	"communicate"
)

type User struct {
	Name string
	Conn communicate.Connection
}

func (u User) Address() string {
	return u.Conn.GetAddr()
}
