package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/repository"
	userservice "github.com/Ndraaa15/foreglyc-server/internal/domain/user/service"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/ai"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/storage"
	"github.com/sirupsen/logrus"
)

type IFoodService interface {
	CreateDietaryPlan(ctx context.Context, request dto.CreateDietaryPlanRequest, userId string) (dto.DietaryPlanResponse, error)
	UpdateInsulineQuestionnaire(ctx context.Context, request dto.UpdateInsulineQuestionnaireRequest, userId string) (dto.DietaryPlanResponse, error)

	GenerateFoodInformation(ctx context.Context, request dto.CreateFoodInformationRequest) (dto.FoodInformationResponse, error)
	CreateFoodMonitoring(ctx context.Context, request dto.FoodMonitoringRequest, userId string) (dto.FoodMonitoringResponse, error)

	GenerateFoodRecomendation(ctx context.Context, userId string) ([]dto.MenuChatBotResponse, error)

	GetStatus3J(ctx context.Context, userId string) (dto.Status3JResponse, error)

	CreateDietaryInformation(ctx context.Context, request dto.CreateDietaryInformationRequest, userId string) (dto.DietaryInformationResponse, error)

	GenerateDietaryInformation(ctx context.Context, userId string) (dto.DietaryInformationChatbotResponse, error)

	GetFoodMonitoring(ctx context.Context, filter dto.GetFoodMonitoringFilter) ([]dto.FoodMonitoringResponse, error)

	GetFoodRecommendation(ctx context.Context, filter dto.GetFoodRecommendationFilter) ([]dto.FoodRecommendationResponse, error)

	GetDietaryPlan(ctx context.Context, userId string) (dto.DietaryPlanResponse, error)

	GetDietaryInformation(ctx context.Context, userId string) (dto.DietaryInformationResponse, error)
}

type FoodService struct {
	log                    *logrus.Logger
	foodRepository         repository.RepositoryItf
	geminiAiService        ai.IGemini
	firebaseStorageService storage.IFirebaseStorage
	userService            userservice.IUserService
}

func New(log *logrus.Logger, foodRepository repository.RepositoryItf, geminiAiService ai.IGemini, firebaseStorageService storage.IFirebaseStorage, userService userservice.IUserService) IFoodService {
	return &FoodService{
		log:                    log,
		foodRepository:         foodRepository,
		geminiAiService:        geminiAiService,
		firebaseStorageService: firebaseStorageService,
		userService:            userService,
	}
}
