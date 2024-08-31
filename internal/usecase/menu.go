package usecase

import (
	"github.com/getlantern/systray"
)

func LaunchMenu() error {
	systray.Run(startMenu, endMenu)

	return nil
}

func startMenu() {
	systray.SetTitle("Hi")
	systray.SetTooltip("Hello")
	systray.AddMenuItem("Quit", "Quit")
}

func endMenu() {}
