package database

import (
	"fmt"
	"os"

	"github.com/rodrigospano/hotel/internal/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	var err error
	var (
		db_user     = os.Getenv("DB_USER")
		db_password = os.Getenv("DB_PASSWORD")
		db_host     = os.Getenv("DB_HOST")
		db_port     = os.Getenv("DB_PORT")
		db_name     = os.Getenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_password, db_host, db_port, db_name)

	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
	if err := DB.AutoMigrate(&data.Room{}, &data.HotelAccount{}); err != nil {
		panic(err)
	}

	// uncomment only if need seeds to be executed
	// RoomsSeed(DB, 40)
	// uncomment to drop a table
	// DropDbTable(DB, "rooms")
	return nil
}
