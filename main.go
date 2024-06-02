package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} ${reqHeader:X-GitHub-Event} ${body}\n",
	}))

	// webhook の送信先
	app.Post("/", func(c fiber.Ctx) error {
		return nil
	})

	app.Listen(":3000")
}
