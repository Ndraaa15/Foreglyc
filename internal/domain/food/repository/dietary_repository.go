package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
	"github.com/jmoiron/sqlx"
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
			dietaryPlan.UserId,
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

func (r *FoodRepository) GetDietaryPlan(ctx context.Context, userId string) (entity.DietaryPlan, error) {
	query, args, err := squirrel.
		Select(
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
			"updated_at",
		).
		From(DietaryPlanTable).
		Where("user_id = ?", userId).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return entity.DietaryPlan{}, err
	}

	var dietaryPlan entity.DietaryPlan
	err = sqlx.GetContext(ctx, r.q, &dietaryPlan, query, args...)
	if err != nil {
		return entity.DietaryPlan{}, err
	}

	return dietaryPlan, nil
}

func (r *FoodRepository) UpdateDietaryPlan(ctx context.Context, dietaryPlan *entity.DietaryPlan) error {
	query, args, err := squirrel.Update(DietaryPlanTable).
		Set("live_with", dietaryPlan.LiveWith).
		Set("breakfast_time", dietaryPlan.BreakfastTime).
		Set("lunch_time", dietaryPlan.LunchTime).
		Set("dinner_time", dietaryPlan.DinnerTime).
		Set("morning_snack_time", dietaryPlan.MorningSnackTime).
		Set("afternoon_snack_time", dietaryPlan.AfternoonSnackTime).
		Set("is_use_insuline", dietaryPlan.IsUseInsuline).
		Set("insulise_questionnaires", dietaryPlan.InsuliseQuestionnaires).
		Set("total_daily_insuline", dietaryPlan.TotalDailyInsuline).
		Set("meal_plan_type", dietaryPlan.MealPlanType).
		Set("updated_at", dietaryPlan.UpdatedAt).
		Where("user_id = ?", dietaryPlan.UserId).
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

func (r *FoodRepository) CreateDietaryInformation(ctx context.Context, dietaryInformation *entity.DietaryInformation) error {
	query, args, err := squirrel.Insert(DietaryInformationTable).
		Columns(
			"user_id",
			"total_snack_calory",
			"total_calory",
			"total_breakfast_calory",
			"total_lunch_calory",
			"total_dinner_calory",
			"created_at",
		).
		Values(
			dietaryInformation.UserId,
			dietaryInformation.TotalSnackCalory,
			dietaryInformation.TotalCalory,
			dietaryInformation.TotalBreakfastCalory,
			dietaryInformation.TotalLunchCalory,
			dietaryInformation.TotalDinnerCalory,
			dietaryInformation.CreatedAt,
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

func (r *FoodRepository) GetDietaryInformation(ctx context.Context, userId string) (entity.DietaryInformation, error) {
	query, args, err := squirrel.
		Select(
			"user_id",
			"total_snack_calory",
			"total_calory",
			"total_breakfast_calory",
			"total_lunch_calory",
			"total_dinner_calory",
			"created_at",
			"updated_at",
		).
		From(DietaryInformationTable).
		Where("user_id = ?", userId).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return entity.DietaryInformation{}, err
	}

	var dietaryInformation entity.DietaryInformation
	err = sqlx.GetContext(ctx, r.q, &dietaryInformation, query, args...)
	if err != nil {
		return entity.DietaryInformation{}, err
	}

	return dietaryInformation, nil
}
