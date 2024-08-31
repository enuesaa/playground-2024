package main

import (
	"github.com/getlantern/systray"
)

func startMenu() {
	systray.SetTitle("Hi")
	systray.SetTooltip("Hello")
	systray.AddMenuItem("Quit", "Quit")
}

func endMenu() {}
