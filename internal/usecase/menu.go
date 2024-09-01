package usecase

import (
	"fmt"
	"os"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/getlantern/systray"
	"github.com/pkg/browser"
)

func LaunchMenu() error {
	systray.Run(startMenu, endMenu)

	return nil
}

func startMenu() {
	systray.SetTitle("CodeTrailer")

	captureBtn := systray.AddMenuItem("Capture", "Capture")
	go func() {
		for {
			<-captureBtn.ClickedCh
			fmt.Println("capture")
		}
	}()

	recordBtn := systray.AddMenuItem("Record", "Record")
	go func() {
		for {
			<-recordBtn.ClickedCh
			fmt.Println("record")

			repos := repository.New()
			priority, _ := repos.Log.Ask("Priority", "1")
			fmt.Printf("%s\n", priority)
			name, _ := repos.Log.Ask("Name", "")
			fmt.Printf("%s\n", name)
			description, _ := repos.Log.Ask("Description", "")
			fmt.Printf("%s\n", description)
		}
	}()

	previewBtn := systray.AddMenuItem("Preview", "Preview")
	go func() {
		for {
			<-previewBtn.ClickedCh
			fmt.Println("preview")

			browser.OpenURL("https://example.com")
		}
	}()

	exportPdfBtn := systray.AddMenuItem("ExportPdf", "ExportPdf")
	go func() {
		for {
			<-exportPdfBtn.ClickedCh
			fmt.Println("exportPdf")
		}
	}()

	quitBtn := systray.AddMenuItem("Quit", "Quit")
	go func() {
		for {
			<-quitBtn.ClickedCh
			os.Exit(0)
		}
	}()
}

func endMenu() {}
