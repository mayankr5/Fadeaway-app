package main

import (
	"log"

	"github.com/fadeaway-app/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
