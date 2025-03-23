package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigospano/hotel/internal/controllers"
)

func HotelAccountRouter(api *fiber.App) {
	hotelAccountGroup := api.Group("/hotel_account")

	hotelAccountGroup.Post("/", controllers.AddRoom)
}
