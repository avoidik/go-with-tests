package main

import "fmt"

const (
	englishHelloPrexif = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	langSpanish        = "Spanish"
	langFrench         = "French"
)

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case langFrench:
		prefix = frenchHelloPrefix
	case langSpanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrexif
	}
	return
}

func Hello(name, lang string) string {
	if len(name) == 0 {
		name = "world"
	}
	return greetingPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
