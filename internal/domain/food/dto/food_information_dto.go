package dto

type CreateFoodInformationRequest struct {
	MealTime string `json:"mealTime"`
	ImageUrl string `json:"imageUrl"`
}

type FoodInformationResponse struct {
	FoodName           string                   `json:"foodName"`
	MealTime           string                   `json:"mealTime"`
	ImageUrl           string                   `json:"imageUrl"`
	Nutrition          []NutritionGroupResponse `json:"nutritions"`
	TotalCalory        int64                    `json:"totalCalory"`
	TotalCarbohydrates int64                    `json:"totalCarbohydrate"`
	TotalProtein       int64                    `json:"totalProtein"`
	TotalFat           int64                    `json:"totalFat"`
	GlyecemicIndex     int64                    `json:"glyecemicIndex"`
}

type NutritionGroupResponse struct {
	Type       string              `json:"type"`
	Components []ComponentResponse `json:"components"`
}

type ComponentResponse struct {
	Name    string  `json:"name"`
	Portion float64 `json:"portion"`
	Unit    string  `json:"unit"`
	Calory  int64   `json:"calory"`
}
