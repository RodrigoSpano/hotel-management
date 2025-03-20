package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	database "github.com/rodrigospano/hotel/internal/config"
	"github.com/rodrigospano/hotel/internal/data"
)

func GetRooms(c *fiber.Ctx) error {
	var rooms []data.Room
	result := database.DB.Find(&rooms)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(rooms)
}

// get room by id
func GetRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	var room data.Room

	result := database.DB.Find(&room, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(room)
}

func AddRoom(c *fiber.Ctx) error {
	room := new(data.Room)
	if err := c.BodyParser(room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	// todo => add validation if room with that name already exists
	result := database.DB.Create(room)
	fmt.Print(result)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": result.Error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(&room)
}

func DeleteRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	// todo => add validation if room exists
	result := database.DB.Delete(&data.Room{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": result.Error.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Room deleted successfully, id:" + id})
}
