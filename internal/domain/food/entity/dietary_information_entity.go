package entity

import "github.com/lib/pq"

type DietaryInformation struct {
	UserId               string      `db:"user_id"`
	TotalCalory          int         `db:"total_calory"`
	TotalBreakfastCalory int         `db:"total_breakfast_calory"`
	TotalSnackCalory     int         `db:"total_snack_calory"`
	TotalLunchCalory     int         `db:"total_lunch_calory"`
	TotalDinnerCalory    int         `db:"total_dinner_calory"`
	CreatedAt            pq.NullTime `db:"created_at"`
	UpdatedAt            pq.NullTime `db:"updated_at"`
}
