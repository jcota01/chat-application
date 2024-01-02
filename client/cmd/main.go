package main

import (
	"client/internal/profile"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the address of the server")
	var addr string

	fmt.Scanln(&addr)

	fmt.Printf("Addr is set to: %s\n", addr)

	conn, err := profile.NewConnection(addr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("You are now ready to send messages")
	var msg string
	for {
		fmt.Scanln(&msg)

		conn.Send(msg)
	}
}
