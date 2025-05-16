package service

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/mapper"
	"github.com/lib/pq"
)

func (s *FoodService) GenerateFoodRecommendation(ctx context.Context, userId string) ([]dto.MenuResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return nil, err
	}

	menuChatBotResponse, err := s.geminiAiService.FoodRecomendationsN8N(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to generate food recommendation")
		return nil, err
	}

	var flatRecs []*entity.FoodRecommendation
	dateOrder := make([]time.Time, 0, len(menuChatBotResponse))
	byDate := make(map[time.Time][]*entity.FoodRecommendation)

	for _, day := range menuChatBotResponse {
		d, err := time.Parse("Monday, 02 Jan 2006", day.Date)
		if err != nil {
			s.log.WithError(err).Error("failed to parse date")
			return nil, err
		}
		dateOrder = append(dateOrder, d)

		for _, f := range day.FoodRecomendations {
			rec := &entity.FoodRecommendation{
				UserId:                 userId,
				FoodName:               f.FoodName,
				MealTime:               f.MealTime,
				Ingredients:            f.Ingredients,
				CaloriesPerIngredients: f.CaloriesPerIngredients,
				TotalCalory:            f.TotalCalories,
				GlycemicIndex:          f.GlycemicIndex,
				ImageUrl:               f.ImageUrl,
				Date:                   d,
				CreatedAt:              pq.NullTime{Time: time.Now(), Valid: true},
			}
			flatRecs = append(flatRecs, rec)
			byDate[d] = append(byDate[d], rec)
		}
	}

	if err := repository.CreateFoodRecommendations(ctx, flatRecs); err != nil {
		s.log.WithError(err).Error("failed to create food recommendations")
		return nil, err
	}

	out := make([]dto.MenuResponse, 0, len(dateOrder))
	for _, d := range dateOrder {
		dayRecs := byDate[d]
		dtoRecs := make([]dto.FoodRecommendationResponse, len(dayRecs))
		for i, er := range dayRecs {
			dtoRecs[i] = mapper.FoodRecommendationToResponse(er)
		}
		out = append(out, dto.MenuResponse{
			Date:               d.Format("Monday, 02 Jan 2006"),
			FoodRecomendations: dtoRecs,
		})
	}

	return out, nil
}

func (s *FoodService) GetFoodRecommendation(ctx context.Context, filter dto.GetFoodRecommendationFilter) ([]dto.FoodRecommendationResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return []dto.FoodRecommendationResponse{}, err
	}

	foodRecomendation, err := repository.GetFoodRecommendation(ctx, filter)
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
