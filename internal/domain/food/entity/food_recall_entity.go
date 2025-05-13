package entity

import (
	"encoding/json"

	"github.com/lib/pq"
)

type FoodRecall struct {
	Id            int64           `db:"id"`
	UserID        string          `db:"user_id"`
	FoodName      string          `db:"food_name"`
	TimeType      string          `db:"time_type"`
	ImageUrl      string          `db:"image_url"`
	Nutritions    json.RawMessage `db:"nutritions"`
	TotalCalories int             `db:"total_calories"`
	CreatedAt     pq.NullTime     `db:"created_at"`
	UpdatedAt     pq.NullTime     `db:"updated_at"`
}
