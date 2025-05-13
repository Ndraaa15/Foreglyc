package http

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (h *FileHandler) UploadFile(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.UserContext(), 10*time.Second)
	defer cancel()

	file, err := ctx.FormFile("file")
	if err != nil {
		h.log.WithError(err).Error("failed to get file from request")
		return err
	}

	resp, err := h.FileService.UploadFile(c, file)
	if err != nil {
		h.log.WithError(err).Error("failed to upload file")
		return err
	}

	return response.SuccessResponse(ctx, fiber.StatusCreated, resp, "success upload file")
}
