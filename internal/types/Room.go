package types

type Room struct {
	Id              int     `json:"id"`
	Guest           string  `json:"guest"`
	Price_per_night float64 `json:"price_per_night"`
	Beds            int     `json:"beds"`
}
