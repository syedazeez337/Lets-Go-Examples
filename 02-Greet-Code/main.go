package main

import (
	"fmt"
)

func main() {
	greetings := greet()
	fmt.Println(greetings)
}

func greet() string {
	return "Hello Greetings"
}