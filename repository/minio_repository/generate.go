package minio_repository

import (
	"fmt"
	"time"
)

func (m *MinioRepositoryImpl) GenerateFileName(fileExt string, path string) string {
	nameFile := fmt.Sprintf("%s%d%s", path, time.Now().UnixNano(), fileExt)
	return nameFile
}
