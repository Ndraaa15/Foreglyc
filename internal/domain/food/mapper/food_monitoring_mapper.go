package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
)

func FoodMonitoringToResponse(data *entity.FoodMonitoring) dto.FoodMonitoringResponse {
	return dto.FoodMonitoringResponse{
		Id:                data.Id,
		FoodName:          data.FoodName,
		MealTime:          data.MealTime,
		ImageUrl:          data.ImageUrl,
		Nutritions:        data.Nutritions,
		TotalCalory:       data.TotalCalory,
		TotalCarbohydrate: data.TotalCarbohydrate,
		TotalFat:          data.TotalFat,
		TotalProtein:      data.TotalProtein,
	}
}
