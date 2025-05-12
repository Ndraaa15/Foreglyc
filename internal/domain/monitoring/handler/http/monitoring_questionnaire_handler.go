package http

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *MonitoringHandler) CreateMonitoringQuiestionnare(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.CreateMonitoringQuestionnaire
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

	questionnare, err := h.monitoringService.CreateMonitoringQuestionnaire(c, request, userId)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, questionnare, "success create monitoring questionnare")
}
