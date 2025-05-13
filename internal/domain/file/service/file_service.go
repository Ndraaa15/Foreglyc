package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/file/dto"
	"github.com/spf13/viper"
)

func (s *FileService) UploadFile(
	ctx context.Context,
	requestFile *multipart.FileHeader,
) (dto.FileResponse, error) {
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

	if _, err := s.firebaseStorageService.UploadFile(ctx, fileData, filePath); err != nil {
		s.log.WithError(err).Error("failed to upload file")
		return dto.FileResponse{}, err
	}

	encodedPath := url.PathEscape(filePath)

	imageURLTemplate := viper.GetString("firebase.storage.image_url")
	publicURL := fmt.Sprintf(imageURLTemplate, encodedPath)

	return dto.FileResponse{Url: publicURL}, nil
}
