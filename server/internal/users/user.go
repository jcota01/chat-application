package users

import (
	"fmt"
	"server/internal/server"
)

type User struct {
	name string
	Conn server.Connection
}

func (u User) Address() string {
	return u.Conn.GetAddr()
}

func (u User) Username() string { return u.name }

func (u User) SetUsername(name string) {
	u.name = name
}

func NewUser(conn server.Connection) (User, error) {
	buffer := make([]byte, 1024)

	name, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}

	fmt.Printf("New User:\n\tName:%s\n\tAddress:%s\n", name, conn.GetAddr())

	return User{name, conn}, nil
}
