package config

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

func NewMinioConn() (*minio.Client, error) {
	minioConn, err := minio.New(MinIoEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinIoID, MinIoSecretKey, ""),
		Secure: MinIoSSL,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed open connection minio")
	}

	return minioConn, err
}
