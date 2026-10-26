package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) string {
	switch lang {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	case german:
		return germanHelloPrefix
	}
	return englishHelloPrefix
}

func main() {
	fmt.Println(Hello("world", "English"))
}
