package dto

import "time"

type MenuChatBotResponse struct {
	Date               string                              `json:"date"`
	FoodRecomendations []FoodRecommendationChatBotResponse `json:"foodRecomendations"`
}

type FoodRecommendationChatBotResponse struct {
	MealTime               string `json:"mealTime"`
	FoodName               string `json:"foodName"`
	Ingredients            string `json:"ingredients"`
	CaloriesPerIngredients string `json:"caloriesPerIngredients"`
	TotalCalories          int    `json:"totalCalories"`
	GlycemicIndex          int    `json:"glycemicIndex"`
	ImageUrl               string `json:"imageUrl"`
}

type MenuResponse struct {
	Date               string                       `json:"date"`
	FoodRecomendations []FoodRecommendationResponse `json:"foodRecomendations"`
}

type FoodRecommendationResponse struct {
	Id                     int64  `json:"id"`
	MealTime               string `json:"mealTime"`
	FoodName               string `json:"foodName"`
	Ingredients            string `json:"ingredients"`
	CaloriesPerIngredients string `json:"caloriesPerIngredients"`
	TotalCalories          int    `json:"totalCalories"`
	GlycemicIndex          int    `json:"glycemicIndex"`
	ImageUrl               string `json:"imageUrl"`
}

type CreateMenuRequest struct {
	Date               string                            `json:"date"`
	FoodRecomendations []CreateFoodRecommendationRequest `json:"foodRecomendations"`
}

type CreateFoodRecommendationRequest struct {
	MealTime               string `json:"mealTime"`
	FoodName               string `json:"foodName"`
	Ingredients            string `json:"ingredients"`
	CaloriesPerIngredients string `json:"caloriesPerIngredients"`
	TotalCalories          int    `json:"totalCalories"`
	GlycemicIndex          int    `json:"glycemicIndex"`
	ImageURL               string `json:"imageUrl"`
}

type GetFoodRecommendationFilter struct {
	UserId string
	Date   time.Time
}
