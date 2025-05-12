package dto

type CreateGlucometerMonitoringPreferenceRequest struct {
	StartWakeUpTime               string   `json:"startWakeUpTime" validate:"required"`
	EndWakeUpTime                 string   `json:"endWakeUpTime" validate:"required"`
	PhysicalActivityDays          []string `json:"physicalActivityDays" validate:"required"`
	StartSleepTime                string   `json:"startSleepTime" validate:"required"`
	EndSleepTime                  string   `json:"endSleepTime" validate:"required"`
	HypoglycemiaAccuteThreshold   float64  `json:"hypoglycemiaAccuteThreshold" validate:"required"`
	HypoglycemiaChronicThreshold  float64  `json:"hypoglycemiaChronicThreshold" validate:"required"`
	HyperglycemiaAccuteThreshold  float64  `json:"hyperglycemiaAccuteThreshold" validate:"required"`
	HyperglycemiaChronicThreshold float64  `json:"hyperglycemiaChronicThreshold" validate:"required"`
	SendNotification              *bool    `json:"sendNotification" validate:"required"`
}

type CreateCGMMonitoringPreferenceRequest struct {
	PhysicalActivityDays          []string `json:"physicalActivityDays" validate:"required"`
	HypoglycemiaAccuteThreshold   float64  `json:"hypoglycemiaAccuteThreshold" validate:"required"`
	HypoglycemiaChronicThreshold  float64  `json:"hypoglycemiaChronicThreshold" validate:"required"`
	HyperglycemiaAccuteThreshold  float64  `json:"hyperglycemiaAccuteThreshold" validate:"required"`
	HyperglycemiaChronicThreshold float64  `json:"hyperglycemiaChronicThreshold" validate:"required"`
	SendNotification              *bool    `json:"sendNotification" validate:"required"`
}

type GlucometerMonitoringPrefereceResponse struct {
	UserId                        string   `json:"userId"`
	StartWakeUpTime               string   `json:"startWakeUpTime"`
	EndWakeUpTime                 string   `json:"endWakeUpTime"`
	PhysicalActivityDays          []string `json:"physicalActivityDays"`
	StartSleepTime                string   `json:"startSleepTime"`
	EndSleepTime                  string   `json:"endSleepTime"`
	HypoglycemiaAccuteThreshold   float64  `json:"hypoglycemiaAccuteThreshold"`
	HypoglycemiaChronicThreshold  float64  `json:"hypoglycemiaChronicThreshold"`
	HyperglycemiaAccuteThreshold  float64  `json:"hyperglycemiaAccuteThreshold"`
	HyperglycemiaChronicThreshold float64  `json:"hyperglycemiaChronicThreshold"`
	SendNotification              bool     `json:"sendNotification"`
}

type CGMMonitoringPrefereceResponse struct {
	UserId                        string   `json:"userId"`
	PhysicalActivityDays          []string `json:"physicalActivityDays"`
	HypoglycemiaAccuteThreshold   float64  `json:"hypoglycemiaAccuteThreshold"`
	HypoglycemiaChronicThreshold  float64  `json:"hypoglycemiaChronicThreshold"`
	HyperglycemiaAccuteThreshold  float64  `json:"hyperglycemiaAccuteThreshold"`
	HyperglycemiaChronicThreshold float64  `json:"hyperglycemiaChronicThreshold"`
	SendNotification              bool     `json:"sendNotification"`
}
