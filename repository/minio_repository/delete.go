package minio_repository

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (m *MinioRepositoryImpl) DeleteFile(ctx context.Context, objectName string) error {
	if err := m.c.RemoveObject(ctx, infra.MinIoBucket, objectName, minio.RemoveObjectOptions{}); err != nil {
		log.Warn().Msgf(util.LogErrDelObjectMinio, err)
		return err
	}

	return nil
}
