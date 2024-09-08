package usecase

import (
	"github.com/playwright-community/playwright-go"
)

func ExportPdf() error {
	if err := playwright.Install(&playwright.RunOptions{
		Browsers: []string{"chromium"},
	}); err != nil {
		return err
	}
	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	defer pw.Stop()
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		return err
	}
	defer browser.Close()

	page, err := browser.NewPage(playwright.BrowserNewPageOptions{})
	if err != nil {
		return err
	}
	if _, err := page.Goto("https://yahoo.co.jp"); err != nil {
		return err
	}
	_, err = page.PDF(playwright.PagePdfOptions{
		Path: playwright.String("aa.pdf"),
	})
	if err != nil {
		return err
	}

	return nil
}
