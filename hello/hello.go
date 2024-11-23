package main

import "fmt"

const (
	spanish = "Spanish"
	german  = "German"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	germanPrefix  = "Heil, "
)

func main() {
	fmt.Println(Hello("you", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return Prefix2(language) + name
}

func Prefix2(language string) string {
	switch language {
	case "":
		return englishPrefix
	case spanish:
		return spanishPrefix
	case german:
		return germanPrefix
	}
	return ""
}
