package entity

import "github.com/lib/pq"

type DietaryInformation struct {
	UserId               string      `db:"user_id"`
	TotalCalory          int64       `db:"total_calory"`
	TotalBreakfastCalory int64       `db:"total_breakfast_calory"`
	TotalSnackCalory     int64       `db:"total_snack_calory"`
	TotalLunchCalory     int64       `db:"total_lunch_calory"`
	TotalDinnerCalory    int64       `db:"total_dinner_calory"`
	TotalCarbohydrate    int64       `db:"total_carbohydrate"`
	TotalFat             int64       `db:"total_fat"`
	TotalProtein         int64       `db:"total_protein"`
	CreatedAt            pq.NullTime `db:"created_at"`
	UpdatedAt            pq.NullTime `db:"updated_at"`
}
