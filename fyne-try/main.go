package main

import (
	"fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := fyneapp.New()

	window := app.NewWindow("Hello")
	window.Resize(fyne.NewSize(1000, 800))

	// label
	helloLabel := widget.NewLabel("Hello Fyne!")

	// textarea
	textarea := widget.NewMultiLineEntry()
	textarea.SetPlaceHolder("text")

	// button
	hiButton := widget.NewButton("Hi!", func() {
		helloLabel.SetText("Welcome!")
	})

	window.SetContent(container.NewVBox(
		textarea,
		hiButton,
		helloLabel,
	))

	window.ShowAndRun()
}
