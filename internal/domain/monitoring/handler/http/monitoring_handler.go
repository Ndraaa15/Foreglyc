package http

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *MonitoringHandler) CreateGlucometerMonitoring(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.CreateGlucometerMonitoringRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.WithError(err).Error("failed to validate request")
		return err
	}

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		h.log.Error("failed to get userId from context")
		return errx.Unauthorized("user not found")
	}

	monitoring, err := h.monitoringService.CreateGlucometerMonitoring(c, request, userId)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusCreated, monitoring, "success create glucometer monitoring")
}

func (h *MonitoringHandler) GetGlucometerMonitorings(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		h.log.Error("failed to get userId from context")
		return errx.Unauthorized("user not found")
	}

	var filter dto.GetGlucometerMonitoringFilter
	if err := ctx.QueryParser(&filter); err != nil {
		h.log.WithError(err).Error("failed to parse query")
		return err
	}

	filter.UserId = userId

	monitoring, err := h.monitoringService.GetGlucometerMonitorings(c, filter)
	if err != nil {
		h.log.WithError(err).Error("failed to get glucometer monitoring")
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, monitoring, "success get glucometer monitoring")
}

func (h *MonitoringHandler) GetGlucometerMonitoringGraph(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		h.log.Error("failed to get userId from context")
		return errx.Unauthorized("user not found")
	}

	var filter dto.GetGlucometerMonitoringGraphFilter
	if err := ctx.QueryParser(&filter); err != nil {
		h.log.WithError(err).Error("failed to parse query")
		return err
	}

	filter.UserId = userId

	monitoring, err := h.monitoringService.GetGlucometerMonitorignGraph(c, filter)
	if err != nil {
		h.log.WithError(err).Error("failed to get glucometer monitoring graph")
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, monitoring, "success get glucometer monitoring graph")
}
