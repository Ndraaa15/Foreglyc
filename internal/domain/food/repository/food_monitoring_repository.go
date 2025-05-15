package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
	"github.com/jmoiron/sqlx"
)

func (r *FoodRepository) CreateFoodMonitoring(ctx context.Context, foodRecall *entity.FoodMonitoring) error {
	query, args, err := squirrel.Insert(FoodMonitoringTable).
		Columns(
			"user_id",
			"food_name",
			"meal_time",
			"image_url",
			"nutritions",
			"total_calory",
			"total_carbohydrate",
			"total_fat",
			"total_protein",
			"created_at",
		).
		Values(
			foodRecall.UserID,
			foodRecall.FoodName,
			foodRecall.MealTime,
			foodRecall.ImageUrl,
			foodRecall.Nutritions,
			foodRecall.TotalCalory,
			foodRecall.TotalCarbohydrate,
			foodRecall.TotalFat,
			foodRecall.TotalProtein,
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

func (r *FoodRepository) GetFoodMonitoring(ctx context.Context, filter dto.GetFoodMonitoringFilter) ([]entity.FoodMonitoring, error) {
	queryBuilder := squirrel.Select("*").
		From(FoodMonitoringTable)

	if filter.UserId != "" {
		queryBuilder = queryBuilder.Where("user_id = ?", filter.UserId)
	}

	if !filter.Date.IsZero() {
		queryBuilder = queryBuilder.Where("DATE(created_at) = ?", filter.Date)
	}

	query, args, err := queryBuilder.PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var foodMonitoring []entity.FoodMonitoring
	err = sqlx.SelectContext(ctx, r.q, &foodMonitoring, query, args...)
	if err != nil {
		return nil, err
	}

	return foodMonitoring, nil
}

func (r *FoodRepository) CountFoodMonitoringByFilter(ctx context.Context, filter dto.CountFoodMonitoringFilter) (int64, error) {
	queryBuilder := squirrel.Select("count(*)").
		From(FoodMonitoringTable)

	if filter.UserId != "" {
		queryBuilder = queryBuilder.Where("user_id = ?", filter.UserId)
	}

	if !filter.Time.IsZero() {
		queryBuilder = queryBuilder.Where("created_at <= ?", filter.Time)
	}

	query, args, err := queryBuilder.
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
