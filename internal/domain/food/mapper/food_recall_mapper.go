package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
)

func FoodRecallToResponse(data *entity.FoodMonitoring) dto.FoodMonitoringResponse {
	return dto.FoodMonitoringResponse{
		Id:                data.Id,
		FoodName:          data.FoodName,
		TimeType:          data.TimeType,
		ImageUrl:          data.ImageUrl,
		Nutritions:        data.Nutritions,
		TotalCalory:       data.TotalCalory,
		TotalCarbohydrate: data.TotalCarbohydrate,
		TotalFat:          data.TotalFat,
		TotalProtein:      data.TotalProtein,
	}
}
