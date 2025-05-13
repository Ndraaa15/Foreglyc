package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
)

func (r *FoodRepository) CreateDietaryPlan(ctx context.Context, dietaryPlan *entity.DietaryPlan) error {
	query, args, err := squirrel.Insert(DietaryPlanTable).
		Columns(
			"user_id",
			"live_with",
			"breakfast_time",
			"lunch_time",
			"dinner_time",
			"morning_snack_time",
			"afternoon_snack_time",
			"is_use_insuline",
			"insulise_questionnaires",
			"total_daily_insuline",
			"meal_plan_type",
			"created_at",
		).
		Values(
			dietaryPlan.UserID,
			dietaryPlan.LiveWith,
			dietaryPlan.BreakfastTime,
			dietaryPlan.LunchTime,
			dietaryPlan.DinnerTime,
			dietaryPlan.MorningSnackTime,
			dietaryPlan.AfternoonSnackTime,
			dietaryPlan.IsUseInsuline,
			dietaryPlan.InsuliseQuestionnaires,
			dietaryPlan.TotalDailyInsuline,
			dietaryPlan.MealPlanType,
			dietaryPlan.CreatedAt,
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.q.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *FoodRepository) CreateFoodRecall(ctx context.Context, foodRecall *entity.FoodRecall) error {
	query, args, err := squirrel.Insert(FoodRecallTable).
		Columns(
			"user_id",
			"food_name",
			"time_type",
			"image_url",
			"nutritions",
			"total_calories",
			"created_at",
		).
		Values(
			foodRecall.UserID,
			foodRecall.FoodName,
			foodRecall.TimeType,
			foodRecall.ImageUrl,
			foodRecall.Nutritions,
			foodRecall.TotalCalories,
			foodRecall.CreatedAt,
		).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	err = r.q.QueryRowxContext(ctx, query, args...).Scan(&foodRecall.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *FoodRepository) GetCountFoodTotal(ctx context.Context, userId string) (int64, error) {
	query, args, err := squirrel.Select("count(*)").
		From(FoodRecallTable).
		Where("user_id = ?", userId).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return 0, err
	}

	var count int64
	err = r.q.QueryRowxContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
