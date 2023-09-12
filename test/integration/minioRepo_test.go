package integration

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra/config"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/_repository"
)

func createBucket(t *testing.T) {
	err := minioClient.MakeBucket(context.Background(), "files", minio.MakeBucketOptions{})
	config.MinIoBucket = "files"
	assert.NoError(t, err)
}

func minioRepo(t *testing.T) {
	fileHeader := newFileHeader()
	fileExt := filepath.Ext(fileHeader.Filename)
	ctx := context.TODO()
	minioRepo := _repository.NewMinioImpl(minioClient)

	fileName := minioRepo.GenerateFileName(fileExt, "payment-images/public/")

	t.Run("SUCCESS_upload", func(t *testing.T) {
		err := minioRepo.UploadFile(ctx, fileHeader, fileName)
		assert.NoError(t, err)
	})

	t.Run("SUCCESS_delete", func(t *testing.T) {
		err := minioRepo.DeleteFile(ctx, fileName)
		assert.NoError(t, err)
	})
}
