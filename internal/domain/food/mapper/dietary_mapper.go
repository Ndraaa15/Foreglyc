package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
)

func DietaryPlanToResponse(data *entity.DietaryPlan) dto.DietaryPlanResponse {
	return dto.DietaryPlanResponse{
		UserId:                 data.UserId,
		LiveWith:               data.LiveWith,
		BreakfastTime:          data.BreakfastTime.Format("15:04"),
		LunchTime:              data.LunchTime.Format("15:04"),
		DinnerTime:             data.DinnerTime.Format("15:04"),
		MorningSnackTime:       data.MorningSnackTime.Format("15:04"),
		AfternoonSnackTime:     data.AfternoonSnackTime.Format("15:04"),
		IsUseInsuline:          data.IsUseInsuline,
		InsuliseQuestionnaires: data.InsuliseQuestionnaires,
		TotalDailyInsuline:     data.TotalDailyInsuline,
		MealPlanType:           data.MealPlanType,
	}
}

func DietaryInformationToResponse(data *entity.DietaryInformation) dto.DietaryInformationResponse {
	return dto.DietaryInformationResponse{
		UserId:               data.UserId,
		TotalCalory:          data.TotalCalory,
		TotalBreakfastCalory: data.TotalBreakfastCalory,
		TotalSnackCalory:     data.TotalSnackCalory,
		TotalLunchCalory:     data.TotalLunchCalory,
		TotalDinnerCalory:    data.TotalDinnerCalory,
		TotalCarbohydrate:    data.TotalCarbohydrate,
		TotalFat:             data.TotalFat,
		TotalProtein:         data.TotalProtein,
	}
}
