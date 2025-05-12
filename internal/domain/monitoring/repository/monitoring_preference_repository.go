package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
	"github.com/jmoiron/sqlx"
)

func (r *MonitoringRepository) CreateGlucometerMonitoringPreference(ctx context.Context, data *entity.GlucometerMonitoringPreference) error {
	query, args, err := squirrel.
		Insert(GlucometerMonitoringPrefereceTable).
		Columns(
			"user_id",
			"start_wake_up_time",
			"end_wake_up_time",
			"physical_activity_days",
			"start_sleep_time",
			"end_sleep_time",
			"hyphoglycemia_accute_threshold",
			"hyphoglycemia_chronic_threshold",
			"hyperglycemia_accute_threshold",
			"hyperglycemia_chronic_threshold",
			"send_notification",
			"created_at",
		).
		Values(
			data.UserId,
			data.StartWakeUpTime,
			data.EndWakeUpTime,
			data.PhysicalActivityDays,
			data.StartSleepTime,
			data.EndSleepTime,
			data.HypoglycemiaAccuteThreshold,
			data.HypoglycemiaChronicThreshold,
			data.HyperglycemiaAccuteThreshold,
			data.HyperglycemiaChronicThreshold,
			data.SendNotification,
			data.CreatedAt,
		).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return err
	}

	_, err = r.q.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *MonitoringRepository) CreateCGMMonitoringPreference(ctx context.Context, data *entity.CGMMonitoringPreference) error {
	query, args, err := squirrel.
		Insert(CGMMonitoringPreferenceTable).
		Columns(
			"user_id",
			"hyphoglycemia_accute_threshold",
			"hyphoglycemia_chronic_threshold",
			"hyperglycemia_accute_threshold",
			"hyperglycemia_chronic_threshold",
			"send_notification",
			"created_at",
		).
		Values(
			data.UserId,
			data.HypoglycemiaAccuteThreshold,
			data.HypoglycemiaChronicThreshold,
			data.HyperglycemiaAccuteThreshold,
			data.HyperglycemiaChronicThreshold,
			data.SendNotification,
			data.CreatedAt,
		).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return err
	}

	_, err = r.q.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *MonitoringRepository) GetGlucometerMonitoringPreference(ctx context.Context, userId string) (entity.GlucometerMonitoringPreference, error) {
	query, args, err := squirrel.
		Select(
			"id",
			"user_id",
			"start_wake_up_time",
			"end_wake_up_time",
			"physical_activity_days",
			"start_sleep_time",
			"end_sleep_time",
			"hyphoglycemia_accute_threshold",
			"hyphoglycemia_chronic_threshold",
			"hyperglycemia_accute_threshold",
			"hyperglycemia_chronic_threshold",
			"send_notification",
			"created_at",
			"updated_at",
		).
		From(GlucometerMonitoringPrefereceTable).
		Where("user_id = ?", userId).
		OrderBy("created_at desc").
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return entity.GlucometerMonitoringPreference{}, err
	}
	query = r.q.Rebind(query)

	var data entity.GlucometerMonitoringPreference
	err = sqlx.GetContext(ctx, r.q, &data, query, args...)
	if err != nil {
		return entity.GlucometerMonitoringPreference{}, err
	}

	return data, nil
}

func (r *MonitoringRepository) GetCGMMonitoringPreference(ctx context.Context, userId string) (entity.CGMMonitoringPreference, error) {
	query, args, err := squirrel.
		Select(
			"id",
			"user_id",
			"hyphoglycemia_accute_threshold",
			"hyphoglycemia_chronic_threshold",
			"hyperglycemia_accute_threshold",
			"hyperglycemia_chronic_threshold",
			"send_notification",
			"created_at",
			"updated_at",
		).
		From(CGMMonitoringPreferenceTable).
		Where("user_id = ?", userId).
		OrderBy("created_at desc").
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return entity.CGMMonitoringPreference{}, err
	}
	query = r.q.Rebind(query)

	var data entity.CGMMonitoringPreference
	err = sqlx.GetContext(ctx, r.q, &data, query, args...)
	if err != nil {
		return entity.CGMMonitoringPreference{}, err
	}

	return data, nil
}
