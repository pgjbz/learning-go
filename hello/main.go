package main

import (
	"fmt"
	"pgjbz.dev/greetings"
)

func main() {
	var message string = greetings.Hello("Paulo")
	fmt.Println(message)
}
