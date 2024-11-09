package main

import (
	"bytes"
	_ "embed"
	"log"
	"text/template"

	"github.com/gofiber/fiber/v3"
)

//go:embed downloader.sh.tmpl
var downloader string

type DownloaderVars struct {
	Host string
}

func main() {
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
		host := c.Get("Host")

		tmpl, err := template.New("downloader").Parse(downloader)
		if err != nil {
			return err
		}

		vars := DownloaderVars{
			Host: host,
		}
	
		var rendered bytes.Buffer
		if err := tmpl.Execute(&rendered, vars); err != nil {
			return err
		}
	
        return c.SendString(rendered.String())
    })

	app.Get("/archive.tar.gz", func (c fiber.Ctx) error {
		archive, err := ArchiveCurrentDir()
		if err != nil {
			return err
		}
		return c.Send(archive.Bytes())
	})

    if err := app.Listen(":80"); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}
