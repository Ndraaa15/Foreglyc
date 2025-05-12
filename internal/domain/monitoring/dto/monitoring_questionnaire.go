package dto

import "encoding/json"

type CreateMonitoringQuiestionnare struct {
	GlucareMonitoringId int64             `json:"glucareMonitoringId" validate:"required"`
	Quiestionnare       []json.RawMessage `json:"questionnare" validate:"required"`
}
