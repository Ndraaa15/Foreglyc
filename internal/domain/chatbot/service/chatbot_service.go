package service

import (
	"context"
	"encoding/json"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/chatbot/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/storage"
	"google.golang.org/genai"
)

func (c *ChatBotService) ChatForeglycExpert(ctx context.Context, requests []dto.ChatMessageRequest) ([]dto.ChatMessageResponse, error) {
	var contents []*genai.Content
	var history []dto.ChatMessageResponse

	var fileInformation storage.FileInformation

	for i, m := range requests {
		if m.Role == genai.RoleModel {
			contents = append(contents, &genai.Content{
				Role: genai.RoleModel,
				Parts: []*genai.Part{
					{Text: m.Message},
				},
			})
		} else if m.Role == genai.RoleUser {
			isLast := i == len(requests)-1

			var parts []*genai.Part
			if isLast && m.FileUrl != "" {
				var err error
				fileInformation, err = c.firebaseStorageService.GetFile(ctx, m.FileUrl)
				if err != nil {
					c.log.WithError(err).Error("failed to retrieve image")
					continue
				}
				parts = []*genai.Part{
					{Text: m.Message},
					{InlineData: &genai.Blob{Data: fileInformation.Data, MIMEType: fileInformation.Type}},
				}
			} else {
				parts = []*genai.Part{{Text: m.Message}}
			}

			contents = append(contents, &genai.Content{
				Role:  genai.RoleUser,
				Parts: parts,
			})
		}

		history = append(
			history,
			dto.ChatMessageResponse{
				Role:    m.Role,
				Message: m.Message,
				FileUrl: m.FileUrl,
			},
		)
	}

	aiResponseText, err := c.geminiAiService.ChatForeglycExpert(ctx, contents)
	if err != nil {
		c.log.WithError(err).Error("AI service failed")
		return nil, err
	}

	history = append(history, dto.ChatMessageResponse{
		Role:    genai.RoleModel,
		Message: aiResponseText,
	})

	return history, nil
}

func (c *ChatBotService) GlucosePrediction(ctx context.Context, userId string) (dto.PredictionResponse, error) {
	glucometerMonitoringIds, err := c.monitoringService.GetGlucometerMonitoringIds(ctx, userId)
	if err != nil {
		c.log.WithError(err).Error("failed to get glucometer monitoring ids")
		return dto.PredictionResponse{}, err
	}

	res, err := c.geminiAiService.GlucosePredictionN8N(ctx, userId, glucometerMonitoringIds)
	if err != nil {
		c.log.WithError(err).Error("failed to get glucose prediction")
		return dto.PredictionResponse{}, err
	}

	resp := dto.PredictionResponse{
		Scenario: res,
		Chats:    []dto.ChatMessageResponse{},
	}

	return resp, nil
}

func (c *ChatBotService) PredictionChatForeglycExpert(ctx context.Context, request dto.PredictionChatRequest) (dto.PredictionResponse, error) {
	var contents []*genai.Content
	var history []dto.ChatMessageResponse

	var fileInformation storage.FileInformation

	predictionInformationJson, err := json.Marshal(request.Scenario)
	if err != nil {
		c.log.WithError(err).Error("failed to marshal prediction information")
		return dto.PredictionResponse{}, err
	}

	contents = genai.Text(string(predictionInformationJson))

	for i, m := range request.Chats {
		if m.Role == genai.RoleModel {
			contents = append(contents, &genai.Content{
				Role: genai.RoleModel,
				Parts: []*genai.Part{
					{Text: m.Message},
				},
			})
		} else if m.Role == genai.RoleUser {
			isLast := i == len(request.Chats)-1

			var parts []*genai.Part
			if isLast && m.FileUrl != "" {
				var err error
				fileInformation, err = c.firebaseStorageService.GetFile(ctx, m.FileUrl)
				if err != nil {
					c.log.WithError(err).Error("failed to retrieve image")
					continue
				}
				parts = []*genai.Part{
					{Text: m.Message},
					{InlineData: &genai.Blob{Data: fileInformation.Data, MIMEType: fileInformation.Type}},
				}
			} else {
				parts = []*genai.Part{{Text: m.Message}}
			}

			contents = append(contents, &genai.Content{
				Role:  genai.RoleUser,
				Parts: parts,
			})
		}

		history = append(history, dto.ChatMessageResponse{
			Role:    m.Role,
			Message: m.Message,
			FileUrl: m.FileUrl,
		})
	}

	aiResponseText, err := c.geminiAiService.ChatForeglycExpertWithPrediction(ctx, contents)
	if err != nil {
		c.log.WithError(err).Error("AI service failed")
		return dto.PredictionResponse{}, err
	}

	history = append(history, dto.ChatMessageResponse{
		Role:    genai.RoleModel,
		Message: aiResponseText,
	})

	return dto.PredictionResponse{
		Scenario: request.Scenario,
		Chats:    history,
	}, nil
}
