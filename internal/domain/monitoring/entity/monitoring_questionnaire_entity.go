package entity

import (
	"github.com/lib/pq"
)

type QuestionnaireItem struct {
	Question string      `json:"question"`
	Answer   interface{} `json:"answer"`
}

type MonitoringQuestionnaire struct {
	Id                     int64               `db:"id"`
	GlucometerMonitoringID string              `db:"glucometer_monitoring_id"`
	Questionnaires         []QuestionnaireItem `db:"questionnaires,type:jsonb"`
	ManagementType         *string             `db:"management_type"`
	CreatedAt              pq.NullTime         `db:"created_at"`
	UpdatedAt              pq.NullTime         `db:"updated_at"`
}
