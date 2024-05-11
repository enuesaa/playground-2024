package main

import (
	"log"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/enuesaa/lkaw/presenter"
)

var cli struct {
	Browser bool `short:"b" help:"open in web browser" default:"false"`
	Port int `help:"port to serve." default:"3000"`
	Filename string `short:"f" help:"filename to open"`
}

func main() {
	kong.Parse(&cli)

	if cli.Browser {
		if err := Serve(cli.Port); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
	}

	if cli.Filename != "" && strings.HasSuffix(cli.Filename, ".db") {
		if err := presenter.HandleSqlite(cli.Filename); err != nil {
			log.Fatalf(err.Error())
		}
	}
}
