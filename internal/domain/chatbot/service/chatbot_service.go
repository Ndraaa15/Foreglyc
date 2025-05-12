package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/chatbot/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/storage"
	"google.golang.org/genai"
)

func (c *ChatBotService) ChatForeglycExpert(ctx context.Context, requests []dto.ChatMessageRequest) ([]dto.ChatMessageResponse, error) {
	var contents []*genai.Content
	var history []dto.ChatMessageResponse

	var fileInformation storage.FileInformation

	// Process all previous messages
	for i, m := range requests {
		if m.Role == genai.RoleModel {
			contents = append(contents, &genai.Content{
				Role: genai.RoleModel,
				Parts: []*genai.Part{
					{Text: m.Message},
				},
			})
		} else if m.Role == genai.RoleUser {
			// Check if this is the last message
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

		// Append to response history
		history = append(history, dto.ChatMessageResponse{
			Role:    m.Role,
			Message: m.Message,
			FileUrl: m.FileUrl,
		})
	}

	// Get AI response
	aiResponseText, err := c.geminiAiService.ChatForeglycExpert(ctx, contents)
	if err != nil {
		c.log.WithError(err).Error("AI service failed")
		return nil, err
	}

	// Append AI model reply to history
	history = append(history, dto.ChatMessageResponse{
		Role:    "model",
		Message: aiResponseText,
	})

	return history, nil
}

func (c *ChatBotService) GlucosePrediction() {

}
