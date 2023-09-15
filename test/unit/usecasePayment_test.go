package unit

import (
	"bytes"
	"database/sql"
	"errors"
	"mime/multipart"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain/mocks"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/_usecase"
)

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
func TestUsecaseCreatePayment(t *testing.T) {
	paymentRepo := &mocks.FakePaymentRepository{}
	minioRepo := &mocks.FakeMinioRepo{}

	paymentUsecase := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)

	t.Run("SUCCESS", func(t *testing.T) {
		ctx := context.TODO()
		fileHeader := newFileHeader()

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetByName(ctx, "bca")
		paymentRepo.GetByNameReturns(nil, sql.ErrNoRows)

		minioRepo.GenerateFileName(filepath.Ext(fileHeader.Filename), "payment-images/public/")
		minioRepo.GenerateFileNameReturns("/files/payment-images/public/12345678.png")

		paymentRepo.StartTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		}, func() error {
			return nil
		})
		paymentRepo.StartTxReturns(nil)

		payment, err := paymentUsecase.Create(ctx, &domain.RequestCreatePayment{
			Name:        "bca",
			Description: "bcatransfer",
			Image:       newFileHeader(),
		})
		assert.NoError(t, err)
		assert.NotNil(t, payment)
	})

	t.Run("ERROR_startx", func(t *testing.T) {
		ctx := context.TODO()
		fileHeader := newFileHeader()

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetByName(ctx, "bca")
		paymentRepo.GetByNameReturns(&domain.Payment{}, nil)

		minioRepo.GenerateFileName(filepath.Ext(fileHeader.Filename), "payment-images/public/")
		minioRepo.GenerateFileNameReturns("/files/payment-images/public/12345678.png")

		paymentRepo.StartTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		}, func() error {
			return nil
		})
		paymentRepo.StartTxReturns(errors.New("errors new"))

		payment, err := paymentUsecase.Create(ctx, &domain.RequestCreatePayment{
			Name:        "bca",
			Description: "bcatransfer",
			Image:       newFileHeader(),
		})
		assert.Error(t, err)
		assert.Nil(t, payment)
		var errHTTP *domain.ErrHTTP
		assert.Equal(t, true, errors.As(err, &errHTTP))
	})

	t.Run("ERROR_GetPaymentByName", func(t *testing.T) {
		ctx := context.TODO()

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetByName(ctx, "nil")
		paymentRepo.GetByNameReturns(&domain.Payment{}, nil)

		payment, err := paymentUsecase.Create(ctx, &domain.RequestCreatePayment{
			Name:        "bca",
			Description: "bcatransfer",
			Image:       newFileHeader(),
		})
		assert.Error(t, err)
		assert.Nil(t, payment)
		var errHTTP *domain.ErrHTTP
		assert.Equal(t, true, errors.As(err, &errHTTP))
		assert.Equal(t, 409, errHTTP.Code)
	})
}

