package http

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/chatbot/dto"
	"github.com/Ndraaa15/foreglyc-server/pkg/errx"
	"github.com/Ndraaa15/foreglyc-server/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *ChatBotHandler) ChatForeglycExpert(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 60*time.Second)
	defer cancel()

	var request []dto.ChatMessageRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	for i, msg := range request {
		if err := h.validator.Struct(msg); err != nil {
			h.log.WithError(err).Errorf("validation failed at index %d", i)
			return err
		}
	}

	// Process the chat request
	data, err := h.chatBotService.ChatForeglycExpert(c, request)
	if err != nil {
		h.log.WithError(err).Error("failed to process chat request")
		return err
	}

	// Return the response
	return response.SuccessResponse(ctx, fiber.StatusOK, data, "success generate response chat")
}

func (h *ChatBotHandler) GlucosePrediction(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 20*time.Second)
	defer cancel()

	userId, ok := ctx.Locals("userId").(string)
	if !ok {
		h.log.Error("failed to get user id from context")
		return errx.Unauthorized("failed to get user id from context")
	}

	data, err := h.chatBotService.GlucosePrediction(c, userId)
	if err != nil {
		h.log.WithError(err).Error("failed to process chat request")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, data, "success generate response chat")
}

func (h *ChatBotHandler) PredictionChatForeglycExpert(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 60*time.Second)
	defer cancel()

	var request dto.PredictionChatRequest
	if err := ctx.BodyParser(&request); err != nil {
		h.log.WithError(err).Error("failed to parse request")
		return err
	}

	for i, msg := range request.Chats {
		if err := h.validator.Struct(msg); err != nil {
			h.log.WithError(err).Errorf("validation failed at index %d", i)
			return err
		}
	}

	data, err := h.chatBotService.PredictionChatForeglycExpert(c, request)
	if err != nil {
		h.log.WithError(err).Error("failed to process chat request")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusOK, data, "success generate response chat")
}
