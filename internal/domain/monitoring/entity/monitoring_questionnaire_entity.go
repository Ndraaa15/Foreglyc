package entity

import (
	"encoding/json"

	"github.com/lib/pq"
)

type MonitoringQuestionnaire struct {
	Id                     int64           `db:"id"`
	GlucometerMonitoringID string          `db:"glucometer_monitoring_id"`
	Questionnaires         json.RawMessage `db:"questionnaires,type:jsonb"`
	ManagementType         *string         `db:"management_type"`
	CreatedAt              pq.NullTime     `db:"created_at"`
	UpdatedAt              pq.NullTime     `db:"updated_at"`
}
