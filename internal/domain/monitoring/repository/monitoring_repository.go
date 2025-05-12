package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
	"github.com/jmoiron/sqlx"
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
			data.CreatedAt,
		).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	res, err := r.q.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	data.Id = lastInsertId

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
