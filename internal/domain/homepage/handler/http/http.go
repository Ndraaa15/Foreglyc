package http

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/homepage/service"
	"github.com/Ndraaa15/foreglyc-server/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HomepageHandler struct {
	homepageService service.IHomepageService
	log             *logrus.Logger
	validator       *validator.Validate
}

func New(homepageService service.IHomepageService, log *logrus.Logger, validator *validator.Validate) *HomepageHandler {
	return &HomepageHandler{
		homepageService: homepageService,
		log:             log,
		validator:       validator,
	}
}

func (c *HomepageHandler) SetEndpoint(router *fiber.App) {
	v1 := router.Group("/api/v1/homepages")
	v1.Get("/self", middleware.Authentication(), c.GetHomepage)
}
