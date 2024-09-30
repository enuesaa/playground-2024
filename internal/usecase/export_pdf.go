package usecase

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/playwright-community/playwright-go"
)

func ExportPdf(repos repository.Repos) error {
	fbytes, err := repos.Fs.Read(repos.Config.DocPath)
	if err != nil {
		return err
	}
	readme := string(fbytes)
	slides := strings.Split(readme, "\n\n---\n\n")
	slidesCount := len(slides)


	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch()
	if err != nil {
		return err
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		return err
	}
	defer page.Close()

	_, err = page.Goto("http://localhost:3000?preview")
	if err != nil {
		return err
	}

	tmpfiles := []string{}
	for i := range slidesCount {
		err = page.Keyboard().Down("ArrowRight")
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
		fbytes, err = page.PDF(playwright.PagePdfOptions{
			Width: playwright.String("1000"),
			Height: playwright.String("800"),
		})
		if err != nil {
			return err
		}

		filename := fmt.Sprintf("tmp-%d.pdf", i)
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		f.Write(fbytes)

		tmpfiles = append(tmpfiles, filename)
	}

	if err := api.MergeCreateFile(tmpfiles, "out.pdf", false, nil); err != nil {
		return err
	}

	for _, filename := range tmpfiles {
		if err := repos.Fs.Delete(filename); err != nil {
			return err
		}
	}

	return nil
}
