package main

import (
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/playwright-community/playwright-go"
)

func main() {
	if err := playwright.Install(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		Screen: &playwright.Size{
			Width: 1000,
			Height: 500,
		},
	})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	_, err = page.Goto("https://example.com")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	fbytes, err := page.PDF(playwright.PagePdfOptions{
		Width: playwright.String("1000"),
		Height: playwright.String("500"),
	})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	f, err := os.Create("a.pdf")
	defer f.Close()
	f.Write(fbytes)

	// merge pdf files
	if err := api.MergeCreateFile([]string{"a.pdf", "a.pdf"}, "o.pdf", false, nil); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
