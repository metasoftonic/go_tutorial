package main

import (
	"fmt"
	"github.com/metasoftonic/go_tutorial/session_two/greetings"
)

func main() {
	message := greetings.Hello("shadow")
	fmt.Println(message)
}
