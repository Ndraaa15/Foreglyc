package http

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/service"
	"github.com/Ndraaa15/foreglyc-server/internal/middleware"
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
	v1 := router.Group("/api/v1/monitorings")

	v1.Post("/cgms/preferences", middleware.Authentication(), c.CreateCGMMonitoringPreference)
	v1.Post("/glucometers/preferences", middleware.Authentication(), c.CreateGlucometerMonitoringPreference)

	v1.Post("/glucometers", middleware.Authentication(), c.CreateGlucometerMonitoring)
	v1.Get("/glucometers", middleware.Authentication(), c.GetGlucometerMonitorings)
	v1.Get("/glucometers/graph", middleware.Authentication(), c.GetGlucometerMonitoringGraph)

	v1.Post("/questionnaires", middleware.Authentication(), c.CreateMonitoringQuiestionnare)
}
