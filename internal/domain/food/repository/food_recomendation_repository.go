package repository

import (
	"context"

	"fmt"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
	"github.com/jmoiron/sqlx"

	"github.com/Masterminds/squirrel"
)

func (r *FoodRepository) CreateFoodRecommendations(ctx context.Context, recs []*entity.FoodRecommendation) error {
	if len(recs) == 0 {
		return nil
	}

	builder := squirrel.
		Insert(FoodRecomendationTable).
		Columns(
			"user_id",
			"food_name",
			"meal_time",
			"ingredients",
			"calories_per_ingredients",
			"image_url",
			"total_calory",
			"glycemic_index",
			"date",
			"created_at",
		).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id")

	for _, f := range recs {
		builder = builder.Values(
			f.UserId,
			f.FoodName,
			f.MealTime,
			f.Ingredients,
			f.CaloriesPerIngredients,
			f.ImageUrl,
			f.TotalCalory,
			f.GlycemicIndex,
			f.Date,
			f.CreatedAt,
		)
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build batch insert SQL: %w", err)
	}

	rows, err := r.q.QueryContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("exec batch insert with returning id: %w", err)
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return fmt.Errorf("scan returned id: %w", err)
		}
		recs[i].Id = id
		i++
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("iterate returned ids: %w", err)
	}
	if i != len(recs) {
		return fmt.Errorf("expected %d returned ids, got %d", len(recs), i)
	}

	return nil
}

func (r *FoodRepository) GetFoodRecommendation(ctx context.Context, filter dto.GetFoodRecommendationFilter) ([]entity.FoodRecommendation, error) {
	queryBuilder := squirrel.Select("*").From(FoodRecomendationTable)

	if filter.UserId != "" {
		queryBuilder = queryBuilder.Where("user_id = ?", filter.UserId)
	}

	if !filter.Date.IsZero() {
		queryBuilder = queryBuilder.Where("DATE(date) = ?", filter.Date)
	}

	query, args, err := queryBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var foodRecommendations []entity.FoodRecommendation
	err = sqlx.SelectContext(ctx, r.q, &foodRecommendations, query, args...)
	if err != nil {
		return nil, err
	}

	return foodRecommendations, nil
}
