package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	bookworms, err := loadBookworms("testdata/bookworm.json")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(bookworms)
}
