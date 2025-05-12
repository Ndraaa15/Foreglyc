package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/entity"
)

func ToUserResponse(user *entity.User) dto.UserResponse {
	dto := dto.UserResponse{
		Id:           user.Id.String(),
		Email:        user.Email,
		FullName:     user.FullName,
		PhotoProfile: user.PhotoProfile,
		IsVerified:   user.IsVerified,
	}

	if user.BodyWeight.Valid {
		dto.BodyWeight = user.BodyWeight.Float64
	}

	if user.DateOfBirth.Valid {
		dto.DateOfBirth = user.DateOfBirth.Time.Format("2006-01-02")
	}

	if user.Address.Valid {
		dto.Address = user.Address.String
	}

	if user.CaregiverContact.Valid {
		dto.CaregiverContact = user.CaregiverContact.String
	}

	return dto
}
