package main

import (
	"log"

    "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Post("/hook", func(c fiber.Ctx) error {
		log.Printf("%+v", c.Body())
		return nil
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
