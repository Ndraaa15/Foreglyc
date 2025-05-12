package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/repository"
	"github.com/sirupsen/logrus"
)

type IUserService interface {
	GetUserById(ctx context.Context, userId string) (dto.UserResponse, error)
	UpdateUser(ctx context.Context, userId string, request dto.UpdateUserRequest) (dto.UserResponse, error)
}

type UserService struct {
	log            *logrus.Logger
	UserRepository repository.RepositoryItf
}

func New(log *logrus.Logger, UserRepository repository.RepositoryItf) IUserService {
	return &UserService{
		log:            log,
		UserRepository: UserRepository,
	}
}
