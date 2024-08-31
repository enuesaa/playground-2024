package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/getlantern/systray"
	"github.com/urfave/cli/v2"
)

func main() {
	repos := repository.New()

	for {
		args, err := repos.Log.Ask(">", "")
		if err != nil {
			log.Fatal(err)
		}
		if args == "q" {
			break
		}
		cmd := exec.Command("bash", "-c", args)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			os.Exit(1)
		}
	}

	// start up prompt here.
	// - write contents
	// - capture screenshot
	// struct に詰める
	// create pdf file
	systray.Run(startMenu, endMenu)

	app := &cli.App{
		Name:  "codetrailer",
		Usage: "",
		Action: func(*cli.Context) error {
			fmt.Println("hello")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
