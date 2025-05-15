package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
)

func FoodRecommendationToResponse(foodRecomendation *entity.FoodRecommendation) dto.FoodRecommendationResponse {
	return dto.FoodRecommendationResponse{
		Id:       foodRecomendation.Id,
		FoodName: foodRecomendation.FoodName,
		MealTime: foodRecomendation.MealTime,
	}
}
