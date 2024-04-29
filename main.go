package main

import (
	"log"

	"github.com/alecthomas/kong"
)

var cli struct {
	Browser bool `short:"b" help:"open in web browser" default:"false"`
	Port int `help:"port to serve." default:"3000"`
}

func main() {
	kong.Parse(&cli)

	if cli.Browser {
		if err := Serve(cli.Port); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	}
}
