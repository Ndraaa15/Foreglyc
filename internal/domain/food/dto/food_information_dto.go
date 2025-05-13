package dto

type CreateFoodInformationRequest struct {
	TimeType string `json:"TimeType"`
	ImageUrl string `json:"imageUrl"`
}

type FoodInformationResponse struct {
	FoodName    string                   `json:"foodName"`
	TimeType    string                   `json:"timeType"`
	ImageUrl    string                   `json:"imageUrl"`
	Nutrition   []NutritionGroupResponse `json:"nutritions"`
	TotalCalory int                      `json:"totalCalory"`
}

type NutritionGroupResponse struct {
	Type       string              `json:"type"`
	Components []ComponentResponse `json:"components"`
}

type ComponentResponse struct {
	Name    string  `json:"name"`
	Portion float64 `json:"portion"`
	Unit    string  `json:"unit"`
	Calory  int     `json:"calory"`
}
