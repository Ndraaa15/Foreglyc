package dto

import "encoding/json"

type FoodMonitoringRequest struct {
	FoodName          string          `json:"foodName"`
	TimeType          string          `json:"timeType"`
	ImageUrl          string          `json:"imageUrl"`
	Nutritions        json.RawMessage `json:"nutritions"`
	TotalCalory       int             `json:"totalCalory"`
	TotalCarbohydrate int             `json:"totalCarbohydrate"`
	TotalFat          int             `json:"totalFat"`
	TotalProtein      int             `json:"totalProtein"`
}

type FoodMonitoringResponse struct {
	Id                int64           `json:"id"`
	FoodName          string          `json:"foodName"`
	TimeType          string          `json:"timeType"`
	ImageUrl          string          `json:"imageUrl"`
	Nutritions        json.RawMessage `json:"nutritions"`
	TotalCalory       int             `json:"totalCalory"`
	TotalCarbohydrate int             `json:"totalCarbohydrate"`
	TotalFat          int             `json:"totalFat"`
	TotalProtein      int             `json:"totalProtein"`
}
