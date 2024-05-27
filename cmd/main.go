package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		log.Info("Receive Request")
		return c.SendString("WORK")
	})

	app.Post("/send", func(c *fiber.Ctx) error {
		body := c.Body()

		log.Info(string(body))
		return nil
	})

	app.Listen(":1234")
}
