package main

import (
	"log"
	"models"
)

func main() {
	ev := new(models.MyEvent)
	ev.Type = "My Type"
	ev.SecondParam = "Another param"
	log.Printf("Event %v", ev)
}
