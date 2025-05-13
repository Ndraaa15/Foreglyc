package entity

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type DietaryPlan struct {
	UserID                 string          `db:"user_id"`
	LiveWith               string          `db:"live_with"`
	BreakfastTime          time.Time       `db:"breakfast_time"`
	LunchTime              time.Time       `db:"lunch_time"`
	DinnerTime             time.Time       `db:"dinner_time"`
	MorningSnackTime       time.Time       `db:"morning_snack_time"`
	AfternoonSnackTime     time.Time       `db:"afternoon_snack_time"`
	IsUseInsuline          bool            `db:"is_use_insuline"`
	InsuliseQuestionnaires json.RawMessage `db:"insulise_questionnaires" type:"jsonb"`
	TotalDailyInsuline     float64         `db:"total_daily_insuline"`
	MealPlanType           string          `db:"meal_plan_type"`
	CreatedAt              pq.NullTime     `db:"created_at"`
	UpdatedAt              pq.NullTime     `db:"updated_at"`
}
