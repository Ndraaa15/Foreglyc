package entity

import "github.com/lib/pq"

type CGMMonitoringPreference struct {
	UserId                        string      `db:"user_id"`
	HypoglycemiaAccuteThreshold   float64     `db:"hypoglycemia_accute_threshold"`
	HypoglycemiaChronicThreshold  float64     `db:"hypoglycemia_chronic_threshold"`
	HyperglycemiaAccuteThreshold  float64     `db:"hyperglycemia_accute_threshold"`
	HyperglycemiaChronicThreshold float64     `db:"hyperglycemia_chronic_threshold"`
	SendNotification              bool        `db:"send_notification"`
	CreatedAt                     pq.NullTime `db:"created_at"`
	UpdatedAt                     pq.NullTime `db:"updated_at"`
}
