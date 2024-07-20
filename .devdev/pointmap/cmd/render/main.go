package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
)

//go:embed svg.tmpl
var svgTmpl string

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

func main() {
	tmpl, err := template.New("svg").Parse(svgTmpl)
	if err != nil {
		panic(err)
	}

	rects := []Rect{
		{X: 10, Y: 10, Width: 100, Height: 50},
		{X: 150, Y: 50, Width: 120, Height: 70},
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, rects); err != nil {
		panic(err)
	}
	fmt.Printf("%s", out.String())
}
