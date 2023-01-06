package main

// Delcares a package "main": package is a way to group functions, made up of all files in same dir

import (
	"fmt"
	"log"

	"example.com/greetings"
	// Part of standard library, for formatting text
	// "rsc.io/quote"
)

//To add as a requirement and generate a go.sum for authenticating:
//$ go mod tidy

func main() {
	// Set props for the Logger:
	// log entry prefix and a flag to disable printing the time,
	// source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Kyle")
	//If error, print it and exit
	//Returns: greetings: 	empty name
	// exit status 1
	if err != nil {
		log.Fatal(err)
	}
	//if no error, print the message
	fmt.Println(message)
}

//main function executes by default when main package runs ($go run .)

//$go help for commands
