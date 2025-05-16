package entity

import (
	"encoding/json"

	"github.com/lib/pq"
)

type FoodMonitoring struct {
	Id                int64           `db:"id"`
	UserID            string          `db:"user_id"`
	FoodName          string          `db:"food_name"`
	MealTime          string          `db:"meal_time"`
	ImageUrl          string          `db:"image_url"`
	Nutritions        json.RawMessage `db:"nutritions"`
	TotalCalory       int64           `db:"total_calory"`
	TotalCarbohydrate int64           `db:"total_carbohydrate"`
	TotalProtein      int64           `db:"total_protein"`
	TotalFat          int64           `db:"total_fat"`
	GlyecemicIndex    int64           `db:"glycemic_index"`
	CreatedAt         pq.NullTime     `db:"created_at"`
	UpdatedAt         pq.NullTime     `db:"updated_at"`
}
