package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
)

func FoodRecallToResponse(data *entity.FoodRecall) dto.FoodRecallResponse {
	return dto.FoodRecallResponse{
		Id:            data.Id,
		FoodName:      data.FoodName,
		TimeType:      data.TimeType,
		ImageUrl:      data.ImageUrl,
		Nutritions:    data.Nutritions,
		TotalCalories: data.TotalCalories,
	}
}
