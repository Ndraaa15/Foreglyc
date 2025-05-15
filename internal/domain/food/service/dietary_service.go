package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/mapper"
	"github.com/lib/pq"
	"google.golang.org/genai"
)

func (s FoodService) CreateDietaryPlan(ctx context.Context, request dto.CreateDietaryPlanRequest, userId string) (dto.DietaryPlanResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return dto.DietaryPlanResponse{}, err
	}

	breakfastTime, err := time.Parse("15:04", request.BreakfastTime)
	if err != nil {
		s.log.WithError(err).Error("failed to parse breakfast time")
		return dto.DietaryPlanResponse{}, err
	}

	lunchTime, err := time.Parse("15:04", request.LunchTime)
	if err != nil {
		s.log.WithError(err).Error("failed to parse lunch time")
		return dto.DietaryPlanResponse{}, err
	}

	dinnerTime, err := time.Parse("15:04", request.DinnerTime)
	if err != nil {
		s.log.WithError(err).Error("failed to parse dinner time")
		return dto.DietaryPlanResponse{}, err
	}

	morningSnackTime, err := time.Parse("15:04", request.MorningSnackTime)
	if err != nil {
		s.log.WithError(err).Error("failed to parse morning snack time")
		return dto.DietaryPlanResponse{}, err
	}

	afternoonSnackTime, err := time.Parse("15:04", request.AfternoonSnackTime)
	if err != nil {
		s.log.WithError(err).Error("failed to parse afternoon snack time")
		return dto.DietaryPlanResponse{}, err
	}

	data := entity.DietaryPlan{
		LiveWith:               request.LiveWith,
		BreakfastTime:          breakfastTime,
		LunchTime:              lunchTime,
		DinnerTime:             dinnerTime,
		MorningSnackTime:       morningSnackTime,
		AfternoonSnackTime:     afternoonSnackTime,
		IsUseInsuline:          *request.IsUseInsuline,
		InsuliseQuestionnaires: request.InsuliseQuestionnaires,
		TotalDailyInsuline:     request.TotalDailyInsuline,
		MealPlanType:           request.MealPlanType,
		UserId:                 userId,
		CreatedAt:              pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = repository.CreateDietaryPlan(ctx, &data)
	if err != nil {
		s.log.WithError(err).Error("failed to create dietary plan")
		return dto.DietaryPlanResponse{}, err
	}

	return mapper.DietaryPlanToResponse(&data), nil
}

func (s *FoodService) UpdateInsulineQuestionnaire(ctx context.Context, request dto.UpdateInsulineQuestionnaireRequest, userId string) (dto.DietaryPlanResponse, error) {

	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return dto.DietaryPlanResponse{}, err
	}

	dietaryPlan, err := repository.GetDietaryPlan(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get dietary plan")
		return dto.DietaryPlanResponse{}, err
	}

	dietaryPlan.InsuliseQuestionnaires = request.InsuliseQuestionnaires
	dietaryPlan.TotalDailyInsuline = request.TotalDailyInsuline

	err = repository.UpdateDietaryPlan(ctx, &dietaryPlan)
	if err != nil {
		s.log.WithError(err).Error("failed to update dietary plan")
		return dto.DietaryPlanResponse{}, err
	}

	return mapper.DietaryPlanToResponse(&dietaryPlan), nil
}

func (s *FoodService) GetDietaryPlan(ctx context.Context, userId string) (dto.DietaryPlanResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return dto.DietaryPlanResponse{}, err
	}

	dietaryPlan, err := repository.GetDietaryPlan(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get dietary plan")
		return dto.DietaryPlanResponse{}, err
	}

	return mapper.DietaryPlanToResponse(&dietaryPlan), nil
}

func (s *FoodService) GenerateDietaryInformation(ctx context.Context, userId string) (dto.DietaryInformationChatbotResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return dto.DietaryInformationChatbotResponse{}, err
	}

	foodMonitorings, err := repository.GetFoodMonitoring(ctx, dto.GetFoodMonitoringFilter{UserId: userId})
	if err != nil {
		s.log.WithError(err).Error("failed to get food monitorings")
		return dto.DietaryInformationChatbotResponse{}, err
	}

	foodMonitoringJson, err := json.Marshal(foodMonitorings)
	if err != nil {
		s.log.WithError(err).Error("failed to marshal food monitorings")
		return dto.DietaryInformationChatbotResponse{}, err
	}

	contents := genai.Text(string(foodMonitoringJson))

	dietaryInformation, err := s.geminiAiService.GenerateDietaryInformation(ctx, contents)
	if err != nil {
		s.log.WithError(err).Error("failed to generate dietary information")
		return dto.DietaryInformationChatbotResponse{}, err
	}

	return dietaryInformation, nil
}

func (s *FoodService) CreateDietaryInformation(ctx context.Context, request dto.CreateDietaryInformationRequest, userId string) (dto.DietaryInformationResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return dto.DietaryInformationResponse{}, err
	}

	dietaryInformation := entity.DietaryInformation{
		UserId:               userId,
		TotalSnackCalory:     request.TotalSnackCalory,
		TotalCalory:          request.TotalCalory,
		TotalBreakfastCalory: request.TotalBreakfastCalory,
		TotalLunchCalory:     request.TotalLunchCalory,
		TotalDinnerCalory:    request.TotalDinnerCalory,
		CreatedAt:            pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = repository.CreateDietaryInformation(ctx, &dietaryInformation)
	if err != nil {
		s.log.WithError(err).Error("failed to create dietary information")
		return dto.DietaryInformationResponse{}, err
	}

	return mapper.DietaryInformationToResponse(&dietaryInformation), nil
}

func (s *FoodService) GetDietaryInformation(ctx context.Context, userId string) (dto.DietaryInformationResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return dto.DietaryInformationResponse{}, err
	}

	dietaryInformation, err := repository.GetDietaryInformation(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get dietary information")
		return dto.DietaryInformationResponse{}, err
	}

	return mapper.DietaryInformationToResponse(&dietaryInformation), nil
}
