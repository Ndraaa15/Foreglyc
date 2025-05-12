package http

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/auth/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	authService service.IAuthService
	log         *logrus.Logger
	validator   *validator.Validate
}

func New(authService service.IAuthService, log *logrus.Logger, validator *validator.Validate) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		log:         log,
		validator:   validator,
	}
}

func (c *AuthHandler) SetEndpoint(router *fiber.App) {
	v1 := router.Group("/api/v1/auth")
	v1.Post("/signup", c.SignUp)
	v1.Post("/signin", c.SignIn)
	v1.Post("/verify-email/:id", c.VerifyEmail)
	v1.Get("/resend-verification-email/:id", c.ResendVerificationEmail)

	v1.Post("/forgot-password", c.ForgotPassword)
	v1.Post("/verify-forgot-password", c.VerifyForgotPassword)
	v1.Post("/resend-forgot-password", c.ForgotPassword)
	v1.Post("/change-password", c.ChangePassword)
}
