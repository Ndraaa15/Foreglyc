package http

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/chatbot/service"
	"github.com/Ndraaa15/foreglyc-server/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ChatBotHandler struct {
	chatBotService service.IChatBotService
	log            *logrus.Logger
	validator      *validator.Validate
}

func New(chatBotService service.IChatBotService, log *logrus.Logger, validator *validator.Validate) *ChatBotHandler {
	return &ChatBotHandler{
		chatBotService: chatBotService,
		log:            log,
		validator:      validator,
	}
}

func (h *ChatBotHandler) SetEndpoint(router *fiber.App) {
	v1 := router.Group("/api/v1/chatbots")
	v1.Post("/foreglyc-expert", middleware.Authentication(), h.ChatForeglycExpert)
	v1.Get("/glucoses/prediction", middleware.Authentication(), h.GlucosePrediction)
	v1.Post("/glucoses/chat/prediction", middleware.Authentication(), h.PredictionChatForeglycExpert)
}
