package main

import (
	"fmt"
	"greetings"
)

func main() {
	message := greetings.Hello("Lucas")
	fmt.Println(message)
}
