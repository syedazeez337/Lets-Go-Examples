package main

import "fmt"

type language string

func main() {

	greetings := GreetMultiple("vi")
	fmt.Println(greetings)

	for i, val := range phrasebook {
		fmt.Println(i, val)
	}
	
}

var phrasebook = map[language]string{
	"en": "Hello world",
	"gk": "Χαίρετε Κόσμε",
	"fr": "Bonjour le monde",
	"ur": " ہیلو ",
	"vi": "Xin chào Thế Giới",
}

func GreetMultiple(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}

	return greeting
}
