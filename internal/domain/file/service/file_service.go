package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/file/dto"
	"github.com/spf13/viper"
)

func (s *FileService) UploadFile(ctx context.Context, requestFile *multipart.FileHeader) (dto.FileResponse, error) {
	file, err := requestFile.Open()
	if err != nil {
		return dto.FileResponse{}, err
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return dto.FileResponse{}, err
	}

	folder := viper.GetString("firebase.storage.folder")
	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), requestFile.Filename)
	filePath := fmt.Sprintf("%s/%s", folder, fileName)

	_, err = s.firebaseStorageService.UploadFile(ctx, fileData, filePath)
	if err != nil {
		s.log.WithError(err).Error("failed to upload file")
		return dto.FileResponse{}, err
	}

	lastUri := fmt.Sprintf("%%2F%s?alt=media", fileName)
	return dto.FileResponse{
		Url: fmt.Sprintf(viper.GetString("firebase.storage.image_url"), lastUri),
	}, nil
}
