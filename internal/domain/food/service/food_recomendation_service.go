package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/mapper"
)

func (s *FoodService) GenerateFoodRecomendation(ctx context.Context, userId string) ([]dto.MenuChatBotResponse, error) {
	return s.geminiAiService.FoodRecomendationsN8N(ctx, userId)
}

func (s *FoodService) GetFoodRecommendation(ctx context.Context, filter dto.GetFoodRecommendationFilter) ([]dto.FoodRecommendationResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return []dto.FoodRecommendationResponse{}, err
	}

	foodRecomendation, err := repository.GetFoodRecomendation(ctx, filter)
	if err != nil {
		s.log.WithError(err).Error("failed to get food recomendation")
		return []dto.FoodRecommendationResponse{}, err
	}

	foodRecomendationResponse := make([]dto.FoodRecommendationResponse, len(foodRecomendation))

	for i, item := range foodRecomendation {
		foodRecomendationResponse[i] = mapper.FoodRecommendationToResponse(&item)
	}

	return foodRecomendationResponse, nil
}

func (s *FoodService) CreateFoodRecomendation(ctx context.Context, request []dto.CreateMenuRequest, userId string) (dto.FoodRecommendationResponse, error) {
	return dto.FoodRecommendationResponse{}, nil
}
