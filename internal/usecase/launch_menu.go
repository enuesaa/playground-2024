package usecase

import (
	"fmt"

	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/getlantern/systray"
)

func LaunchMenu(repos repository.Repos, path string) {
	onready := func ()  {
		systray.SetTitle("CodeTrailer")

		captureBtn := systray.AddMenuItem("Capture", "Capture")	
		go func() {
			<-captureBtn.ClickedCh
			capturePath := fmt.Sprintf("%s/aa.png", path)
			Capture(repos, capturePath)
			systray.Quit()
		}()
	}
	onexit := func() {}

	systray.Run(onready, onexit)
}
