package data

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Room_name       string  `json:"room_name"`
	Price_per_night float64 `json:"price_per_night"`
	Beds            int     `json:"beds"`
	Description     string  `json:"description"`
}
