package main

import (
	"fmt"
	"log"

	"go-modules/greetings"
)

func main() {
	//LOGGER SETUP
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Ong", "Katana", "Suka"}
	message, err := greetings.Hellos(names)

	//PRINT ERROR
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
