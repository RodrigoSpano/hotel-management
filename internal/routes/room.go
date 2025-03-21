package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigospano/hotel/internal/handlers"
)

func RoomsRouter(api *fiber.App) {
	roomsGroup := api.Group("/rooms")

	roomsGroup.Get("/", handlers.GetRooms)
	roomsGroup.Get("/:id", handlers.GetRoom)
	roomsGroup.Post("/", handlers.AddRoom)
	roomsGroup.Patch("/:id", handlers.UpdateRoom)
	roomsGroup.Delete("/:id", handlers.DeleteRoom)
}
