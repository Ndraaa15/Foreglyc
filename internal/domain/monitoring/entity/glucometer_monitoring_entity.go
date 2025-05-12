package entity

import (
	"github.com/Ndraaa15/foreglyc-server/pkg/enum"
	"github.com/lib/pq"
)

type GlucometerMonitoring struct {
	Id           int64              `db:"id"`
	UserId       string             `db:"user_id"`
	BloodGlucose float64            `db:"blood_glucose"`
	Unit         string             `db:"unit"`
	Status       enum.GlucoseStatus `db:"status"`
	CreatedAt    pq.NullTime        `db:"created_at"`
	UpdatedAt    pq.NullTime        `db:"updated_at"`
}
