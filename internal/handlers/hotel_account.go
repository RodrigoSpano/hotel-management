package handlers

import (
	"fmt"

	database "github.com/rodrigospano/hotel/internal/config"
	"github.com/rodrigospano/hotel/internal/data"
)

func CreateAccount(account data.HotelAccount) error {
	// call here logic of findACcount to verify if exists
	// it is creating this in the wrong table
	result := database.DB.Table("hotel_accounts").Create(&account)
	if result.Error != nil {
		return result.Error
	}
	fmt.Print(result)
	return nil
}
