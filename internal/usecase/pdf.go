package usecase

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func Pdf() {
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
}
