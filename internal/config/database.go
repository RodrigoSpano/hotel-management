package database

import (
	"fmt"
	"os"

	"github.com/rodrigospano/hotel/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

func ConnectDB() error {
	var err error
	Database, err := gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Room{})
	return nil
}
