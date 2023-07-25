package services

import (
	dbimpl "github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/infrastructures/db/dbImpl"
	converterEntity "github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/helpers/converter-entity"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/repositories"
	"github.com/go-playground/validator/v10"
)

type PaymentService interface{}

type PaymentServiceImpl struct {
	PaymentRepository repositories.Repository
	DbImpl            *dbimpl.DbImpl
	Convert           *converterEntity.ConvertImpl
	Validation        *validator.Validate
}

func NewPaymentServiceImpl(
	paymentRepository repositories.Repository,
	dbImpl *dbimpl.DbImpl,
	convert *converterEntity.ConvertImpl,
	validation *validator.Validate,
) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: paymentRepository,
		DbImpl:            dbImpl,
		Convert:           convert,
		Validation:        validation,
	}
}
