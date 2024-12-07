package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Serhat")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	pairs, err := greetings.Hellos([]string{"ali", "maçürı"})
	// pairs2, err := greetings.Hellos([]string{})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(pairs)
	// fmt.Println(pairs2)
}
