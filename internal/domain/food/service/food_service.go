package service

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/mapper"
	"github.com/lib/pq"
	"google.golang.org/genai"
)

func (s FoodService) CreateDietaryPlan(ctx context.Context, request dto.CreateDietaryPlanRequest, userId string) (dto.DietaryPlanResponse, error) {
	repository, err := s.foodRepository.WithTx(false)

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
		UserID:                 userId,
		CreatedAt:              pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = repository.CreateDietaryPlan(ctx, &data)
	if err != nil {
		s.log.WithError(err).Error("failed to create dietary plan")
		return dto.DietaryPlanResponse{}, err
	}

	return mapper.DietaryPlanToResponse(&data), nil
}

func (s *FoodService) GenerateFoodInformation(ctx context.Context, request dto.CreateFoodInformationRequest) (dto.FoodInformationResponse, error) {
	fileInformation, err := s.firebaseStorageService.GetFile(ctx, request.ImageUrl)
	if err != nil {
		s.log.WithError(err).Error("failed to retrieve image")
		return dto.FoodInformationResponse{}, err
	}

	contents := genai.Blob{Data: fileInformation.Data, MIMEType: fileInformation.Type}
	response, err := s.geminiAiService.GenerateFoodInformation(ctx, []*genai.Content{{Parts: []*genai.Part{{InlineData: &contents}}}})
	if err != nil {
		s.log.WithError(err).Error("failed to generate food information")
		return dto.FoodInformationResponse{}, err
	}

	response.ImageUrl = request.ImageUrl
	response.TimeType = request.TimeType
	return response, nil
}

func (s *FoodService) CreateFoodRecall(ctx context.Context, request dto.FoodRecallRequest, userId string) (dto.FoodRecallResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create food recall")
		return dto.FoodRecallResponse{}, err
	}

	data := entity.FoodRecall{
		UserID:        userId,
		FoodName:      request.FoodName,
		TimeType:      request.TimeType,
		ImageUrl:      request.ImageUrl,
		Nutritions:    request.Nutritions,
		TotalCalories: request.TotalCalories,
		CreatedAt:     pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = repository.CreateFoodRecall(ctx, &data)
	if err != nil {
		s.log.WithError(err).Error("failed to create food recall")
		return dto.FoodRecallResponse{}, err
	}

	return mapper.FoodRecallToResponse(&data), nil
}

func (r *FoodService) GetStatus3J(ctx context.Context, userId string) (dto.Status3JResponse, error) {
	repository, err := r.foodRepository.WithTx(false)
	if err != nil {
		r.log.WithError(err).Error("failed to create food recall")
		return dto.Status3JResponse{}, err
	}

	totalFoodCall, err := repository.GetCountFoodTotal(ctx, userId)
	if err != nil {
		r.log.WithError(err).Error("failed to get total food call")
		return dto.Status3JResponse{}, err
	}

	if totalFoodCall < 3 {
		return dto.Status3JResponse{IsEligible: false}, nil
	}

	return dto.Status3JResponse{IsEligible: true}, nil
}
