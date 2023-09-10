package minio

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/config"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
)

type MinioImpl struct {
	Bucket string
}

func NewMinioImpl() *MinioImpl {
	return &MinioImpl{
		Bucket: config.Get().ThirdParty.Minio.Bucket,
	}
}

func (m *MinioImpl) UploadFile(ctx context.Context, file *multipart.FileHeader, objectName string) error {
	fileReader, err := file.Open()
	if err != nil {
		log.Err(err).Msg("cannot open file header")
		return err
	}
	defer fileReader.Close()

	minioClient, err := Connection()
	if err != nil {
		return err
	}
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	info, err := minioClient.PutObject(ctx, m.Bucket, objectName, fileReader, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Err(err).Msg("cannot put object")
		return err
	}

	log.Info().Msgf("info upload : %v", info)

	return nil
}

func (m *MinioImpl) DeleteFile(ctx context.Context, objectName string) error {
	minioClient, err := Connection()
	if err != nil {
		return err
	}

	if err := minioClient.RemoveObject(ctx, m.Bucket, objectName, minio.RemoveObjectOptions{}); err != nil {
		log.Err(err).Msg("cannot remove object")
		return err
	}

	return nil
}

func (m *MinioImpl) GenerateFileName(file *multipart.FileHeader, path, prefix string) string {
	nameFile := fmt.Sprintf("%s%s%d%s", path, prefix, time.Now().UnixNano(), filepath.Ext(file.Filename))
	return nameFile
}
