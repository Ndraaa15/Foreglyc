package dto

type ChatMessageRequest struct {
	Role    string `json:"role" validate:"required,oneof=user model"`
	Message string `json:"message" validate:"required"`
	FileUrl string `json:"fileUrl"`
}

type ChatMessageResponse struct {
	Role    string `json:"role"`
	Message string `json:"message"`
	FileUrl string `json:"fileUrl"`
}

type ChatResponse struct {
	History []ChatMessageResponse `json:"history"`
}
