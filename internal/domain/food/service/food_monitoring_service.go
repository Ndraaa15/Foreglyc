package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/mapper"
	"github.com/Ndraaa15/foreglyc-server/pkg/constant"
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
	response.MealTime = request.MealTime
	return response, nil
}

func (s *FoodService) CreateFoodMonitoring(ctx context.Context, request dto.FoodMonitoringRequest, userId string) (dto.FoodMonitoringResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return dto.FoodMonitoringResponse{}, err
	}

	data := entity.FoodMonitoring{
		UserID:            userId,
		FoodName:          request.FoodName,
		MealTime:          request.MealTime,
		ImageUrl:          request.ImageUrl,
		Nutritions:        request.Nutritions,
		TotalCalory:       request.TotalCalory,
		TotalCarbohydrate: request.TotalCarbohydrate,
		TotalFat:          request.TotalFat,
		TotalProtein:      request.TotalProtein,
		GlyecemicIndex:    request.GlyecemicIndex,
		CreatedAt:         pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = repository.CreateFoodMonitoring(ctx, &data)
	if err != nil {
		s.log.WithError(err).Error("failed to create food monitoring")
		return dto.FoodMonitoringResponse{}, err
	}

	return mapper.FoodMonitoringToResponse(&data), nil
}

func (r *FoodService) GetStatus3J(ctx context.Context, userId string) (dto.Status3JResponse, error) {
	repository, err := r.foodRepository.WithTx(false)
	if err != nil {
		r.log.WithError(err).Error("failed to initialize food repository")
		return dto.Status3JResponse{}, err
	}

	user, err := r.userService.GetUserById(ctx, userId)
	if err != nil {
		r.log.WithError(err).Error("failed to get user by id")
		return dto.Status3JResponse{}, err
	}

	createdAt, err := time.Parse("2006-01-02", user.CreatedAt)
	if err != nil {
		r.log.WithError(err).Error("failed to parse created at")
		return dto.Status3JResponse{}, err
	}

	totalFoodCall, err := repository.CountFoodMonitoringByFilter(ctx, dto.CountFoodMonitoringFilter{
		UserId: userId,
		Time:   createdAt,
	})
	if err != nil {
		r.log.WithError(err).Error("failed to get total food call")
		return dto.Status3JResponse{}, err
	}

	if totalFoodCall < 3 {
		return dto.Status3JResponse{IsEligible: false}, nil
	}

	return dto.Status3JResponse{IsEligible: true}, nil
}

func (s *FoodService) GetFoodMonitoring(ctx context.Context, filter dto.GetFoodMonitoringFilter) ([]dto.FoodMonitoringResponse, error) {
	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return []dto.FoodMonitoringResponse{}, err
	}

	data, err := repository.GetFoodMonitoring(ctx, filter)
	if err != nil {
		s.log.WithError(err).Error("failed to get food monitoring")
		return []dto.FoodMonitoringResponse{}, err
	}

	var resp []dto.FoodMonitoringResponse
	for _, item := range data {
		resp = append(resp, mapper.FoodMonitoringToResponse(&item))
	}

	return resp, nil
}

func (s *FoodService) GetFoodHomepage(ctx context.Context, userId string) (dto.FoodHomepageResponse, error) {
	today := time.Now()

	repository, err := s.foodRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize food repository")
		return dto.FoodHomepageResponse{}, err
	}

	monitoringList, err := repository.GetFoodMonitoring(ctx, dto.GetFoodMonitoringFilter{
		UserId: userId,
		Date:   today,
	})
	if err != nil {
		return dto.FoodHomepageResponse{}, err
	}

	recommendationList, err := repository.GetFoodRecommendation(ctx, dto.GetFoodRecommendationFilter{
		UserId: userId,
		Date:   today,
	})
	if err != nil {
		return dto.FoodHomepageResponse{}, err
	}

	mealTime, err := repository.GetDietaryPlan(ctx, userId)
	var mealTimeResponse dto.DietaryPlanResponse
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			mealTimeResponse = dto.DietaryPlanResponse{
				BreakfastTime:      "-",
				LunchTime:          "-",
				DinnerTime:         "-",
				MorningSnackTime:   "-",
				AfternoonSnackTime: "-",
			}
		} else {
			return dto.FoodHomepageResponse{}, err
		}
	}

	mealTimeResponse = mapper.DietaryPlanToResponse(&mealTime)
	mealTimeMap := make(map[string]string, 0)
	for _, mt := range constant.MealOrder {
		if mt == "Morning Snack" {
			mealTimeMap[mt] = mealTimeResponse.MorningSnackTime
		} else if mt == "Afternoon Snack" {
			mealTimeMap[mt] = mealTimeResponse.AfternoonSnackTime
		} else if mt == "Dinner" {
			mealTimeMap[mt] = mealTimeResponse.DinnerTime
		} else if mt == "Lunch" {
			mealTimeMap[mt] = mealTimeResponse.LunchTime
		} else if mt == "Breakfast" {
			mealTimeMap[mt] = mealTimeResponse.BreakfastTime
		}
	}

	monMap := make(map[string]dto.FoodMonitoringResponse, len(monitoringList))
	for _, m := range monitoringList {
		monMap[m.MealTime] = mapper.FoodMonitoringToResponse(&m)
	}

	recMap := make(map[string]dto.FoodRecommendationResponse, len(recommendationList))
	for _, r := range recommendationList {
		recMap[r.MealTime] = mapper.FoodRecommendationToResponse(&r)
	}

	daily := make([]dto.DailyFoodResponse, 0, len(constant.MealOrder))
	for _, meal := range constant.MealOrder {
		daily = append(daily, dto.DailyFoodResponse{
			MealTime:          meal,
			Time:              mealTimeMap[meal],
			FoodMonitoring:    monMap[meal],
			FoodRecomendation: recMap[meal],
		})
	}

	dietaryInformation, err := repository.GetDietaryInformation(ctx, userId)
	if err != nil {
		return dto.FoodHomepageResponse{}, err
	}

	return dto.FoodHomepageResponse{
		DailyFoodResponses: daily,
		TotalCalory:        int64(dietaryInformation.TotalCalory),
		TotalCarbohydrate:  int64(dietaryInformation.TotalCarbohydrate),
		TotalFat:           int64(dietaryInformation.TotalFat),
		TotalProtein:       int64(dietaryInformation.TotalProtein),
	}, nil

}
