package usecase

import (
	"fmt"

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
}

func endMenu() {}
