package usecase

import (
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/playwright-community/playwright-go"
)

func ExportPdf() error {
	if err := playwright.Install(); err != nil {
		return err
	}

	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		return err
	}
	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		Screen: &playwright.Size{
			Width: 1000,
			Height: 500,
		},
	})
	if err != nil {
		return err
	}
	_, err = page.Goto("https://example.com")
	if err != nil {
		return err
	}

	fbytes, err := page.PDF(playwright.PagePdfOptions{
		Width: playwright.String("1000"),
		Height: playwright.String("500"),
	})
	if err != nil {
		return err
	}
	f, err := os.Create("a.pdf")
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(fbytes)

	// merge pdf files
	if err := api.MergeCreateFile([]string{"a.pdf", "a.pdf"}, "o.pdf", false, nil); err != nil {
		return err
	}

	return nil
}
