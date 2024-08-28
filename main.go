package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func main() {
	// start up prompt here.
	// - write contents
	// - capture screenshot
	// struct に詰める
	// create pdf file

    app := &cli.App{
        Name:  "codetrailer",
        Usage: "",
        Action: func(*cli.Context) error {
            fmt.Println("hello")
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
