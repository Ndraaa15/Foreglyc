package entity

import (
	"time"

	"github.com/lib/pq"
)

type FoodRecommendation struct {
	Id                     int64       `db:"id"`
	UserId                 string      `db:"user_id"`
	FoodName               string      `db:"food_name"`
	MealTime               string      `db:"meal_time"`
	Ingredients            string      `db:"ingredients"`
	CaloriesPerIngredients string      `db:"calories_per_ingredients"`
	ImageUrl               string      `db:"image_url"`
	TotalCalory            int         `db:"total_calory"`
	GlycemicIndex          int         `db:"glycemic_index"`
	Date                   time.Time   `db:"date"`
	CreatedAt              pq.NullTime `db:"created_at"`
	UpdatedAt              pq.NullTime `db:"updated_at"`
}
