package domain

import (
	"context"
	"mime/multipart"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o ./mocks . MinioRepo
type MinioRepo interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, objectName string) error
	DeleteFile(ctx context.Context, objectName string) error
	GenerateFileName(fileExt string, path string) string
}
