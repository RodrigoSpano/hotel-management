package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	Id              primitive.ObjectID `json:"id,omitempty"`
	Guest           string             `json:"guest"`
	Price_per_night float64            `json:"price_per_night"`
	Beds            int                `json:"beds"`
}
