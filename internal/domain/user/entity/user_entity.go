package entity

import (
	"database/sql"

	"github.com/Ndraaa15/foreglyc-server/pkg/enum"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	Id               uuid.UUID         `db:"id"`
	FullName         string            `db:"full_name"`
	Email            string            `db:"email"`
	Password         string            `db:"password"`
	PhotoProfile     string            `db:"photo_profile"`
	IsVerified       bool              `db:"is_verified"`
	BodyWeight       sql.NullFloat64   `db:"body_weight"`
	DateOfBirth      pq.NullTime       `db:"date_of_birth"`
	Address          sql.NullString    `db:"address"`
	CaregiverContact sql.NullString    `db:"caregiver_contact"`
	PhoneNumber      sql.NullString    `db:"phone_number"`
	AuthProvider     enum.AuthProvider `db:"auth_provider"`
	CreatedAt        pq.NullTime       `db:"created_at"`
	UpdatedAt        pq.NullTime       `db:"updated_at"`
}
