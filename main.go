package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})
	pdf.AddPage()
	if err := pdf.AddTTFFont("HackNerdFont", "./HackNerdFont-Regular.ttf"); err != nil {
		panic(err)
	}
	if err := pdf.SetFont("HackNerdFont", "", 14); err != nil {
		panic(err)
	}
	pdf.Cell(nil, "aa")
	pdf.WritePdf("hello.pdf")

}

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
