package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/report/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/report/entity"
	"github.com/jmoiron/sqlx"
)

func (r *ReportRepository) GetReportInformations(ctx context.Context, filter dto.GetReportInformationFilter) ([]entity.ReportInformation, error) {
	queryBuilder := squirrel.Select("*").From(ReportInformationTable)

	if filter.UserId != "" {
		queryBuilder = queryBuilder.Where("user_id = ?", filter.UserId)
	}

	if !filter.StartDate.IsZero() {
		queryBuilder = queryBuilder.Where("created_at >= ?", filter.StartDate)
	}

	if !filter.EndDate.IsZero() {
		queryBuilder = queryBuilder.Where("created_at <= ?", filter.EndDate)
	}

	query, args, err := queryBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	var reportInformations []entity.ReportInformation
	err = sqlx.SelectContext(ctx, r.q, &reportInformations, query, args...)
	if err != nil {
		return nil, err
	}

	return reportInformations, nil
}
