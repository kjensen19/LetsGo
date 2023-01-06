package main

// Delcares a package "main": package is a way to group functions, made up of all files in same dir

import "fmt"

// Part of standard library, for formatting text

import "rsc.io/quote"

//To add as a requirement and generate a go.sum for authenticating:
//$ go mod tidy

func main() {
	fmt.Println(quote.Go())
}

//main function executes by default when main package runs ($go run .)

//$go help for commands
