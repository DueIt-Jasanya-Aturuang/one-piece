package services

import (
	"context"

	dbimpl "github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/infrastructures/db/dbImpl"
	converterEntity "github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/helpers/converter-entity"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/helpers/minio"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/dto"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/entities"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/repositories"
	"github.com/go-playground/validator/v10"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, req *dto.PaymentCreateRequest) (*entities.Payment, error)
	UpdatePayment(ctx context.Context, req *dto.PaymentUpdateRequest, id string) (*dto.PaymentResponse, error)
	GetPaymentByName(ctx context.Context, name string) (*dto.PaymentResponse, error)
	DeletePayment(ctx context.Context, id string) (bool, error)
}

type PaymentServiceImpl struct {
	PaymentRepository repositories.Repository
	DbImpl            *dbimpl.DbImpl
	Convert           *converterEntity.ConvertImpl
	Validation        *validator.Validate
	Minio             *minio.MinioImpl
}

func NewPaymentServiceImpl(
	paymentRepository repositories.Repository,
	dbImpl *dbimpl.DbImpl,
	convert *converterEntity.ConvertImpl,
	validation *validator.Validate,
	minio *minio.MinioImpl,
) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: paymentRepository,
		DbImpl:            dbImpl,
		Convert:           convert,
		Validation:        validation,
		Minio:             minio,
	}
}
