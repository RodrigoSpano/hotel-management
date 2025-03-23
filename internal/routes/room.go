package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigospano/hotel/internal/controllers"
)

func RoomsRouter(api *fiber.App) {
	roomsGroup := api.Group("/rooms")

	roomsGroup.Get("/", controllers.GetRooms)
	roomsGroup.Get("/:id", controllers.GetRoom)
	roomsGroup.Post("/", controllers.AddRoom)
	roomsGroup.Patch("/:id", controllers.UpdateRoom)
	roomsGroup.Delete("/:id", controllers.DeleteRoom)
}
