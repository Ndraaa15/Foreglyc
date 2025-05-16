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
	TotalCalory       int64           `json:"totalCalory"`
	TotalCarbohydrate int64           `json:"totalCarbohydrate"`
	TotalFat          int64           `json:"totalFat"`
	TotalProtein      int64           `json:"totalProtein"`
	GlyecemicIndex    int64           `json:"glyecemicIndex"`
}

type FoodMonitoringResponse struct {
	Id                int64           `json:"id"`
	FoodName          string          `json:"foodName"`
	MealTime          string          `json:"mealTime"`
	ImageUrl          string          `json:"imageUrl"`
	Nutritions        json.RawMessage `json:"nutritions"`
	TotalCalory       int64           `json:"totalCalory"`
	TotalCarbohydrate int64           `json:"totalCarbohydrate"`
	TotalFat          int64           `json:"totalFat"`
	TotalProtein      int64           `json:"totalProtein"`
	GlyecemicIndex    int64           `json:"glyecemicIndex"`
}

type CountFoodMonitoringFilter struct {
	UserId string
	Time   time.Time
}

type GetFoodMonitoringFilter struct {
	UserId string
	Date   time.Time
}

type DailyFoodResponse struct {
	MealTime          string                     `json:"mealTime"`
	Time              string                     `json:"time"`
	FoodMonitoring    FoodMonitoringResponse     `json:"foodMonitoring,omitempty"`
	FoodRecomendation FoodRecommendationResponse `json:"foodRecomendation,omitempty"`
}

type FoodHomepageResponse struct {
	DailyFoodResponses []DailyFoodResponse `json:"dailyFoodResponses"`
	TotalCalory        int64               `json:"totalCalory"`
	TotalCarbohydrate  int64               `json:"totalCarbohydrate"`
	TotalFat           int64               `json:"totalFat"`
	TotalProtein       int64               `json:"totalProtein"`
	TotalInsuline      int64               `json:"totalInsuline"`
}
