package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/getlantern/systray"
	"github.com/playwright-community/playwright-go"
	"github.com/urfave/cli/v2"
)

func main() {
    if err := playwright.Install(&playwright.RunOptions{
        Browsers: []string{"chromium"},
    }); err != nil {
        log.Fatal(err)
    }
	pw, err := playwright.Run()
	if err != nil {
        log.Fatal(err)
	}
    defer pw.Stop()
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
        Headless: playwright.Bool(false),
    })
	if err != nil {
        log.Fatal(err)
	}
    defer browser.Close()

    page, err := browser.NewPage(playwright.BrowserNewPageOptions{})
	if err != nil {
        log.Fatal(err)
	}
    if _, err := page.Goto("https://yahoo.co.jp"); err != nil {
        log.Fatal(err)
    }
    _, err = page.PDF(playwright.PagePdfOptions{
        Path: playwright.String("aa.pdf"),
    })
    if err != nil {
        log.Fatal(err)
    }

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
