package usecase

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
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

	quitBtn := systray.AddMenuItem("Quit", "Quit")
	go func() {
		for {
			<-quitBtn.ClickedCh
			os.Exit(0)
		}
	}()
}

func endMenu() {}
