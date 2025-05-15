package dto

import "encoding/json"

type CreateDietaryPlanRequest struct {
	LiveWith               string          `json:"liveWith" validate:"required"`
	BreakfastTime          string          `json:"breakfastTime" validate:"required"`
	LunchTime              string          `json:"lunchTime" validate:"required"`
	DinnerTime             string          `json:"dinnerTime" validate:"required"`
	MorningSnackTime       string          `json:"morningSnackTime" validate:"required"`
	AfternoonSnackTime     string          `json:"afternoonSnackTime" validate:"required"`
	IsUseInsuline          *bool           `json:"isUseInsuline" validate:"required"`
	InsuliseQuestionnaires json.RawMessage `json:"insuliseQuestionnaires"`
	TotalDailyInsuline     float64         `json:"totalDailyInsuline"`
	MealPlanType           string          `json:"mealPlanType"`
}

type UpdateInsulineQuestionnaireRequest struct {
	InsuliseQuestionnaires json.RawMessage `json:"insuliseQuestionnaires" validate:"required,dive,required"`
	TotalDailyInsuline     float64         `json:"totalDailyInsuline" validate:"required"`
}

type DietaryPlanResponse struct {
	UserId                 string          `json:"userId"`
	LiveWith               string          `json:"liveWith"`
	BreakfastTime          string          `json:"breakfastTime"`
	LunchTime              string          `json:"lunchTime"`
	DinnerTime             string          `json:"dinnerTime"`
	MorningSnackTime       string          `json:"morningSnackTime"`
	AfternoonSnackTime     string          `json:"afternoonSnackTime"`
	IsUseInsuline          bool            `json:"isUseInsuline"`
	InsuliseQuestionnaires json.RawMessage `json:"insuliseQuestionnaires"`
	TotalDailyInsuline     float64         `json:"totalDailyInsuline"`
	MealPlanType           string          `json:"mealPlanType"`
}

type DietaryInformationChatbotResponse struct {
	TotalCalory          int `json:"totalCalory"`
	TotalBreakfastCalory int `json:"totalBreakfastCalory"`
	TotalSnackCalory     int `json:"totalSnackCalory"`
	TotalLunchCalory     int `json:"totalLunchCalory"`
	TotalDinnerCalory    int `json:"totalDinnerCalory"`
}

type CreateDietaryInformationRequest struct {
	TotalCalory          int `json:"totalCalory"`
	TotalBreakfastCalory int `json:"totalBreakfastCalory"`
	TotalSnackCalory     int `json:"totalSnackCalory"`
	TotalLunchCalory     int `json:"totalLunchCalory"`
	TotalDinnerCalory    int `json:"totalDinnerCalory"`
}

type DietaryInformationResponse struct {
	UserId               string `json:"userId"`
	TotalCalory          int    `json:"totalCalory"`
	TotalBreakfastCalory int    `json:"totalBreakfastCalory"`
	TotalSnackCalory     int    `json:"totalSnackCalory"`
	TotalLunchCalory     int    `json:"totalLunchCalory"`
	TotalDinnerCalory    int    `json:"totalDinnerCalory"`
}
