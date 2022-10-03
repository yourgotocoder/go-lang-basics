package main

import (
	"fmt"

	"example/greeting"
)

func main() {
	message := greetings.Greetings("Sudu")
	fmt.Println(message)
}
