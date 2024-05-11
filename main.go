package main

import (
	"log"

	"github.com/alecthomas/kong"
)

var cli struct {
	Port int `help:"port to serve." default:"3000"`
    Filenames []string `arg:"" required:"" name:"filename" help:"Files to open"`
	Version   Version `name:"version" short:"v" help:"Print version"`
}

func main() {
	kong.Parse(&cli,
		kong.UsageOnError(),
		kong.Name("lkaw"),
		kong.Description("A CLI tool to look at files anyway."),
	)

	if err := Serve(cli.Port); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
