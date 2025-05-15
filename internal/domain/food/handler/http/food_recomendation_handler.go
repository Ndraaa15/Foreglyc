package http

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *FoodHandler) GetFoodRecomendation(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 180*time.Second)
	defer cancel()

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		h.log.Error("failed to get userId from context")
		return errx.Unauthorized("user not found")
	}

	res, err := h.foodService.GenerateFoodRecommendation(c, userId)
	if err != nil {
		h.log.WithError(err).Error("failed to get food recomendation")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, res, "success get food recomendation")
}
