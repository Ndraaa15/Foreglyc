package dto

import "time"

type GetReportInformationFilter struct {
	UserId    string
	StartDate time.Time
	EndDate   time.Time
}

type MonthResponse struct {
	Year  int
	Month string
}

type ListMonthResponse struct {
	LastUpdated string          `json:"lastUpdated"`
	Months      []MonthResponse `json:"months"`
}

type ReportInformationResponse struct {
	Id                         int64  `json:"id"`
	Date                       string `json:"date"`
	TotalBloodGlucose          int    `json:"totalBloodGlucose"`
	RecommendationBloodGlucose string `json:"recommendationBloodGlucose"`
	Recommendation             string `json:"recommendation"`
}
