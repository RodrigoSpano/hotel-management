package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(api *fiber.App) {
	// room routes /////////
	RoomsRouter(api)
	HotelAccountRouter(api)
}
