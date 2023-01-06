package main

// Delcares a package "main": package is a way to group functions, made up of all files in same dir

import (
	"fmt"

	"example.com/greetings"

	// Part of standard library, for formatting text

	"rsc.io/quote"
)

//To add as a requirement and generate a go.sum for authenticating:
//$ go mod tidy

func main() {
	message := greetings.Hello("Kyle")
	fmt.Println(message)
}

//main function executes by default when main package runs ($go run .)

//$go help for commands
