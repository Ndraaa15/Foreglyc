package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/entity"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/jmoiron/sqlx"
)

func (r *UserRepository) GetUserById(ctx context.Context, userId string) (entity.User, error) {
	query, args, err := squirrel.
		Select(
			"id",
			"full_name",
			"email",
			"password",
			"photo_profile",
			"is_verified",
			"body_weight",
			"date_of_birth",
			"address",
			"caregiver_contact",
			"auth_provider",
			"created_at",
			"updated_at",
		).
		From(UserTableName).
		Where("id = ?", userId).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return entity.User{}, err
	}

	query = r.q.Rebind(query)

	var user entity.User
	err = sqlx.GetContext(ctx, r.q, &user, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errx.NotFound("user not found")
		}
		return entity.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	query, args, err := squirrel.
		Update(UserTableName).
		Set("full_name", user.FullName).
		Set("email", user.Email).
		Set("password", user.Password).
		Set("photo_profile", user.PhotoProfile).
		Set("is_verified", user.IsVerified).
		Set("body_weight", user.BodyWeight).
		Set("date_of_birth", user.DateOfBirth).
		Set("address", user.Address).
		Set("caregiver_contact", user.CaregiverContact).
		Set("auth_provider", user.AuthProvider).
		Where("id = ?", user.Id).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	query = r.q.Rebind(query)

	result, err := r.q.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errx.NotFound("user not found")
	}

	return nil
}
