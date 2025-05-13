package http

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/file/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FileHandler struct {
	FileService service.IFileService
	log         *logrus.Logger
	validator   *validator.Validate
}

func New(FileService service.IFileService, log *logrus.Logger, validator *validator.Validate) *FileHandler {
	return &FileHandler{
		FileService: FileService,
		log:         log,
		validator:   validator,
	}
}

func (c *FileHandler) SetEndpoint(router *fiber.App) {
	v1 := router.Group("/api/v1/files")
	v1.Post("/upload", c.UploadFile)

}
