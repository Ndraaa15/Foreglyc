package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/auth/dto"
	userdto "github.com/Ndraaa15/foreglyc-server/internal/domain/user/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/entity"
	usermapper "github.com/Ndraaa15/foreglyc-server/internal/domain/user/mapper"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"

	"github.com/Ndraaa15/foreglyc-server/internal/infra/cache"
	"github.com/Ndraaa15/foreglyc-server/pkg/constant"
	"github.com/Ndraaa15/foreglyc-server/pkg/enum"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/jwt"
	"github.com/Ndraaa15/foreglyc-server/pkg/util"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) SignUp(ctx context.Context, request dto.SignUpRequest) (userdto.UserResponse, error) {
	authRepository, err := s.authRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create initialize client")
		return userdto.UserResponse{}, err
	}

	uuidV7, err := uuid.NewV7()
	if err != nil {
		s.log.WithError(err).Error("failed to generate uuid")
		return userdto.UserResponse{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed to hash password")
		return userdto.UserResponse{}, err
	}

	user := entity.User{
		Id:           uuidV7,
		Email:        request.Email,
		Password:     string(hashedPassword),
		FullName:     request.FullName,
		IsVerified:   false,
		PhotoProfile: "http://www.gravatar.com/avatar/?d=mp",
		AuthProvider: enum.AuthProviderBasic,
		Level:        constant.UserLevelGluMaster,
		CreatedAt:    pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = authRepository.CreateUser(ctx, &user)
	if err != nil {
		s.log.WithError(err).Error("failed to create user")
		return userdto.UserResponse{}, err
	}

	verificationCode := util.GenerateCode(4)
	err = s.cacheRepository.Set(ctx, fmt.Sprintf("%s:sign_up", user.Id.String()), verificationCode, cache.DefaultExpiration)
	if err != nil {
		s.log.WithError(err).Error("failed to set verification code in cache")
		return userdto.UserResponse{}, err
	}

	s.emailService.SetReciever(user.Email)
	s.emailService.SetSubject("Verification Code")
	s.emailService.SetSender(viper.GetString("email.from"))
	s.emailService.SetBodyHTML("verification_email.html",
		struct {
			OTP string
		}{
			OTP: verificationCode,
		},
	)

	err = s.emailService.Send()
	if err != nil {
		s.log.WithError(err).Error("failed to send email")
		return userdto.UserResponse{}, err
	}

	return usermapper.ToUserResponse(&user), nil
}

func (s *AuthService) SignIn(ctx context.Context, request dto.SignInRequest) (dto.SignInResponse, error) {
	authRepository, err := s.authRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create initialize client")
		return dto.SignInResponse{}, err
	}

	user, err := authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		s.log.WithError(err).Error("failed to get user by email")
		return dto.SignInResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		s.log.WithError(err).Error("failed to compare password")
		return dto.SignInResponse{}, errx.BadRequest("invalid password or email")
	}

	accessToken, err := jwt.EncodeToken(&user)
	if err != nil {
		s.log.WithError(err).Error("failed to encode token")
		return dto.SignInResponse{}, err
	}

	return dto.SignInResponse{
		TokenType:   jwt.TokenTypeBearer,
		AccessToken: accessToken,
		ExpiresAt:   int64(viper.GetDuration("jwt.expiration").Seconds()),
	}, nil
}

func (s *AuthService) VerifyUser(ctx context.Context, id string, request dto.VerifyEmailRequest) (userdto.UserResponse, error) {
	verificationCode, err := s.cacheRepository.Get(ctx, fmt.Sprintf("%s:sign_up", id))
	if err != nil {
		return userdto.UserResponse{}, errx.Timeout("verification code expired")
	}

	if verificationCode != request.Code {
		s.log.Error("invalid verification code")
		return userdto.UserResponse{}, errx.BadRequest("invalid verification code")
	}

	authRepository, err := s.authRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create initialize client")
		return userdto.UserResponse{}, err
	}

	user, err := authRepository.GetUserById(ctx, id)
	if err != nil {
		s.log.WithError(err).Error("failed to get user by id")
		return userdto.UserResponse{}, err
	}

	user.IsVerified = true
	user.UpdatedAt = pq.NullTime{Time: time.Now(), Valid: true}

	err = authRepository.UpdateUser(ctx, &user)
	if err != nil {
		s.log.WithError(err).Error("failed to update user")
		return userdto.UserResponse{}, err
	}

	err = s.cacheRepository.Delete(ctx, fmt.Sprintf("%s:sign_up", id))
	if err != nil {
		return userdto.UserResponse{}, err
	}

	return usermapper.ToUserResponse(&user), nil
}

