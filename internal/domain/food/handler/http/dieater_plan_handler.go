package http

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *FoodHandler) CreateDietaryPlan(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.CreateDietaryPlanRequest
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

	res, err := h.foodService.CreateDietaryPlan(c, request, userId)
	if err != nil {
		h.log.WithError(err).Error("failed to create dietary plan")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusCreated, res, "success create dietary plan")
}

func (h *FoodHandler) UpdateInsulineQuestionnaire(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.UpdateInsulineQuestionnaireRequest
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

	res, err := h.foodService.UpdateInsulineQuestionnaire(c, request, userId)
	if err != nil {
		h.log.WithError(err).Error("failed to update insuline questionnaire")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, res, "success update insuline questionnaire")
}
