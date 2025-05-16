package repository

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/report/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/report/entity"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/jmoiron/sqlx"
)

var (
	ErrFailedToCommit   = errx.InternalServerError("failed to commit transaction")
	ErrFailedToRollback = errx.InternalServerError("failed to rollback transaction")
)

const (
	ReportInformationTable = "report_informations"
)

type Repository struct {
	DB *sqlx.DB
}

type RepositoryItf interface {
	WithTx(tx bool) (ReportRepositoryItf, error)
}

func New(db *sqlx.DB) RepositoryItf {
	return &Repository{db}
}

type ReportRepository struct {
	q sqlx.ExtContext
}

type ReportRepositoryItf interface {
	Commit() error
	Rollback() error

	GetReportInformations(ctx context.Context, filter dto.GetReportInformationFilter) ([]entity.ReportInformation, error)
}

func (r *Repository) WithTx(tx bool) (ReportRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &ReportRepository{db}, nil
}

func (r *ReportRepository) Commit() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Commit()
	}

	return ErrFailedToCommit
}

func (r *ReportRepository) Rollback() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Rollback()
	}

	return ErrFailedToRollback
}
