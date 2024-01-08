package main

import (
	"bufio"
	"client/internal/client"
	"communicate"
	"fmt"
	"os"
	"time"
)

func main() {
	clientObj := client.NewClient("127.0.0.1:8080")

	go loopRead(clientObj)

	fmt.Println("Enter message:")

	scanner := bufio.NewScanner(os.Stdin)
	var msg string
	for {
		if scanner.Scan() {
			msg = scanner.Text()
			m := communicate.Message{MsgType: communicate.Msg, Msg: []byte(msg)}

			clientObj.Send(m)
			time.Sleep(1000 * time.Millisecond)
		} else {
			fmt.Println("Error reading input")
			return
		}
	}
}

// / Loop reading messages from the server
func loopRead(c *client.Client) {
	for {
		time.Sleep(500 * time.Millisecond)

		m, e := c.Read()
		if e == nil && m.MsgType == communicate.Msg {
			fmt.Printf("%s\n", string(m.Msg))
		}
	}
}
