package service

import (
	"context"

	foodservice "github.com/Ndraaa15/foreglyc-server/internal/domain/food/service"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/homepage/dto"
	monitoringservice "github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/service"
	userservice "github.com/Ndraaa15/foreglyc-server/internal/domain/user/service"
	"github.com/sirupsen/logrus"
)

type IHomepageService interface {
	GetHomepage(ctx context.Context, userId string) (dto.HomepageResponse, error)
}

type HomepageService struct {
	log               *logrus.Logger
	monitoringService monitoringservice.IMonitoringService
	foodService       foodservice.IFoodService
	userService       userservice.IUserService
}

func New(log *logrus.Logger, monitoringService monitoringservice.IMonitoringService, foodService foodservice.IFoodService, userService userservice.IUserService) IHomepageService {
	return &HomepageService{
		log:               log,
		monitoringService: monitoringService,
		foodService:       foodService,
		userService:       userService,
	}
}
