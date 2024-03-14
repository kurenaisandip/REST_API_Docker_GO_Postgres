package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kurenaisandip/REST_API_Docker_GO_Postgres/handlers"
)

func setupRoutes(app *fiber.App) {

	// app.Get("/", handlers.Home)

	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)

}
