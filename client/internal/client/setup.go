package client

import (
	"communicate"
	"fmt"
	"github.com/google/uuid"
	"os"
)

func name(conn *communicate.Connection) (string, uuid.UUID) {
	// Get the uuid from the server
	m, e := conn.Read()
	if e != nil || m.MsgType != communicate.AskName {
		fmt.Println("Server did not ask for name")
		os.Exit(1)
	}

	// Parse the uuid
	id, e := uuid.Parse(string(m.Msg))
	if e != nil {
		fmt.Println("Error receiving uuid from server")
		os.Exit(1)
	}

	// Send the name to the server
	name := askName()
	conn.Send(communicate.Message{MsgType: communicate.Name, Msg: []byte(name)})

	// Get the ready message from the server
	m, e = conn.Read()
	if e != nil || m.MsgType != communicate.Ready {
		fmt.Println("Server did not accept name")
		os.Exit(1)
	}

	return name, id
}
