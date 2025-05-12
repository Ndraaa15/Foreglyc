package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/repository"
	userservice "github.com/Ndraaa15/foreglyc-server/internal/domain/user/service"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/ai"
	"github.com/sirupsen/logrus"
)

type IMonitoringService interface {
	CreateCGMMonitoringPreference(ctx context.Context, request dto.CreateCGMMonitoringPreferenceRequest, userId string) (dto.CGMMonitoringPrefereceResponse, error)
	CreateGlucometerMonitoringPreference(ctx context.Context, request dto.CreateGlucometerMonitoringPreferenceRequest, userId string) (dto.GlucometerMonitoringPrefereceResponse, error)
	GetCGMMonitoringPreference(ctx context.Context, userId string) (dto.CGMMonitoringPrefereceResponse, error)
	GetGlucometerMonitoringPreference(ctx context.Context, userId string) (dto.GlucometerMonitoringPrefereceResponse, error)

	CreateGlucometerMonitoring(ctx context.Context, request dto.CreateGlucometerMonitoringRequest, userId string) (dto.GlucometerMonitoringResponse, error)
	GetGlucometerMonitorings(ctx context.Context, filter dto.GetGlucometerMonitoringFilter) ([]dto.GlucometerMonitoringResponse, error)
	GetGlucometerMonitorignGraph(ctx context.Context, filter dto.GetGlucometerMonitoringGraphFilter) ([]dto.GlucometerMonitoringGraphResponse, error)

	CreateMonitoringQuestionnaire(ctx context.Context, request dto.CreateMonitoringQuestionnaire, userId string) (dto.MonitoringQuestionnaireResponse, error)
}

type MonitoringService struct {
	log                  *logrus.Logger
	monitoringRepository repository.RepositoryItf
	geminiAiService      ai.IGemini
	userService          userservice.IUserService
}

func New(log *logrus.Logger, monitoringRepository repository.RepositoryItf, geminiAiService ai.IGemini, userService userservice.IUserService) IMonitoringService {
	return &MonitoringService{
		log:                  log,
		monitoringRepository: monitoringRepository,
		geminiAiService:      geminiAiService,
		userService:          userService,
	}
}
