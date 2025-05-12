package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/entity"
)

func ToUserResponse(user *entity.User) dto.UserResponse {
	return dto.UserResponse{
		Id:               user.Id.String(),
		Email:            user.Email,
		FullName:         user.FullName,
		PhotoProfile:     user.PhotoProfile,
		IsVerified:       user.IsVerified,
		BodyWeight:       user.BodyWeight,
		DateOfBirth:      user.DateOfBirth.Format("2006-01-02"),
		Address:          user.Address,
		CaregiverContact: user.CaregiverContact,
	}
}
