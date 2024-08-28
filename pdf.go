package main

import (
	"github.com/signintech/gopdf"
)

func CreatePdf() {
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
