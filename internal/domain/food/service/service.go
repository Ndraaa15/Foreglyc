package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/repository"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/ai"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/storage"
	"github.com/sirupsen/logrus"
)

type IFoodService interface {
	CreateDietaryPlan(ctx context.Context, request dto.CreateDietaryPlanRequest, userId string) (dto.DietaryPlanResponse, error)
	UpdateInsulineQuestionnaire(ctx context.Context, request dto.UpdateInsulineQuestionnaireRequest, userId string) (dto.DietaryPlanResponse, error)

	GenerateFoodInformation(ctx context.Context, request dto.CreateFoodInformationRequest) (dto.FoodInformationResponse, error)
	CreateFoodMonitoring(ctx context.Context, request dto.FoodMonitoringRequest, userId string) (dto.FoodMonitoringResponse, error)

	GetStatus3J(ctx context.Context, userId string) (dto.Status3JResponse, error)
}

type FoodService struct {
	log                    *logrus.Logger
	foodRepository         repository.RepositoryItf
	geminiAiService        ai.IGemini
	firebaseStorageService storage.IFirebaseStorage
}

func New(log *logrus.Logger, foodRepository repository.RepositoryItf, geminiAiService ai.IGemini, firebaseStorageService storage.IFirebaseStorage) IFoodService {
	return &FoodService{
		log:                    log,
		foodRepository:         foodRepository,
		geminiAiService:        geminiAiService,
		firebaseStorageService: firebaseStorageService,
	}
}
