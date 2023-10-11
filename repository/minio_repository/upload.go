package minio_repository

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (m *MinioRepositoryImpl) UploadFile(ctx context.Context, file *multipart.FileHeader, objectName string) error {
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

	log.Info().Msgf(objectName)
	info, err := m.c.PutObject(ctx, infra.MinIoBucket, objectName, fileReader, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Warn().Msgf(util.LogErrPutObjectMinio, err)
		return err
	}

	log.Info().Msgf(util.LogInfoFileUploadMinio, info)
	return nil
}