func TestUsecaseUpdatePayment(t *testing.T) {
	paymentRepo := &mocks.FakePaymentRepository{}
	minioRepo := &mocks.FakeMinioRepo{}

	paymentUsecase := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)

	t.Run("SUCCESS_sama-name-req", func(t *testing.T) {
		ctx := context.TODO()
		fileHeader := newFileHeader()

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetByID(ctx, "123")
		paymentRepo.GetByIDReturns(&domain.Payment{Name: "sama"}, nil)

		paymentRepo.GetByName(ctx, "bca")
		paymentRepo.GetByNameReturns(&domain.Payment{Name: "beda"}, nil)

		minioRepo.GenerateFileName(filepath.Ext(fileHeader.Filename), "payment-images/public/")
		minioRepo.GenerateFileNameReturns("/files/payment-images/public/12345678.png")

		paymentRepo.StartTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		}, func() error {
			return nil
		})
		paymentRepo.StartTxReturns(nil)

		payment, err := paymentUsecase.Update(ctx, &domain.RequestUpdatePayment{
			ID:          "123",
			Name:        "sama",
			Description: "bcatransfer",
			Image:       newFileHeader(),
		})
		assert.NoError(t, err)
		assert.NotNil(t, payment)
	})

	t.Run("SUCCESS_beda-name-req", func(t *testing.T) {
		ctx := context.TODO()
		fileHeader := newFileHeader()

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetByID(ctx, "123")
		paymentRepo.GetByIDReturns(&domain.Payment{Name: "sama2", Image: "ada"}, nil)

		paymentRepo.GetByName(ctx, "bca")
		paymentRepo.GetByNameReturns(nil, sql.ErrNoRows)

		minioRepo.GenerateFileName(filepath.Ext(fileHeader.Filename), "payment-images/public/")
		minioRepo.GenerateFileNameReturns("/files/payment-images/public/12345678.png")

		paymentRepo.StartTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		}, func() error {
			return nil
		})
		paymentRepo.StartTxReturns(nil)

		payment, err := paymentUsecase.Update(ctx, &domain.RequestUpdatePayment{
			ID:          "123",
			Name:        "sama",
			Description: "bcatransfer",
			Image:       newFileHeader(),
		})
		assert.NoError(t, err)
		assert.NotNil(t, payment)
	})

	t.Run("SUCCESS_image-nil", func(t *testing.T) {
		ctx := context.TODO()
		fileHeader := newFileHeader()

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetByID(ctx, "123")
		paymentRepo.GetByIDReturns(&domain.Payment{Name: "sama2", Image: "ada"}, nil)

		paymentRepo.GetByName(ctx, "bca")
		paymentRepo.GetByNameReturns(nil, sql.ErrNoRows)

		minioRepo.GenerateFileName(filepath.Ext(fileHeader.Filename), "payment-images/public/")
		minioRepo.GenerateFileNameReturns("/files/payment-images/public/12345678.png")

		paymentRepo.StartTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		}, func() error {
			return nil
		})
		paymentRepo.StartTxReturns(nil)

		payment, err := paymentUsecase.Update(ctx, &domain.RequestUpdatePayment{
			ID:          "123",
			Name:        "sama",
			Description: "bcatransfer",
			Image:       nil,
		})
		assert.NoError(t, err)
		assert.NotNil(t, payment)
	})

	t.Run("ERROR_invalid-id", func(t *testing.T) {
		ctx := context.TODO()

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetByID(ctx, "nil")
		paymentRepo.GetByIDReturns(nil, sql.ErrNoRows)

		payment, err := paymentUsecase.Update(ctx, &domain.RequestUpdatePayment{
			ID:          "123",
			Name:        "bca",
			Description: "bcatransfer",
			Image:       newFileHeader(),
		})
		assert.Error(t, err)
		assert.Nil(t, payment)
		var errHTTP *domain.ErrHTTP
		assert.Equal(t, true, errors.As(err, &errHTTP))
		assert.Equal(t, 404, errHTTP.Code)
	})

	t.Run("ERROR_sama-name-req", func(t *testing.T) {
		ctx := context.TODO()

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetByID(ctx, "123")
		paymentRepo.GetByIDReturns(&domain.Payment{Name: "sama"}, nil)

		paymentRepo.GetByName(ctx, "bca")
		paymentRepo.GetByNameReturns(&domain.Payment{Name: "sama"}, nil)

		payment, err := paymentUsecase.Update(ctx, &domain.RequestUpdatePayment{
			ID:          "123",
			Name:        "bca",
			Description: "bcatransfer",
			Image:       newFileHeader(),
		})
		assert.Error(t, err)
		assert.Nil(t, payment)
		var errHTTP *domain.ErrHTTP
		assert.Equal(t, true, errors.As(err, &errHTTP))
		assert.Equal(t, 409, errHTTP.Code)
	})
}

func TestUsecaseGetAllPayment(t *testing.T) {
	paymentRepo := &mocks.FakePaymentRepository{}
	minioRepo := &mocks.FakeMinioRepo{}

	paymentUsecase := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)

	t.Run("SUCCESS", func(t *testing.T) {
		ctx := context.TODO()
		data := &[]domain.Payment{
			{
				ID: "id1",
			},
			{
				ID: "id2",
			},
		}

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetAll(ctx)
		paymentRepo.GetAllReturns(data, nil)

		payments, err := paymentUsecase.GetAll(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, payments)
		assert.Equal(t, 2, len(*payments))
	})

	t.Run("ERROR", func(t *testing.T) {
		ctx := context.TODO()

		paymentRepo.OpenConn(ctx)
		paymentRepo.OpenConnReturns(nil)
		defer paymentRepo.CloseConn()

		paymentRepo.GetAll(ctx)
		paymentRepo.GetAllReturns(nil, errors.New("error db"))

		payments, err := paymentUsecase.GetAll(ctx)
		assert.Error(t, err)
		assert.Nil(t, payments)
		var errHTTP *domain.ErrHTTP
		assert.Equal(t, true, errors.As(err, &errHTTP))
		assert.Equal(t, 500, errHTTP.Code)
	})
}
