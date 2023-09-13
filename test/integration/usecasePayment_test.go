package integration

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/_repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/_usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/helper"
)

func UsecaseCreatePayment(t *testing.T) {
	uow := _repository.NewUnitOfWorkRepositoryImpl(DB)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)
	minioRepo := _repository.NewMinioImpl(minioClient)
	ctx := context.TODO()
	usecasePayment := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)

	t.Run("SUCCESS", func(t *testing.T) {
		req := &domain.RequestCreatePayment{
			Name:        "bcausecase",
			Description: "bca mandiri",
			Image:       newFileHeader(),
		}

		payment, err := usecasePayment.Create(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, *payment.Description, req.Description)
	})

	t.Run("SUCCESS_description-nil", func(t *testing.T) {
		req := &domain.RequestCreatePayment{
			Name:        "bcausecase2",
			Description: "",
			Image:       newFileHeader(),
		}

		payment, err := usecasePayment.Create(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, helper.GetNullString(sql.NullString{}), payment.Description)
	})

}

func UsecaseCreatePayment409ERROR(t *testing.T) {
	uow := _repository.NewUnitOfWorkRepositoryImpl(DB)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)
	minioRepo := _repository.NewMinioImpl(minioClient)
	ctx := context.TODO()
	usecasePayment := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)

	t.Run("ERROR", func(t *testing.T) {
		req := &domain.RequestCreatePayment{
			Name:        "bca",
			Description: "",
			Image:       newFileHeader(),
		}

		payment, err := usecasePayment.Create(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, payment)
		var errHTTP *domain.ErrHTTP
		assert.Equal(t, true, errors.As(err, &errHTTP))
		assert.Equal(t, 409, errHTTP.Code)
	})
}

func UsecaseUpdatePayment(t *testing.T) {
	uow := _repository.NewUnitOfWorkRepositoryImpl(DB)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)
	minioRepo := _repository.NewMinioImpl(minioClient)
	ctx := context.TODO()
	usecasePayment := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)

	t.Run("SUCCESS", func(t *testing.T) {
		req := &domain.RequestUpdatePayment{
			ID:          "payment1",
			Name:        "bca",
			Description: "bca mandiri banget",
			Image:       newFileHeader(),
		}

		payment, err := usecasePayment.Update(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, *payment.Description, req.Description)
	})

	t.Run("SUCCESS_image-nil", func(t *testing.T) {
		req := &domain.RequestUpdatePayment{
			ID:          "payment1",
			Name:        "bca",
			Description: "bca mandiri banget",
			Image:       nil,
		}

		payment, err := usecasePayment.Update(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, *payment.Description, req.Description)
	})

	t.Run("SUCCESS_beda-nama", func(t *testing.T) {
		req := &domain.RequestUpdatePayment{
			ID:          "payment1",
			Name:        "bcabeda",
			Description: "bca mandiri banget",
			Image:       nil,
		}

		payment, err := usecasePayment.Update(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, payment)
		assert.Equal(t, *payment.Description, req.Description)
	})

}

func UsecaseUpdatePaymentERROR(t *testing.T) {
	uow := _repository.NewUnitOfWorkRepositoryImpl(DB)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)
	minioRepo := _repository.NewMinioImpl(minioClient)
	ctx := context.TODO()
	usecasePayment := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)

	t.Run("ERROR_409", func(t *testing.T) {
		req := &domain.RequestUpdatePayment{
			ID:          "payment1",
			Name:        "bcausecase",
			Description: "",
			Image:       newFileHeader(),
		}

		payment, err := usecasePayment.Update(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, payment)
		var errHTTP *domain.ErrHTTP
		assert.Equal(t, true, errors.As(err, &errHTTP))
		assert.Equal(t, 409, errHTTP.Code)
	})

	t.Run("ERROR_404", func(t *testing.T) {
		req := &domain.RequestUpdatePayment{
			ID:          "1",
			Name:        "bcausecase",
			Description: "",
			Image:       newFileHeader(),
		}

		payment, err := usecasePayment.Update(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, payment)
		var errHTTP *domain.ErrHTTP
		assert.Equal(t, true, errors.As(err, &errHTTP))
		assert.Equal(t, 404, errHTTP.Code)
	})
}

func UsecaseGetAllPayment(t *testing.T) {
	uow := _repository.NewUnitOfWorkRepositoryImpl(DB)
	paymentRepo := _repository.NewPaymentRepositoryImpl(uow)
	minioRepo := _repository.NewMinioImpl(minioClient)
	ctx := context.TODO()
	usecasePayment := _usecase.NewPaymentUsecaseImpl(paymentRepo, minioRepo)

	payments, err := usecasePayment.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, payments)
	t.Log(payments)
}