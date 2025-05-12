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
	WithTx(tx bool) (AuthRepositoryItf, error)
}

func New(db *sqlx.DB) RepositoryItf {
	return &Repository{db}
}

type AuthRepository struct {
	q sqlx.ExtContext
}

type AuthRepositoryItf interface {
	Commit() error
	Rollback() error

	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserById(ctx context.Context, id string) (entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
}

func (r *Repository) WithTx(tx bool) (AuthRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &AuthRepository{db}, nil
}

func (r *AuthRepository) Commit() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Commit()
	}

	return ErrFailedToCommit
}

func (r *AuthRepository) Rollback() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Rollback()
	}

	return ErrFailedToRollback
}
