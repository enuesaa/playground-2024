package main

import (
	"log"

	"github.com/enuesaa/lkaw/repository"
)

func main() {
	repos := repository.New()

	cli, err := Parse(repos)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if err := Serve(repos, cli); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
