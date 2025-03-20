package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	database "github.com/rodrigospano/hotel/internal/config"
	"github.com/rodrigospano/hotel/internal/routes"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load("../../.env")
	app := fiber.New()
	micro := fiber.New()

	// db initialization
	database.ConnectDB()

	client_url := os.Getenv("CLIENT_URL")
	app.Use(cors.New(cors.Config{
		AllowOrigins:     client_url,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	// routes
	app.Mount("/api/v1", micro)
	routes.InitializeRoutes(micro)

	app.Listen(os.Getenv("PORT"))
}
