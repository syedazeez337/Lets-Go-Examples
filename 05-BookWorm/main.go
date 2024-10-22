package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("testdata/bookworm.json")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)
}
