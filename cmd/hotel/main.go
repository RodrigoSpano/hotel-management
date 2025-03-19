package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rodrigospano/hotel/internal/routes"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load("../../.env")
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

	app.Listen(os.Getenv("PORT"))
}
