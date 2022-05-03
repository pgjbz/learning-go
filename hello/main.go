package main

import (
	"fmt"
	"log"

	"pgjbz.dev/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Paulo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	names := []string{"Paulo", "Gabriel"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
