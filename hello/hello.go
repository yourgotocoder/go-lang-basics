package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// Get a greeting message and print it.
	log.SetPrefix("greetins: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Sudu")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
