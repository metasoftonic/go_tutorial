package greetings;

import "fmt"

func Hello(name string) string {
	message := fmt.Sprint("Hi, %v. Nice to meet  you", name)
	return message
}
