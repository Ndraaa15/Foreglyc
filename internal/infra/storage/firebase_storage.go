package storage

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"

	"firebase.google.com/go/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type IFirebaseStorage interface {
	UploadFile(ctx context.Context, file []byte, fileName string) (string, error)
	GetFile(ctx context.Context, imageUrl string) (FileInformation, error)
}

type FirebaseStorage struct {
	client *storage.Client
	log    *logrus.Logger
}

func New(client *storage.Client, log *logrus.Logger) IFirebaseStorage {
	return &FirebaseStorage{client: client, log: log}
}

func (f *FirebaseStorage) UploadFile(ctx context.Context, file []byte, fileName string) (string, error) {
	bucket, err := f.client.Bucket(viper.GetString("firebase.storage.bucket"))
	if err != nil {
		return "", err
	}

	wc := bucket.Object(fileName).NewWriter(ctx)
	if _, err = wc.Write(file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	_, err = bucket.Object(fileName).Attrs(ctx)
	if err != nil {
		return "", err
	}

	link := fmt.Sprintf(viper.GetString("firebase.storage.image_url"), fileName)
	return link, nil
}

func (f *FirebaseStorage) GetFile(ctx context.Context, imageUrl string) (FileInformation, error) {
	// 1) Parse URL
	u, err := url.Parse(imageUrl)
	if err != nil {
		return FileInformation{}, fmt.Errorf("invalid image URL: %w", err)
	}

	// 2) Split off the object path after "/o/"
	parts := strings.SplitN(u.Path, "/o/", 2)
	if len(parts) < 2 {
		return FileInformation{}, fmt.Errorf("unable to extract object path from URL %q", u.Path)
	}

	// 3) Strip any query params (e.g. "?alt=media")
	rawObject := parts[1]
	if idx := strings.Index(rawObject, "?"); idx >= 0 {
		rawObject = rawObject[:idx]
	}

	// 4) URL-decode percent-escapes like "%2F"
	objectPath, err := url.QueryUnescape(rawObject)
	if err != nil {
		return FileInformation{}, fmt.Errorf("failed to decode object path %q: %w", rawObject, err)
	}

	// 5) Get bucket and object handle
	bucketName := viper.GetString("firebase.storage.bucket")
	bucket, err := f.client.Bucket(bucketName)
	if err != nil {
		return FileInformation{}, fmt.Errorf("failed to get bucket %q: %w", bucketName, err)
	}
	obj := bucket.Object(objectPath)

	// 6) Fetch metadata (size, MIME type)
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return FileInformation{}, fmt.Errorf("failed to fetch object attrs for %q: %w", objectPath, err)
	}

	// 7) Read contents
	rc, err := obj.NewReader(ctx)
	if err != nil {
		return FileInformation{}, fmt.Errorf("failed to open reader for %q: %w", objectPath, err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return FileInformation{}, fmt.Errorf("failed to read data for %q: %w", objectPath, err)
	}

	// 8) Build result
	return FileInformation{
		Name: path.Base(objectPath), // e.g. "myfile.png"
		Size: attrs.Size,
		Type: attrs.ContentType,
		Data: data,
	}, nil
}
