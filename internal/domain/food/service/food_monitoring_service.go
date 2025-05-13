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

func (s *FoodService) CreateFoodMonitoring(ctx context.Context, request dto.FoodMonitoringRequest, userId string) (dto.FoodMonitoringResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create food recall")
		return dto.FoodMonitoringResponse{}, err
	}

	data := entity.FoodMonitoring{
		UserID:            userId,
		FoodName:          request.FoodName,
		TimeType:          request.TimeType,
		ImageUrl:          request.ImageUrl,
		Nutritions:        request.Nutritions,
		TotalCalory:       request.TotalCalory,
		TotalCarbohydrate: request.TotalCarbohydrate,
		TotalFat:          request.TotalFat,
		TotalProtein:      request.TotalProtein,
		CreatedAt:         pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = repository.CreateFoodMonitoring(ctx, &data)
	if err != nil {
		s.log.WithError(err).Error("failed to create food recall")
		return dto.FoodMonitoringResponse{}, err
	}

	return mapper.FoodRecallToResponse(&data), nil
}

func (r *FoodService) GetStatus3J(ctx context.Context, userId string) (dto.Status3JResponse, error) {
	repository, err := r.foodRepository.WithTx(false)
	if err != nil {
		r.log.WithError(err).Error("failed to create food recall")
		return dto.Status3JResponse{}, err
	}

	totalFoodCall, err := repository.CountFoodMonitoring(ctx, userId)
	if err != nil {
		r.log.WithError(err).Error("failed to get total food call")
		return dto.Status3JResponse{}, err
	}

	if totalFoodCall < 3 {
		return dto.Status3JResponse{IsEligible: false}, nil
	}

	return dto.Status3JResponse{IsEligible: true}, nil
}
