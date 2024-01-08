package client

import (
	"fmt"
)

func askName() string {
	fmt.Println("What is your username?")

	var name string
	fmt.Scanln(&name)

	return name
}
