package dto

import "time"

type GlucosePredictionResponse struct {
	Time  time.Time `json:"time"`
	Value int       `json:"value"`
}

type ScenarioResponse struct {
	Type            string                      `json:"type"`
	Prediction      []GlucosePredictionResponse `json:"prediction"`
	Reason          string                      `json:"reason"`
	Recommendations []string                    `json:"recommendations"`
}

type PredictionResponse struct {
	Scenario []ScenarioResponse    `json:"scenario"`
	Chats    []ChatMessageResponse `json:"chats"`
}
