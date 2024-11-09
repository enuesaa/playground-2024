package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
		archive, err := ArchiveCurrentDir()
		if err != nil {
			return err
		}
		return c.Send(archive.Bytes())
        // return c.SendString("Hello, World!")
    })

    if err := app.Listen(":80"); err != nil {
		log.Panicf("Error: %s", err.Error())
	}
}
