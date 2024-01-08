package main

import (
	"fmt"
	. "server/internal/server"
	"time"
)

func main() {
	// Init Server
	s := NewServer("127.0.0.1:8080")

	s.StartListener()

	fmt.Println("Listening.....")
	for {
		s.IncomingConnection()
		s.IncomingMessage()

		time.Sleep(300 * time.Millisecond)
	}
}
