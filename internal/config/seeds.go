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
		if result := db.Save(&room); result.Error != nil {
			fmt.Println(result.Error.Error())
			fmt.Printf("Error while seeding rooms: %s", room.Room_name)
		}
	}
	fmt.Println("Rooms table seeded successfully with fake data")
	return nil
}
