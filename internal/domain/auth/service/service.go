package service

import (
	"context"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/auth/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/auth/repository"
	userdto "github.com/Ndraaa15/foreglyc-server/internal/domain/user/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/cache"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/email"
	"github.com/sirupsen/logrus"
)

type IAuthService interface {
	SignUp(ctx context.Context, request dto.SignUpRequest) (userdto.UserResponse, error)
	SignIn(ctx context.Context, request dto.SignInRequest) (dto.SignInResponse, error)
	VerifyUser(ctx context.Context, id string, request dto.VerifyEmailRequest) (userdto.UserResponse, error)
	ResendVerificationEmail(ctx context.Context, id string) error

	ChangePassword(ctx context.Context, userId string, request dto.ChangePasswordRequest) (userdto.UserResponse, error)
	ForgotPassword(ctx context.Context, request dto.ForgotPasswordRequest) error
	VerifyForgotPassword(ctx context.Context, id string, request dto.VerifyForgotPasswordRequest) error
	ResendForgotPassword(ctx context.Context, request dto.ForgotPasswordRequest) error
}

type AuthService struct {
	log             *logrus.Logger
	authRepository  repository.RepositoryItf
	cacheRepository cache.ICacheRepository
	emailService    email.IEmail
}

func New(log *logrus.Logger, authRepository repository.RepositoryItf, cacheRepository cache.ICacheRepository, emailService email.IEmail) IAuthService {
	return &AuthService{
		log:             log,
		authRepository:  authRepository,
		cacheRepository: cacheRepository,
		emailService:    emailService,
	}
}
