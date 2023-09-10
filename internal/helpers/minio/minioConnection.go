package minio

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

func Connection() (*minio.Client, error) {
	endpoint := config.Get().ThirdParty.Minio.Endpoint
	accessKeyID := config.Get().ThirdParty.Minio.AccessKey
	secretKey := config.Get().ThirdParty.Minio.SecretKey
	// bucket := config.Get().ThirdParty.Minio.Bucket
	ssl := config.Get().ThirdParty.Minio.SSL

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretKey, ""),
		Secure: ssl,
	})
	if err != nil {
		log.Err(err).Msg("cannot init minio")
	}

	return minioClient, err
}
