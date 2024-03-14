package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kurenaisandip/REST_API_Docker_GO_Postgres/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
