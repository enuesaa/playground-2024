package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/enuesaa/lkaw/repository"
)

type CLI struct {
	Port int `help:"port to serve." default:"3000"`
    Filenames []string `arg:"" required:"" name:"filename" help:"Files to open"`
	Version   Version `name:"version" short:"v" help:"Print version"`
}

func main() {
	var cli CLI
	kong.Parse(&cli,
		kong.UsageOnError(),
		kong.Name("lkaw"),
		kong.Description("A CLI tool to look at files anyway."),
	)

	repos := repository.New()
	if err := Serve(repos, cli); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
