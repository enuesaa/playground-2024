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
		Name:    "codetrailer",
		Usage:   "A CLI tool to capture stdin/stdout and generate a step-by-step document.",
		Version: "0.0.1",
		Flags: []cli.Flag{
			// markdown を作成
			&cli.BoolFlag{
				Name:  "init",
				Usage: "Create new project",
			},
			&cli.StringFlag{
				Name:  "start",
				Usage: "Create new case or continue.",
			},
			&cli.BoolFlag{
				Name:  "preview",
				Usage: "Serve",
			},
			&cli.StringFlag{
				Name:  "project",
				Value: ".",
				Usage: "Path to project. default: current dir.",
			},
		},
		Action: func(*cli.Context) error {
			go usecase.Prompt(repos)

			return usecase.LaunchMenu()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
