package entity

import (
	"time"

	"github.com/lib/pq"
)

type GlucometerMonitoringPreference struct {
	UserId                        string         `db:"user_id" `
	StartWakeUpTime               time.Time      `db:"start_wake_up_time"`
	EndWakeUpTime                 time.Time      `db:"end_wake_up_time"`
	PhysicalActivityDays          pq.StringArray `db:"physical_activity_days"`
	StartSleepTime                time.Time      `db:"start_sleep_time"`
	EndSleepTime                  time.Time      `db:"end_sleep_time"`
	HypoglycemiaAccuteThreshold   float64        `db:"hypoglycemia_accute_threshold"`
	HypoglycemiaChronicThreshold  float64        `db:"hypoglycemia_chronic_threshold"`
	HyperglycemiaAccuteThreshold  float64        `db:"hyperglycemia_accute_threshold"`
	HyperglycemiaChronicThreshold float64        `db:"hyperglycemia_chronic_threshold"`
	SendNotification              bool           `db:"send_notification"`
	CreatedAt                     pq.NullTime    `db:"created_at"`
	UpdatedAt                     pq.NullTime    `db:"updated_at"`
}
