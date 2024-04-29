package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
)

var cli struct {
	Serve bool `help:"serve" default:"false"`
	Port int `help:"port to serve." default:"3000"`
}

func main() {
	kong.Parse(&cli)
	fmt.Println(cli.Port)

	if cli.Serve {
		if err := Serve(); err != nil {
			log.Fatalf("Error: %s", &err)
		}
	}
}
