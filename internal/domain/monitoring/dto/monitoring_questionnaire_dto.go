package dto

import "encoding/json"

type CreateMonitoringQuestionnaire struct {
	GlucometerMonitoringID int64           `json:"glucometerMonitoringId" validate:"required"`
	Questionnaires         json.RawMessage `json:"questionnaires" validate:"required,dive,required"`
	ManagementType         string          `json:"managementType"`
}

type MonitoringQuestionnaireResponse struct {
	Message string `json:"message"`
}
