package routes

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	db_mongo "github.com/rodrigospano/hotel/internal/database"
)

var roomCollection = db_mongo.GetCollection(db_mongo.DB, "rooms")

func RoomsRouter(app fiber.Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		// rooms := []models.Room{}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		fmt.Println(roomCollection.Find(ctx, nil))
		// return c.Status(200).JSON()
		return c.SendStatus(200)
	})

}
