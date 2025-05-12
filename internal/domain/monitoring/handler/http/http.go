package http

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MonitoringHandler struct {
	monitoringService service.IMonitoringService
	log               *logrus.Logger
	validator         *validator.Validate
}

func New(monitoringService service.IMonitoringService, log *logrus.Logger, validator *validator.Validate) *MonitoringHandler {
	return &MonitoringHandler{
		monitoringService: monitoringService,
		log:               log,
		validator:         validator,
	}
}

func (c *MonitoringHandler) SetEndpoint(router *fiber.App) {
	_ = router.Group("/api/v1/monitorings")
}
