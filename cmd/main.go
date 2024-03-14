package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kurenaisandip/REST_API_Docker_GO_Postgres/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
