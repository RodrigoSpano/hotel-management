package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hotelProps struct {
	Name, Email, Phone string
}

type personInCharge struct {
	Fullname                       string `json:"fullname"`
	Phone                          string `json:"phone"`
	Email                          string `json:"email"`
	Gender                         string `json:"gender"`
	National_identification_number int32  `json:"national_identification_number"`
}

type address struct {
	Zip_code      string `json:"zip_code"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Place_address string `json:"place_address"`
}

type HotelAccount struct {
	gorm.Model
	ID               uuid.UUID      `gorm:"type:uuid;primaryKey:true"`
	Hotel            hotelProps     `gorm:"serializer:json"`
	Person_in_charge personInCharge `gorm:"serializer:json"`
	Address          address        `gorm:"serializer:json"`
}
