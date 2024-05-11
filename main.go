package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/enuesaa/lkaw/repository"
)

func main() {
	var cli Cli
	kong.Parse(&cli,
		kong.UsageOnError(),
		kong.Name("lkaw"),
		kong.Description("A CLI tool to look at files anyway."),
	)

	repos := repository.New()
	if cli.Filenames[0] == "." {
		filenames, err := repos.Fs.ListFiles(".")
		if err != nil {
			log.Fatalf("Error: %s", err.Error())
		}
		cli.Filenames = filenames
	}

	if err := Serve(repos, cli); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
