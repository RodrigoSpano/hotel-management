package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/rodrigospano/hotel/internal/types"
)

func ConnectDB() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("hotel.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)
	DB.AutoMigrate(&types.Room{})

	return DB
}

var DB = *ConnectDB()
