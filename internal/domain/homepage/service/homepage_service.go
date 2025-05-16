package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	fooddto "github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/homepage/dto"
	monitoringdto "github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"

	"github.com/Ndraaa15/foreglyc-server/pkg/constant"
)

func (s *HomepageService) GetHomepage(ctx context.Context, userId string) (dto.HomepageResponse, error) {
	user, err := s.userService.GetUserById(ctx, userId)
	if err != nil {
		return dto.HomepageResponse{}, fmt.Errorf("get user: %w", err)
	}

	glucoseGraphs, err := s.monitoringService.GetGlucometerMonitorignGraph(ctx, monitoringdto.GetGlucometerMonitoringGraphFilter{
		UserId: userId,
		Type:   constant.GlucoseMonitoringHourly,
	})
	if err != nil {
		return dto.HomepageResponse{}, fmt.Errorf("get glucose graphs: %w", err)
	}

	today := time.Now()

	monitoringList, err := s.foodService.GetFoodMonitoring(ctx, fooddto.GetFoodMonitoringFilter{
		UserId: userId,
		Date:   today,
	})
	if err != nil {
		return dto.HomepageResponse{}, err
	}

	recommendationList, err := s.foodService.GetFoodRecommendation(ctx, fooddto.GetFoodRecommendationFilter{
		UserId: userId,
		Date:   today,
	})
	if err != nil {
		return dto.HomepageResponse{}, err
	}

	mealTime, err := s.foodService.GetDietaryPlan(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			mealTime = fooddto.DietaryPlanResponse{
				BreakfastTime:      "-",
				LunchTime:          "-",
				DinnerTime:         "-",
				MorningSnackTime:   "-",
				AfternoonSnackTime: "-",
			}
		} else {
			return dto.HomepageResponse{}, err
		}
	}

	mealTimeMap := make(map[string]string, 0)
	for _, mt := range constant.MealOrder {
		if mt == "Morning Snack" {
			mealTimeMap[mt] = mealTime.MorningSnackTime
		} else if mt == "Afternoon Snack" {
			mealTimeMap[mt] = mealTime.AfternoonSnackTime
		} else if mt == "Dinner" {
			mealTimeMap[mt] = mealTime.DinnerTime
		} else if mt == "Lunch" {
			mealTimeMap[mt] = mealTime.LunchTime
		} else if mt == "Breakfast" {
			mealTimeMap[mt] = mealTime.BreakfastTime
		}
	}

	totalCalory := int64(0)
	dietaryInformation, err := s.foodService.GetDietaryInformation(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			totalCalory = 0
		} else {
			totalCalory = dietaryInformation.TotalCalory
		}
	}

	monMap := make(map[string]fooddto.FoodMonitoringResponse, len(monitoringList))
	for _, m := range monitoringList {
		monMap[m.MealTime] = m
	}

	recMap := make(map[string]fooddto.FoodRecommendationResponse, len(recommendationList))
	for _, r := range recommendationList {
		recMap[r.MealTime] = r
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

	isCGMMonitoringPreferenceAvailable := true
	isGlucometerMonitoringPreferenceAvailable := true
	_, err = s.monitoringService.GetCGMMonitoringPreference(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			isCGMMonitoringPreferenceAvailable = false
		} else {
			return dto.HomepageResponse{}, err
		}
	}

	_, err = s.monitoringService.GetGlucometerMonitoringPreference(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			isGlucometerMonitoringPreferenceAvailable = false
		} else {
			return dto.HomepageResponse{}, err
		}
	}

	resp := dto.HomepageResponse{
		FullName:                user.FullName,
		PhotoProfile:            user.PhotoProfile,
		Level:                   user.Level,
		TotalCalory:             totalCalory,
		GlucoseMonitoringGraphs: glucoseGraphs,
		DailyFoodResponses:      daily,
		IsGlucometerMonitoringPreferenceAvailable: isGlucometerMonitoringPreferenceAvailable,
		IsCGMMonitoringPreferenceAvailable:        isCGMMonitoringPreferenceAvailable,
	}

	return resp, nil
}
