package main

import (
	_ "embed"
	"log"

	"github.com/dop251/goja"
)

//go:embed js/dist/index.js
var jsf string

func main() {
    runtime := goja.New()

	_, err := runtime.RunString(jsf)
	if err != nil {
		log.Panic(err)
	}
}
