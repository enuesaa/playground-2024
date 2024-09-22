package usecase

import (
	"fmt"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/getlantern/systray"
)

func LaunchMenu(repos repository.Repos, path string) error {
	onready := func() {
		systray.SetTitle("CodeTrailer")

		captureBtn := systray.AddMenuItem("Capture", "Capture")
		go func() {
			for {
				<-captureBtn.ClickedCh
				capturePath := fmt.Sprintf("%s/aa.png", path)
				if err := Capture(repos, capturePath); err != nil {
					fmt.Printf("Error: %s\n", err.Error())
				}
			}
		}()
	}
	onexit := func() {}

	systray.Run(onready, onexit)

	return nil
}
