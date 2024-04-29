package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var cli struct {
	Port int `help:"port to serve." default:"3000"`
}

func main() {
	kong.Parse(&cli)
	fmt.Println(cli.Port)
}
