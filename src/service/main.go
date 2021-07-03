package main

import (
	"log"
	// Import local 'models' module
	"models"
)

func main() {
	// Create new "MyEvent" struct
	ev := new(models.MyEvent)

	// Defining the properties of the struct
	ev.Type = "My Type"
	ev.SecondParam = "Another param"

	// Print the "MyEvent" struxt to the console
	log.Printf("Event %v", ev)

	// Print the returned result of the "MyPublicFunc()"
	log.Println(models.MyPublicFunc())
}
