package dto

import (
	foodto "github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	monitoringdto "github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
)

type HomepageResponse struct {
	FullName                string                                            `json:"fullName"`
	PhotoProfile            string                                            `json:"photoProfile"`
	Level                   string                                            `json:"level"`
	DailyFoodResponses      []DailyFoodResponse                               `json:"dailyFoodResponses"`
	GlucoseMonitoringGraphs []monitoringdto.GlucometerMonitoringGraphResponse `json:"glucoseMonitoringGraphs"`
	TotalCalory             int                                               `json:"totalCalory"`
}

type DailyFoodResponse struct {
	MealTime          string                            `json:"mealTime"`
	Time              string                            `json:"time"`
	FoodMonitoring    foodto.FoodMonitoringResponse     `json:"foodMonitoring,omitempty"`
	FoodRecomendation foodto.FoodRecommendationResponse `json:"foodRecomendation,omitempty"`
}
