package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rodrigospano/hotel/internal/routes"
)

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api/v1", micro)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	// routes
	routes.RoomsRouter(micro.Group("/rooms"))

	app.Listen(":4000")
}
