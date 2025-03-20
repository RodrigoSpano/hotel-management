package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(api fiber.Router) {
	// room routes /////////
	RoomsRouter(api.Group("/rooms"))

}
