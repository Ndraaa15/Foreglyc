package service

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/mapper"
)

func (s *UserService) GetUserById(ctx context.Context, userId string) (dto.UserResponse, error) {
	repository, err := s.UserRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create initialize client")
		return dto.UserResponse{}, err
	}

	user, err := repository.GetUserById(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get user by id")
		return dto.UserResponse{}, err
	}

	return mapper.ToUserResponse(&user), nil
}

func (s *UserService) UpdateUser(ctx context.Context, userId string, request dto.UpdateUserRequest) (dto.UserResponse, error) {
	repository, err := s.UserRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create initialize client")
		return dto.UserResponse{}, err
	}

	user, err := repository.GetUserById(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get user by id")
		return dto.UserResponse{}, err
	}

	dateOfBirth, err := time.Parse("2006-01-02", request.DateOfBirth)
	if err != nil {
		s.log.WithError(err).Error("failed to parse date of birth")
		return dto.UserResponse{}, err
	}

	user.FullName = request.FullName
	user.Email = request.Email
	user.PhotoProfile = request.PhotoProfile
	user.BodyWeight = request.BodyWeight
	user.DateOfBirth = dateOfBirth
	user.CaregiverContact = request.CaregiverContact

	err = repository.UpdateUser(ctx, &user)
	if err != nil {
		s.log.WithError(err).Error("failed to update user")
		return dto.UserResponse{}, err
	}

	return mapper.ToUserResponse(&user), nil
}
