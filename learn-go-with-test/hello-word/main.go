package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix string = "Hello, "
const portugueseHelloPrefix string = "Olá, "
const spanishHelloPrefix string = "Hola, "

const english = "English"
const portuguese = "Portuguese"
const spanish = "Spanish"

func Hello(name, language string) string {
	prefix := greetingPrefix(language)

	if len(strings.TrimSpace(name)) == 0 {
		name = "World"
	}
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Paulo", "Portuguese"))
}
