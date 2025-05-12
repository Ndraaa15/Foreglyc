package controller

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/service"
	"github.com/Ndraaa15/foreglyc-server/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService service.IUserService
	log         *logrus.Logger
	validator   *validator.Validate
}

func New(userService service.IUserService, log *logrus.Logger, validator *validator.Validate) *UserHandler {
	return &UserHandler{
		userService: userService,
		log:         log,
		validator:   validator,
	}
}

func (c *UserHandler) SetEndpoint(router *fiber.App) {
	v1 := router.Group("/api/v1/users")
	v1.Use(middleware.Authentication())
	v1.Get("/self", middleware.Authentication(), c.GetOwnProfile)
	v1.Put("/self", middleware.Authentication(), c.UpdateOwnProfile)
}
