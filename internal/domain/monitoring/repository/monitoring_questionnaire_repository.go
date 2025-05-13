package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
)

func (r *MonitoringRepository) CreateMonitoringQuestionnaire(ctx context.Context, data *entity.MonitoringQuestionnaire) error {
	query, args, err := squirrel.Insert(MonitoringQuestionnaireTable).
		Columns(
			"glucometer_monitoring_id",
			"questionnaires",
			"management_type",
			"created_at",
		).Values(
		data.GlucometerMonitoringID,
		data.Questionnaires,
		data.ManagementType,
		data.CreatedAt.Time,
	).Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	err = r.q.QueryRowxContext(ctx, query, args...).Scan(&data.Id)
	if err != nil {
		return err
	}

	return nil
}
