package service

import (
	"context"
	"mime/multipart"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/file/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/infra/storage"
	"github.com/sirupsen/logrus"
)

type IFileService interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) (dto.FileResponse, error)
}

type FileService struct {
	log                    *logrus.Logger
	firebaseStorageService storage.IFirebaseStorage
}

func New(log *logrus.Logger, firebaseStorageService storage.IFirebaseStorage) IFileService {
	return &FileService{
		log:                    log,
		firebaseStorageService: firebaseStorageService,
	}
}
