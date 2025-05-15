package dto

import (
	"encoding/json"
	"time"
)

type FoodMonitoringRequest struct {
	FoodName          string          `json:"foodName"`
	MealTime          string          `json:"mealTime"`
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
	MealTime          string          `json:"mealTime"`
	ImageUrl          string          `json:"imageUrl"`
	Nutritions        json.RawMessage `json:"nutritions"`
	TotalCalory       int             `json:"totalCalory"`
	TotalCarbohydrate int             `json:"totalCarbohydrate"`
	TotalFat          int             `json:"totalFat"`
	TotalProtein      int             `json:"totalProtein"`
}

type CountFoodMonitoringFilter struct {
	UserId string
	Time   time.Time
}

type GetFoodMonitoringFilter struct {
	UserId string
	Date   time.Time
}
