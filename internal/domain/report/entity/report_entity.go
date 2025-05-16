package entity

import "github.com/lib/pq"

type ReportInformation struct {
	Id                         int64       `db:"id"`
	UserId                     string      `db:"user_id"`
	TotalBloodGlucose          int         `db:"total_blood_glucose"`
	RecommendationBloodGlucose string      `db:"recommendation_blood_glucose"`
	Recommendation             string      `db:"recommendation"`
	CreatedAt                  pq.NullTime `db:"created_at"`
	UpdatedAt                  pq.NullTime `db:"updated_at"`
}
