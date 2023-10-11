package minio_repository

import (
	"github.com/minio/minio-go/v7"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
)

type MinioRepositoryImpl struct {
	c *minio.Client
}

func NewMinioImpl(c *minio.Client) repository.MinioRepository {
	return &MinioRepositoryImpl{
		c: c,
	}
}
