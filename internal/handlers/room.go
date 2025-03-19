package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/rodrigospano/hotel/internal/entities"
	"gorm.io/gorm/config"
)

func getRooms(c *fiber.Ctx) error {
	var rooms *[]entities.Room
	config.Database.Find(&rooms)
	return c.Status(200).JSON(rooms)
}
