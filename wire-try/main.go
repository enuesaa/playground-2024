package main

import (
	"log"
)

func main() {
	e, err := InitializeEvent()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	e.Start()
}
