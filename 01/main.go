package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter your name: ")

	var name string

	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello, %s! Welcome to Go programming.\n", name)
}
