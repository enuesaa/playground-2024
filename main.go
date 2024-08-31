package main

import (
	"errors"
	"log"
	"os"
	"os/exec"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/erikgeiser/promptkit"
	"github.com/getlantern/systray"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	app := &cli.App{
		Name:  "codetrailer",
		Usage: "",
		Action: func(*cli.Context) error {
			for {
				args, err := repos.Log.Ask(">", "")
				if err != nil {
					if errors.Is(err, promptkit.ErrAborted) {
						return nil
					}
					return err
				}
				if args == "q" {
					break
				}
				cmd := exec.Command("bash", "-c", args)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err != nil {
					return err
				}
			}

			systray.Run(startMenu, endMenu)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
