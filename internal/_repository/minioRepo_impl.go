package _repository

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra/config"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type MinioImpl struct {
	c *minio.Client
}

func NewMinioImpl(c *minio.Client) domain.MinioRepo {
	return &MinioImpl{
		c: c,
	}
}

func (m *MinioImpl) UploadFile(ctx context.Context, file *multipart.FileHeader, objectName string) error {
	fileReader, err := file.Open()
	if err != nil {
		log.Warn().Msgf(util.LogErrOpenFile, file, err)
		return err
	}
	defer func() {
		errClose := fileReader.Close()
		if errClose != nil {
			log.Warn().Msgf(util.LogErrCloseFile, file, errClose)
		}
	}()

	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	info, err := m.c.PutObject(ctx, config.MinIoBucket, objectName, fileReader, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Warn().Msgf(util.LogErrPutObjectMinio, err)
		return err
	}

	log.Info().Msgf(util.LogInfoFileUploadMinio, info)
	return nil
}

func (m *MinioImpl) DeleteFile(ctx context.Context, objectName string) error {
	if err := m.c.RemoveObject(ctx, config.MinIoBucket, objectName, minio.RemoveObjectOptions{}); err != nil {
		log.Warn().Msgf(util.LogErrDelObjectMinio, err)
		return err
	}

	return nil
}

func (m *MinioImpl) GenerateFileName(fileExt string, path string) string {
	nameFile := fmt.Sprintf("/%s/%s%d%s", config.MinIoBucket, path, time.Now().UnixNano(), fileExt)
	return nameFile
}
