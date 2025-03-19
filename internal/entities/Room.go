package entities

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	ID              uint    `json:"id" gorm:"primarykey"`
	Guest_name      string  `json:"guest_name"`
	Price_per_night float64 `json:"price_per_night"`
	Beds            int     `json:"beds"`
}
