package entity

import (
	"encoding/json"

	"github.com/lib/pq"
)

type FoodMonitoring struct {
	Id                int64           `db:"id"`
	UserID            string          `db:"user_id"`
	FoodName          string          `db:"food_name"`
	TimeType          string          `db:"time_type"`
	ImageUrl          string          `db:"image_url"`
	Nutritions        json.RawMessage `db:"nutritions"`
	TotalCalory       int             `db:"total_calory"`
	TotalCarbohydrate int             `db:"total_carbohydrate"`
	TotalProtein      int             `db:"total_protein"`
	TotalFat          int             `db:"total_fat"`
	CreatedAt         pq.NullTime     `db:"created_at"`
	UpdatedAt         pq.NullTime     `db:"updated_at"`
}
