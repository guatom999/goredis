package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		time.Sleep(10 * time.Millisecond)
		return c.SendString("Hello World")
	})
	app.Listen(":8000")

}
