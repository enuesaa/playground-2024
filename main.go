package main

import (
	"log"

	"github.com/enuesaa/txtsout/pkg/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
