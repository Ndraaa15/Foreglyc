package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/chatbot/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/ai"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/storage"
	"github.com/sirupsen/logrus"
)

type IChatBotService interface {
	ChatForeglycExpert(ctx context.Context, requests []dto.ChatMessageRequest) ([]dto.ChatMessageResponse, error)
}

type ChatBotService struct {
	log                    *logrus.Logger
	geminiAiService        ai.IGemini
	firebaseStorageService storage.IFirebaseStorage
}

func New(log *logrus.Logger, geminiAiService ai.IGemini, firebaseStorageService storage.IFirebaseStorage) IChatBotService {
	return &ChatBotService{
		log:                    log,
		geminiAiService:        geminiAiService,
		firebaseStorageService: firebaseStorageService,
	}
}
