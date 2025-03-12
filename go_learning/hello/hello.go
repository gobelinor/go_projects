package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	
	spanish = "Spanish"
	french = "French"
)
	
func Hello(name, country string) string {
	if name == "" {
		name = "World"
	}
	return  greeting_prefix(country) + name
}

func greeting_prefix(country string) (prefix string) {
	switch country {	
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("woeerld", ""))
}
