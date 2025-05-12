package repository

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/jmoiron/sqlx"
)

var (
	ErrFailedToCommit   = errx.InternalServerError("failed to commit transaction")
	ErrFailedToRollback = errx.InternalServerError("failed to rollback transaction")
)

const (
	CGMMonitoringPreferenceTable       = "cgm_monitoring_preferences"
	GlucometerMonitoringPrefereceTable = "glucometer_monitoring_preferences"
	GlucometerMonitoringTable          = "glucometer_monitorings"
	MonitoringQuestionnaireTable       = "monitoring_questionnaires"
)

type Repository struct {
	DB *sqlx.DB
}

type RepositoryItf interface {
	WithTx(tx bool) (MonitoringRepositoryItf, error)
}

func New(db *sqlx.DB) RepositoryItf {
	return &Repository{db}
}

type MonitoringRepository struct {
	q sqlx.ExtContext
}

type MonitoringRepositoryItf interface {
	Commit() error
	Rollback() error

	CreateGlucometerMonitoringPreference(ctx context.Context, data *entity.GlucometerMonitoringPreference) error
	CreateCGMMonitoringPreference(ctx context.Context, data *entity.CGMMonitoringPreference) error
	GetGlucometerMonitoringPreference(ctx context.Context, userId string) (entity.GlucometerMonitoringPreference, error)
	GetCGMMonitoringPreference(ctx context.Context, userId string) (entity.CGMMonitoringPreference, error)

	CreateGlucometerMonitoring(ctx context.Context, data *entity.GlucometerMonitoring) error
	GetGlucometerMonitorings(ctx context.Context, filter dto.GetGlucometerMonitoringFilter) ([]entity.GlucometerMonitoring, error)

	CreateMonitoringQuestionnaire(ctx context.Context, data *entity.MonitoringQuestionnaire) error
}

func (r *Repository) WithTx(tx bool) (MonitoringRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &MonitoringRepository{db}, nil
}

func (r *MonitoringRepository) Commit() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Commit()
	}

	return ErrFailedToCommit
}

func (r *MonitoringRepository) Rollback() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Rollback()
	}

	return ErrFailedToRollback
}
