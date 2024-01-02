package main

import (
	"communicate"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter the address of the server")
	var addr string

	fmt.Scanln(&addr)

	fmt.Printf("Addr is set to: %s\n", addr)
	fmt.Println("What is your username?")

	var name string
	fmt.Scanln(&name)

	conn, err := communicate.StartConn(addr, name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("You are now ready to send messages")
	var msg string
	for {
		fmt.Scanln(&msg)

		m := communicate.Message{MsgType: communicate.MsgEnd, Msg: msg}

		conn.Send(m)
	}

}
