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

	names := []string{"Sudu", "Sudarshan", "Sun"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
