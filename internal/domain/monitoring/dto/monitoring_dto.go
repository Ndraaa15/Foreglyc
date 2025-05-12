package dto

import "time"

type CreateGlucometerMonitoringRequest struct {
	BloodGlucose float64 `json:"bloodGlucose" validate:"required"`
}

type GetGlucometerMonitoringFilter struct {
	UserId    string
	StartDate time.Time
	EndDate   time.Time
}

type GetGlucometerMonitoringGraphFilter struct {
	UserId string
	Type   string `query:"type" validate:"required,oneof=hourly daily"`
}

type GlucometerMonitoringResponse struct {
	BloodGlucose float64 `json:"bloodGlucose"`
	Status       string  `json:"status"`
	Unit         string  `json:"unit"`
	Date         string  `json:"date"`
	Time         string  `json:"time"`
	IsSafe       bool    `json:"isSafe"`
}

type GlucometerMonitoringGraphResponse struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}
