package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello")
    })
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
