package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func Capture() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		filename := fmt.Sprintf("screenshot-%d.png", i)
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := png.Encode(f, img); err != nil {
			panic(err)
		}
	}
}