package database

import (
	"fmt"
	"math/rand"

	"github.com/bxcodec/faker/v4"
	"github.com/rodrigospano/hotel/internal/data"
	"gorm.io/gorm"
)

func RoomsSeed(db *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		room := data.Room{
			Room_name:       fmt.Sprintf("%s-%d", faker.Word(), rand.Intn(50)),
			Price_per_night: rand.Float64() * 10,
			Description:     faker.Paragraph(),
			Beds:            rand.Intn(3),
		}
		if err := db.Save(&room); err != nil {
			fmt.Printf("Error while seeding rooms: %s", room.Room_name)
		}
	}
	return nil
}
