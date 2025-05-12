package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (r *MonitoringRepository) CreateGlucometerMonitoring(ctx context.Context, data *entity.GlucometerMonitoring) error {
	query, args, err := squirrel.
		Insert(GlucometerMonitoringTable).
		Columns(
			"user_id",
			"blood_glucose",
			"unit",
			"status",
			"created_at",
		).
		Values(
			data.UserId,
			data.BloodGlucose,
			data.Unit,
			data.Status,
			data.CreatedAt,
		).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	err = r.q.QueryRowxContext(ctx, query, args...).Scan(&data.Id)
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	return nil
}

func (r *MonitoringRepository) GetGlucometerMonitorings(ctx context.Context, filter dto.GetGlucometerMonitoringFilter) ([]entity.GlucometerMonitoring, error) {
	query, args, err := squirrel.
		Select(
			"id",
			"user_id",
			"blood_glucose",
			"unit",
			"status",
			"created_at",
			"updated_at",
		).
		From(GlucometerMonitoringTable).
		Where("user_id = ?", filter.UserId).
		OrderBy("created_at desc").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}
	query = r.q.Rebind(query)

	var data []entity.GlucometerMonitoring
	err = sqlx.SelectContext(ctx, r.q, &data, query, args...)
	if err != nil {
		return nil, err
	}

	return data, nil
}
