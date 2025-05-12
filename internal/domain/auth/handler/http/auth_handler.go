package http

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/auth/dto"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *AuthHandler) SignUp(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.SignUpRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.WithError(err).Error("failed to validate request")
		return err
	}

	user, err := h.authService.SignUp(c, request)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusCreated, user, "success sign up")
}

func (h *AuthHandler) SignIn(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.SignInRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.WithError(err).Error("failed to validate request")
		return err
	}

	token, err := h.authService.SignIn(c, request)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, token, "success sign in")
}

func (h *AuthHandler) VerifyEmail(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.VerifyEmailRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.WithError(err).Error("failed to validate request")
		return err
	}

	user, err := h.authService.VerifyUser(c, ctx.Params("id"), request)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, user, "success verify email")
}

func (h *AuthHandler) ResendVerificationEmail(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	err := h.authService.ResendVerificationEmail(c, ctx.Params("id"))
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.NoContentResponse(ctx)
}

func (h *AuthHandler) ForgotPassword(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.ForgotPasswordRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.WithError(err).Error("failed to validate request")
		return err
	}

	err := h.authService.ForgotPassword(c, request)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.NoContentResponse(ctx)
}

func (h *AuthHandler) ChangePassword(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.ChangePasswordRequest
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

	resp, err := h.authService.ChangePassword(c, userId, request)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, resp, "success change password")
}

func (h *AuthHandler) VerifyForgotPassword(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		h.log.Error("failed to get userId from context")
		return errx.Unauthorized("user not found in context")
	}

	var request dto.VerifyForgotPasswordRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.WithError(err).Error("failed to validate request")
		return err
	}

	err := h.authService.VerifyForgotPassword(c, userId, request)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.NoContentResponse(ctx)
}

func (h *AuthHandler) ResendForgotPassword(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var request dto.ForgotPasswordRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	if err := h.validator.Struct(request); err != nil {
		h.log.WithError(err).Error("failed to validate request")
		return err
	}

	err := h.authService.ResendForgotPassword(c, request)
	if err != nil {
		return err
	}

	select {
	case <-c.Done():
		h.log.WithError(c.Err()).Error("timeout context")
		return errx.Timeout("request timeout")
	default:
	}

	return response.NoContentResponse(ctx)
}