func (s *AuthService) ResendVerificationEmail(ctx context.Context, id string) error {
	repository, err := s.authRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create transaction")
		return err
	}

	user, err := repository.GetUserById(ctx, id)
	if err != nil {
		s.log.WithError(err).Error("failed to get user by id")
		return err
	}

	if user.IsVerified {
		s.log.Error("user already verified")
		return errx.BadRequest("user already verified")
	}

	verificationCode := util.GenerateCode(4)
	_, err = s.cacheRepository.Get(ctx, fmt.Sprintf("%s:sign_up", id))
	if errors.Is(err, redis.Nil) {
		if err := s.cacheRepository.Set(ctx, fmt.Sprintf("%s:sign_up", id), verificationCode, cache.DefaultExpiration); err != nil {
			return err
		}
	} else if err != nil {
		s.log.WithError(err).Error("failed to get verification code")
		return err
	} else {
		if err := s.cacheRepository.Delete(ctx, fmt.Sprintf("%s:sign_up", id)); err != nil {
			return err
		}
		if err := s.cacheRepository.Set(ctx, fmt.Sprintf("%s:sign_up", id), verificationCode, cache.DefaultExpiration); err != nil {
			return err
		}
	}

	s.emailService.SetReciever(user.Email)
	s.emailService.SetSubject("Verification Code")
	s.emailService.SetSender(viper.GetString("email.from"))
	s.emailService.SetBodyHTML("verification_email.html",
		struct {
			OTP string
		}{
			OTP: verificationCode,
		},
	)

	err = s.emailService.Send()
	if err != nil {
		s.log.WithError(err).Error("failed to send email")
		return err
	}

	return nil
}

func (s *AuthService) ChangePassword(ctx context.Context, userId string, request dto.ChangePasswordRequest) (userdto.UserResponse, error) {
	authRepository, err := s.authRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create initialize client")
		return userdto.UserResponse{}, err
	}

	user, err := authRepository.GetUserById(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get user by id")
		return userdto.UserResponse{}, err
	}

	if request.ConfirmNewPassword != request.NewPassword {
		s.log.Error("new password and confirm new password do not match")
		return userdto.UserResponse{}, errx.BadRequest("new password and confirm new password do not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed to hash password")
		return userdto.UserResponse{}, err
	}

	user.Password = string(hashedPassword)
	user.UpdatedAt = pq.NullTime{Time: time.Now(), Valid: true}

	err = authRepository.UpdateUser(ctx, &user)
	if err != nil {
		s.log.WithError(err).Error("failed to update user")
		return userdto.UserResponse{}, err
	}

	return usermapper.ToUserResponse(&user), nil
}

func (s *AuthService) ForgotPassword(ctx context.Context, request dto.ForgotPasswordRequest) error {
	authRepository, err := s.authRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create initialize client")
		return err
	}

	user, err := authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		s.log.WithError(err).Error("failed to get user by email")
		return err
	}

	resetCode := util.GenerateCode(4)
	err = s.cacheRepository.Set(ctx, fmt.Sprintf("%s:reset_code", user.Id.String()), resetCode, cache.DefaultExpiration)
	if err != nil {
		s.log.WithError(err).Error("failed to create initialize client")
		return err
	}

	s.emailService.SetReciever(user.Email)
	s.emailService.SetSubject("Reset Password")
	s.emailService.SetSender(viper.GetString("email.from"))
	s.emailService.SetBodyHTML("reset_password_email.html",
		struct {
			OTP string
		}{
			OTP: resetCode,
		},
	)

	err = s.emailService.Send()
	if err != nil {
		s.log.WithError(err).Error("failed to send email")
		return err
	}

	return nil
}

func (s *AuthService) VerifyForgotPassword(ctx context.Context, id string, request dto.VerifyForgotPasswordRequest) error {
	resetCode, err := s.cacheRepository.Get(ctx, fmt.Sprintf("%s:reset_code", id))
	if errors.Is(err, redis.Nil) {
		return errx.Timeout("reset code expired")
	} else if err != nil {
		return err
	}

	if resetCode != request.Code {
		s.log.Error("invalid reset code")
		return errx.BadRequest("invalid reset code")
	}

	return nil
}

func (s *AuthService) ResendForgotPassword(ctx context.Context, request dto.ForgotPasswordRequest) error {
	repository, err := s.authRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create transaction")
		return err
	}

	user, err := repository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		s.log.WithError(err).Error("failed to get user by email")
		return err
	}

	resetCode := util.GenerateCode(4)
	_, err = s.cacheRepository.Get(ctx, fmt.Sprintf("%s:reset_code", user.Id.String()))
	if errors.Is(err, redis.Nil) {
		if err := s.cacheRepository.Set(ctx, fmt.Sprintf("%s:reset_code", user.Id.String()), resetCode, cache.DefaultExpiration); err != nil {
			return err
		}
	} else if err != nil {
		s.log.WithError(err).Error("failed to get reset code")
		return err
	} else {
		if err := s.cacheRepository.Delete(ctx, fmt.Sprintf("%s:reset_code", user.Id.String())); err != nil {
			return err
		}
		if err := s.cacheRepository.Set(ctx, fmt.Sprintf("%s:reset_code", user.Id.String()), resetCode, cache.DefaultExpiration); err != nil {
			return err
		}
	}

	s.emailService.SetReciever(user.Email)
	s.emailService.SetSubject("Reset Password")
	s.emailService.SetSender(viper.GetString("email.from"))
	s.emailService.SetBodyHTML("reset_password_email.html",
		struct {
			OTP string
		}{
			OTP: resetCode,
		},
	)

	err = s.emailService.Send()
	if err != nil {
		s.log.WithError(err).Error("failed to send email")
		return err
	}

	return nil
}
