package handler

import (
	"fmt"
	"server/internal/server"
	"server/internal/users"
)

func HandleConnection(conn server.Connection) {
	defer conn.Close()

	user, err := users.NewUser(conn)
	if err != nil {
		return
	}

	buffer := make([]byte, 1024)

	for {
		msg, err := user.Conn.Read(buffer)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("%s: %s\n", user.Username(), msg)
	}
}
