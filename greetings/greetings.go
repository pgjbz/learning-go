package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) { //use a tuple to return string and maybe an error
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomGreetings(), name)
	return message, nil
}

func init() { //init function automatic execute
	rand.Seed(time.Now().UnixNano())
}

func randomGreetings() string {
	var formats []string = []string{
		"Hi, %v. Welcome!",
		"Hello, %v",
		"Salve, %v",
	}
	return formats[rand.Intn(len(formats))]
}
