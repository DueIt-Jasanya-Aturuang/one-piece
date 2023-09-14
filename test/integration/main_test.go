package integration

import (
	"bytes"
	"database/sql"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/ory/dockertest/v3"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra/config"
	_repository2 "github.com/DueIt-Jasanya-Aturuang/one-piece/internal/_repository"
)

var DB *sql.DB
var minioClient *minio.Client
var Uow = _repository2.NewUnitOfWorkRepositoryImpl(DB)
var PaymentRepo = _repository2.NewPaymentRepositoryImpl(Uow)
var SpendingTypeRepo = _repository2.NewSpendingTypeRepositoryImpl(Uow)

func TestMain(m *testing.M) {
	dockerpool := SetupDocker()
	var resources []*dockertest.Resource

	pgResource, dbPg, _ := Postgres(dockerpool)
	resources = append(resources, pgResource)
	DB = dbPg
	Uow = _repository2.NewUnitOfWorkRepositoryImpl(DB)
	PaymentRepo = _repository2.NewPaymentRepositoryImpl(Uow)
	SpendingTypeRepo = _repository2.NewSpendingTypeRepositoryImpl(Uow)
	if DB == nil {
		panic("db nil")
	}

	Migrate(DB)

	minioResourece, endpoint := minioStart(dockerpool)
	resources = append(resources, minioResourece)
	config.MinIoEndpoint, config.MinIoID, config.MinIoSecretKey, config.MinIoSSL = endpoint, "MYACCESSKEY", "MYSECRETKEY", false
	minioConn := config.NewMinioConn()
	minioClient = minioConn

	code := m.Run()

	for _, resource := range resources {
		if err := dockerpool.Purge(resource); err != nil {
			log.Fatal().Msgf("failed purge resource")
		}
	}
	os.Exit(code)
}

func TestInit(t *testing.T) {
	t.Run("PAYMENT_REPO", func(t *testing.T) {
		t.Run("Create", CreatePayment)
		t.Run("GetByID", GetPaymentById)
		t.Run("GetPaymentById_ERROR", GetPaymentByIdERROR)
		t.Run("Update", UpdatePayment)
		t.Run("GetByName", GetPaymentByName)
		t.Run("GetPaymentByName_ERROR", GetPaymentByNameERROR)
	})

	t.Run("SPENDINGTYPE_REPO", func(t *testing.T) {
		t.Run("Create", CreateSpendingTYpe)
	})

	t.Run("MINIO_REPO", func(t *testing.T) {
		t.Run("createBucket", createBucket)
		t.Run("MinioRepo", minioRepo)
	})

	t.Run("PAYMENT_USECASE", func(t *testing.T) {
		t.Run("Create", UsecaseCreatePayment)
		t.Run("CreatePayment409ERROR", UsecaseCreatePayment409ERROR)
		t.Run("Update", UsecaseUpdatePayment)
		t.Run("UpdatePaymentERROR", UsecaseUpdatePaymentERROR)
		t.Run("GetAll", UsecaseGetAllPayment)
	})
}

func newFileHeader() *multipart.FileHeader {
	fileContent := []byte("Contoh isi file")
	fileHeader := &multipart.FileHeader{
		Filename: "example.png",
		Size:     int64(len(fileContent)),
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		log.Fatal().Err(err).Msgf("error")
	}
	part.Write(fileContent)

	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		log.Fatal().Err(err).Msgf("error")
	}
	defer func() {
		_ = file.Close()
	}()

	return fileHeader
}
