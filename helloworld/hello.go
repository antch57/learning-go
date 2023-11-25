package main

import (
	"fmt"
)

// Setup an enum for the languages available
type Language int

const (
	English Language = iota
	Spanish
	Borat
	French
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	boratHelloPrefix   = "Jagshemash, "
)

func Hello(name string, language Language) string {
	greeting := greetingPrefix(language)

	if name == "" {
		name = "world"
	}

	return greeting + name
}

func greetingPrefix(language Language) (prefix string) {
	switch language {
	case English:
		prefix = englishHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = frenchHelloPrefix
	case Borat:
		prefix = boratHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("ant", Borat))
}
