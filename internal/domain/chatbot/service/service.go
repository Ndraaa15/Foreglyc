package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/chatbot/dto"
	monitoringservice "github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/service"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/ai"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/storage"
	"github.com/sirupsen/logrus"
)

type IChatBotService interface {
	ChatForeglycExpert(ctx context.Context, requests []dto.ChatMessageRequest) ([]dto.ChatMessageResponse, error)
	GlucosePrediction(ctx context.Context, userId string) (dto.PredictionResponse, error)
}

type ChatBotService struct {
	log                    *logrus.Logger
	geminiAiService        ai.IGemini
	firebaseStorageService storage.IFirebaseStorage
	monitoringService      monitoringservice.IMonitoringService
}

func New(log *logrus.Logger, geminiAiService ai.IGemini, firebaseStorageService storage.IFirebaseStorage, monitoringService monitoringservice.IMonitoringService) IChatBotService {
	return &ChatBotService{
		log:                    log,
		geminiAiService:        geminiAiService,
		firebaseStorageService: firebaseStorageService,
		monitoringService:      monitoringService,
	}
}
