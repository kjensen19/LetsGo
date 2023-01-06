package main

// Delcares a package "main": package is a way to group functions, made up of all files in same dir

import (
	"fmt"
	"log"

	"example.com/greetings"
	// "golang.org/x/text/message"
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

	//Slice of names
	names := []string{"gladys", "Samantha", "Darrin"}

	// message, err := greetings.Hello("Kyle") for single name

	messages, err := greetings.Hellos(names)
	//If error, print it and exit
	//Returns: greetings: 	empty name
	// exit status 1
	if err != nil {
		log.Fatal(err)
	}
	//if no error, print the message
	fmt.Println(messages)
}

//main function executes by default when main package runs ($go run .)

//$go help for commands
