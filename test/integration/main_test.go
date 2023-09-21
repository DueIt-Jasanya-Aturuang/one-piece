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

	"github.com/DueIt-Jasanya-Aturuang/one-piece/infra"
	_repository2 "github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/_usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/test/integration/setup"
)

var DB *sql.DB
var minioClient *minio.Client
var Uow = _repository2.NewUnitOfWorkRepositoryImpl(DB)
var PaymentRepo = _repository2.NewPaymentRepositoryImpl(Uow)
var SpendingTypeRepo = _repository2.NewSpendingTypeRepositoryImpl(Uow)
var SpendingHistoryRepo = _repository2.NewSpendingHistoryRepositoryImpl(Uow)
var SpendingTypeUsecase = _usecase.NewSpendingTypeUsecaseImpl(SpendingTypeRepo)

func TestMain(m *testing.M) {
	infra.LogInit()
	dockerpool := setup.SetupDocker()
	var resources []*dockertest.Resource

	pgResource, dbPg, _ := setup.Postgres(dockerpool)
	resources = append(resources, pgResource)
	DB = dbPg
	Uow = _repository2.NewUnitOfWorkRepositoryImpl(DB)
	PaymentRepo = _repository2.NewPaymentRepositoryImpl(Uow)
	SpendingTypeRepo = _repository2.NewSpendingTypeRepositoryImpl(Uow)
	SpendingHistoryRepo = _repository2.NewSpendingHistoryRepositoryImpl(Uow)
	SpendingTypeUsecase = _usecase.NewSpendingTypeUsecaseImpl(SpendingTypeRepo)

	if DB == nil {
		panic("db nil")
	}

	setup.Migrate(DB)

	minioResourece, endpoint := setup.MinioStart(dockerpool)
	resources = append(resources, minioResourece)
	infra.MinIoEndpoint, infra.MinIoID, infra.MinIoSecretKey, infra.MinIoSSL = endpoint, "MYACCESSKEY", "MYSECRETKEY", false
	minioConn := infra.NewMinioConn()
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
		t.Run("Create", CreateSpendingType)
		t.Run("Update", UpdateSpendingType)
		t.Run("Delete", DeleteSpendingType)
		t.Run("CheckData", CheckDataSpendingType)
		t.Run("CheckByTitleAndProfileID", CheckByTitleAndProfileIDSpendingType)
		t.Run("GetByID", GetByIDSpendingType)
		t.Run("GetByID_ERROR-deleted_at-null", GetByIDSpendingTypeERRORDeletedAtNull)
		t.Run("GetByID_ERROR-invalid-id", GetByIDSpendingTypeERRORInvalidID)
		t.Run("GetByIDAndProfileID", GetByIDAndProfileIDSpendingType)
		t.Run("GetByIDAndProfileID_ERROR-deleted_at-null", GetByIDAndProfileIDSpendingTypeERRORDeletedAtNull)
		t.Run("GetByIDAndProfileID_ERROR-invalid-id", GetByIDAndProfileIDSpendingTypeERRORInvalidID)
		t.Run("GetByIDAndProfileID_ERROR-invalid-profile_id", GetByIDAndProfileIDSpendingTypeERRORInvalidProfileID)
		t.Run("GetAllByProfileID", GetAllByProfileIDSpendingType)
		t.Run("GetDefault", GetDefaultSpendingType)
	})

	t.Run("SPENDINGHISTORY_REPO", func(t *testing.T) {
		t.Run("Create", CreateSpendingHistory)
		t.Run("Update", UpdateSpendingHistory)
		t.Run("Delete", DeleteSpendingHistory)
		t.Run("GetAllByTimeAndProfileID", GetAllByTimeAndProfileIDSpendingHistory)
		t.Run("GetByIDAndProfileID", GetByIDAndProfileIDSpendingHistory)
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

	t.Run("SPENDINGTYPE_USECASE", func(t *testing.T) {
		t.Run("Create", CreateSpendingTypeUsecase)
		t.Run("Create_ERRORNameAlready", CreateSpendingTypeUsecaseERRORNameAlready)
		t.Run("Update", UpdateSpendingTypeUsecase)
		t.Run("Update_ERRORNameAlready", UpdateSpendingTypeUsecaseERRORNameAlready)
		t.Run("Delete", DeleteSpendingTypeUsecase)
		t.Run("GetByIDAndProfileID", GetByIDAndProfileIDSpendingTypeUsecase)
		t.Run("GetByIDAndProfileID_ERRORNoRow", GetByIDAndProfileIDSpendingTypeUsecaseERRORNoRow)
		t.Run("GetAllByProfileID", GetAllByProfileIDSpendingTypeUsecase)
		t.Run("GetAllByProfileID_WithCreateDefaultType", GetAllByProfileIDSpendingTypeUsecaseWithCreateDefaultType)
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
