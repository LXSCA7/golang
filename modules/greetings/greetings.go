package greetings

import "fmt"

func Hello(name string) string {
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
