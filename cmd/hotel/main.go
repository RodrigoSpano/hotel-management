package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigospano/hotel/internal/routes"
)

func main() {
	app := fiber.New()

	api := app.Group("/api/v1")

	// routes
	routes.RoomsRouter(api.Group("/rooms"))

	app.Listen(":4000")
}
