package controller

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/user/dto"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) GetOwnProfile(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		h.log.Error("failed to get user id from context")
		return errx.Unauthorized("failed to get user id from context")
	}

	user, err := h.userService.GetUserById(c, userId)
	if err != nil {
		h.log.WithError(err).Error("failed to get user by id")
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, user, "success get user profile")
}

func (h *UserHandler) UpdateOwnProfile(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		h.log.Error("failed to get user id from context")
		return errx.Unauthorized("failed to get user id from context")
	}

	var request dto.UpdateUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.WithError(err).Error("failed to validate request")
		return err
	}

	user, err := h.userService.UpdateUser(c, userId, request)
	if err != nil {
		h.log.WithError(err).Error("failed to update user")
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, user, "success update user profile")
}
