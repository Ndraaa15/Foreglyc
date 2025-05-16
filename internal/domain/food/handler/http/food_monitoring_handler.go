package http

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/food/dto"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *FoodHandler) GenerateFoodInformation(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 20*time.Second)
	defer cancel()

	var request dto.CreateFoodInformationRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.WithError(err).Error("failed to validate request")
		return err
	}

	res, err := h.foodService.GenerateFoodInformation(c, request)
	if err != nil {
		h.log.WithError(err).Error("failed to generate food information")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusCreated, res, "success generate food information")
}

func (h *FoodHandler) CreateFoodMonitoring(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.FoodMonitoringRequest
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

	res, err := h.foodService.CreateFoodMonitoring(c, request, userId)
	if err != nil {
		h.log.WithError(err).Error("failed to create food recall")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusCreated, res, "success create food recall")
}

func (h *FoodHandler) GetStatus3J(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		h.log.Error("failed to get userId from context")
		return errx.Unauthorized("user not found")
	}

	res, err := h.foodService.GetStatus3J(c, userId)
	if err != nil {
		h.log.WithError(err).Error("failed to get status 3J")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, res, "success get status 3J")
}

func (s *FoodHandler) GetFoodHomepage(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		s.log.Error("failed to get userId from context")
		return errx.Unauthorized("user not found")
	}

	res, err := s.foodService.GetFoodHomepage(c, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get food homepage")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, res, "success get food homepage")
}
