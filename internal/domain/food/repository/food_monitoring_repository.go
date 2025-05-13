package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
)

func (r *FoodRepository) CreateFoodMonitoring(ctx context.Context, foodRecall *entity.FoodMonitoring) error {
	query, args, err := squirrel.Insert(FoodMonitoringTable).
		Columns(
			"user_id",
			"food_name",
			"time_type",
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
			foodRecall.TimeType,
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

func (r *FoodRepository) CountFoodMonitoring(ctx context.Context, userId string) (int64, error) {
	query, args, err := squirrel.Select("count(*)").
		From(FoodMonitoringTable).
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
