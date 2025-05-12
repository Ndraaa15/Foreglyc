package dto

type CreateQuestionnaireItem struct {
	Question string `json:"question" validate:"required"`
	Answer   any    `json:"answer" validate:"required"`
}

type CreateMonitoringQuestionnaire struct {
	GlucometerMonitoringID int64                     `json:"glucometerMonitoringId" validate:"required"`
	Questionnaire          []CreateQuestionnaireItem `json:"questionnaire" validate:"required,dive,required"`
	ManagementType         string                    `json:"managementType"`
}

type MonitoringQuestionnaireResponse struct {
	Message string `json:"message"`
}
