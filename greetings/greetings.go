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

func Hellos(names []string) (map[string]string, error) {
	hellos := make(map[string]string) //instantiate a map
	for _, name := range names {      //in go only have for loops, this is a basicaly for (idx, name) in names.iter().enumerate() in Rust
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		hellos[name] = message
	}
	return hellos, nil
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
