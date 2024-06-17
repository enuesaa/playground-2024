package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/playwright-community/playwright-go"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetHomeDir() (string, error) {
	return os.UserConfigDir()
}

func (a *App) MakeSomeDir() error {
	homedir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	path := filepath.Join(homedir, ".wailsapp")
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil
	}
	fpath := filepath.Join(path, "config.json")
	f, err := os.Create(fpath)
    if err != nil {
        return err
    }
    defer f.Close()

    if _, err = f.WriteString(`{"a":"b"}`); err != nil {
        return err
    }
	return nil
}

// Greet returns a greeting for the given name
func (a *App) Greet2(name string) string {
	if err := playwright.Install(); err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	// see https://playwright-community.github.io/playwright-go/
	pw, err := playwright.Run()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	browser, err := pw.Firefox.Launch()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	page, err := browser.NewPage()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	entries, err := page.Locator(".athing").All()
	if err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	text := ""
	if len(entries) > 0 {
		entry := entries[0]
		title, err := entry.Locator("td.title > span > a").TextContent()
		if err != nil {
			return fmt.Sprintf("Error: %s", err.Error())
		}
		text = title
	}
	if err = browser.Close(); err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	if err = pw.Stop(); err != nil {
		return fmt.Sprintf("Error: %s", err.Error())
	}
	return text
}
