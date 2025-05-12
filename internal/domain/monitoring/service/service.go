package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/repository"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/cache"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/email"
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
}

type MonitoringService struct {
	log                  *logrus.Logger
	monitoringRepository repository.RepositoryItf
	cacheRepository      cache.ICacheRepository
}

func New(log *logrus.Logger, monitoringRepository repository.RepositoryItf, cacheRepository cache.ICacheRepository, emailService email.IEmail) IMonitoringService {
	return &MonitoringService{
		log:                  log,
		monitoringRepository: monitoringRepository,
		cacheRepository:      cacheRepository,
	}
}
