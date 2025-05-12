package entity

import (
	"time"

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
	BodyWeight       float64           `db:"body_weight"`
	DateOfBirth      time.Time         `db:"date_of_birth"`
	Address          string            `db:"address"`
	CaregiverContact string            `db:"caregiver_contact"`
	PhoneNumber      string            `db:"phone_number"`
	AuthProvider     enum.AuthProvider `db:"auth_provider"`
	CreatedAt        pq.NullTime       `db:"created_at"`
	UpdatedAt        pq.NullTime       `db:"updated_at"`
}
