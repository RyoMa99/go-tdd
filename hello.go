package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "Spanish":
		return spanishHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}
