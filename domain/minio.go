package domain

import (
	"context"
	"mime/multipart"
)

type MinioRepo interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, objectName string) error
	DeleteFile(ctx context.Context, objectName string) error
	GenerateFileName(fileExt string, path string) string
}
