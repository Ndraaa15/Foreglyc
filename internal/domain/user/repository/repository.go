package repository

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/entity"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/jmoiron/sqlx"
)

var (
	ErrFailedToCommit   = errx.InternalServerError("failed to commit transaction")
	ErrFailedToRollback = errx.InternalServerError("failed to rollback transaction")
)

const (
	UserTableName = "users"
)

type Repository struct {
	DB *sqlx.DB
}

type RepositoryItf interface {
	WithTx(tx bool) (UserRepositoryItf, error)
}

func New(db *sqlx.DB) RepositoryItf {
	return &Repository{db}
}

type UserRepository struct {
	q sqlx.ExtContext
}

type UserRepositoryItf interface {
	Commit() error
	Rollback() error

	GetUserById(ctx context.Context, userId string) (entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
}

func (r *Repository) WithTx(tx bool) (UserRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &UserRepository{db}, nil
}

func (r *UserRepository) Commit() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Commit()
	}

	return ErrFailedToCommit
}

func (r *UserRepository) Rollback() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Rollback()
	}

	return ErrFailedToRollback
}
