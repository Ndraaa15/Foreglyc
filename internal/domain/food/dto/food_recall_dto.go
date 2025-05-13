package dto

import "encoding/json"

type FoodRecallRequest struct {
	FoodName      string          `json:"foodName"`
	TimeType      string          `json:"timeType"`
	ImageUrl      string          `json:"imageUrl"`
	Nutritions    json.RawMessage `json:"nutritions"`
	TotalCalories int             `json:"totalCalory"`
}

type FoodRecallResponse struct {
	Id            int64           `json:"id"`
	FoodName      string          `json:"foodName"`
	TimeType      string          `json:"timeType"`
	ImageUrl      string          `json:"imageUrl"`
	Nutritions    json.RawMessage `json:"nutritions"`
	TotalCalories int             `json:"totalCalories"`
}
