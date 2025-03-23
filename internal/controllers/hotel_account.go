package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrigospano/hotel/internal/data"
	"github.com/rodrigospano/hotel/internal/handlers"
)

func CreateAccount(c *fiber.Ctx) error {
	account := new(data.HotelAccount)
	if err := c.BodyParser(account); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	handlers.CreateAccount(*account)
	return c.SendStatus(200)
}
