package main

import (
	"log"
	"os"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/enuesaa/codetrailer/internal/usecase"
	"github.com/urfave/cli/v2"
)

// Dir structure
//   README.md
//   steps/0.aa.md
//   steps/1.bb.md
//   steps/0.image-name.png

// Commands
// - codetrailer init ... this create markdown
// - codetrailer record ... create new step file
// - codetrailer preview
// - codetrailer export-pdf

func main() {
	repos := repository.New()

	app := &cli.App{
		Name:    "codetrailer",
		Usage:   "A CLI tool to capture stdin/stdout and generate a step-by-step document.",
		Version: "0.0.1",
		Action: func(*cli.Context) error {
			go usecase.Prompt(repos)

			return usecase.LaunchMenu()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
