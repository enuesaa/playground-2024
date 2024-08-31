package main

import (
	"log"
	"os"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/usecase"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	app := &cli.App{
		Name:  "codetrailer",
		Usage: "",
		Action: func(*cli.Context) error {
			go usecase.Prompt(repos)

			return usecase.LaunchMenu()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
