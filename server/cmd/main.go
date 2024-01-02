package main

import (
	"fmt"
	"os"
	. "server/internal/server"
)

func main() {
	// Init Server
	s, err := NewServer("127.0.0.1:8080")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	// Close server when done
	defer s.Close()

	fmt.Println("Listening.....")
	for {
		// Accept new incoming profile
		conn, err := s.Accept()

		// Print error if there is one
		if err != nil {
			println(err.Error())
			continue
		}

		// Handle profile
		go HandleConnection(conn)
	}
}
