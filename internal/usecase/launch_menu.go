package usecase

import (
	"github.com/enuesaa/codetrailer/internal/repository"
	"github.com/getlantern/systray"
)

func LaunchMenu(repos repository.Repos) {
	onready := func ()  {
		systray.SetTitle("CodeTrailer")

		captureBtn := systray.AddMenuItem("Capture", "Capture")
		go func() {
			for {
				<-captureBtn.ClickedCh
				Capture(repos)
				systray.Quit()
			}
		}()
	}
	onexit := func() {}

	systray.Run(onready, onexit)
}
