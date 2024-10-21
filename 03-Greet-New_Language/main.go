package main

import "fmt"

func main() {
	engGreeting := Greet("en")
	frGreeting := Greet("fr")
	fmt.Println(engGreeting, frGreeting)
}

type language string

func Greet(lang language) string {
	switch lang {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}
