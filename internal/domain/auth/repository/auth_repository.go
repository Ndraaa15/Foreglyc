package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/entity"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func (r *AuthRepository) CreateUser(ctx context.Context, user *entity.User) error {
	query, values, err := squirrel.
		Insert(UserTableName).
		Columns("id", "email", "password", "full_name", "photo_profile", "is_verified", "auth_provider", "level", "created_at").
		Values(user.Id, user.Email, user.Password, user.FullName, user.PhotoProfile, user.IsVerified, user.AuthProvider, user.Level, user.CreatedAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}
	query = r.q.Rebind(query)

	_, err = r.q.ExecContext(ctx, query, values...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" && pqErr.Constraint == "users_email_key" {
				return errx.Conflict("email already exists")
			}
		}

		return err
	}

	return nil
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	query, values, err := squirrel.
		Select("id", "email", "password", "full_name", "photo_profile", "created_at", "updated_at").
		From(UserTableName).
		Where("email = ?", email).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return entity.User{}, err
	}
	query = r.q.Rebind(query)

	var user entity.User
	err = sqlx.GetContext(ctx, r.q, &user, query, values...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errx.NotFound("user not found")
		}
		return entity.User{}, err
	}

	return user, nil
}

func (r *AuthRepository) GetUserById(ctx context.Context, id string) (entity.User, error) {
	query, values, err := squirrel.
		Select("id", "email", "password", "full_name", "photo_profile", "created_at", "updated_at").
		From(UserTableName).
		Where("id = ?", id).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return entity.User{}, err
	}
	query = r.q.Rebind(query)

	var user entity.User
	err = sqlx.GetContext(ctx, r.q, &user, query, values...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errx.NotFound("user not found")
		}
		return entity.User{}, err
	}

	return user, nil
}

func (r *AuthRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	query, values, err := squirrel.
		Update(UserTableName).
		Set("email", user.Email).
		Set("full_name", user.FullName).
		Set("photo_profile", user.PhotoProfile).
		Set("is_verified", user.IsVerified).
		Set("updated_at", user.UpdatedAt).
		Where("id = ?", user.Id).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}
	query = r.q.Rebind(query)

	_, err = r.q.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
