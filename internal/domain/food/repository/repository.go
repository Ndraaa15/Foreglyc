package repository

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/jmoiron/sqlx"
)

var (
	ErrFailedToCommit   = errx.InternalServerError("failed to commit transaction")
	ErrFailedToRollback = errx.InternalServerError("failed to rollback transaction")
)

const (
	DietaryPlanTable    = "dietary_plans"
	FoodMonitoringTable = "food_monitorings"
)

type Repository struct {
	DB *sqlx.DB
}

type RepositoryItf interface {
	WithTx(tx bool) (FoodRepositoryItf, error)
}

func New(db *sqlx.DB) RepositoryItf {
	return &Repository{db}
}

type FoodRepository struct {
	q sqlx.ExtContext
}

type FoodRepositoryItf interface {
	Commit() error
	Rollback() error

	CreateDietaryPlan(ctx context.Context, dietaryPlan *entity.DietaryPlan) error
	GetDietaryPlan(ctx context.Context, userId string) (entity.DietaryPlan, error)
	UpdateDietaryPlan(ctx context.Context, dietaryPlan *entity.DietaryPlan) error

	CreateFoodMonitoring(ctx context.Context, foodRecall *entity.FoodMonitoring) error
	CountFoodMonitoring(ctx context.Context, userId string) (int64, error)
}

func (r *Repository) WithTx(tx bool) (FoodRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &FoodRepository{db}, nil
}

func (r *FoodRepository) Commit() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Commit()
	}

	return ErrFailedToCommit
}

func (r *FoodRepository) Rollback() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Rollback()
	}

	return ErrFailedToRollback
}
