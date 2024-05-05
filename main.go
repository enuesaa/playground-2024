package main

import (
	"log"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/enuesaa/txtsout/repository"
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
		db := repository.DBRepository{}
		if err := db.Open(cli.Filename); err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		res, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		for res.Next() {
			var schema string
			if err := res.Scan(&schema); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			log.Println(schema)
		}
	}
}
