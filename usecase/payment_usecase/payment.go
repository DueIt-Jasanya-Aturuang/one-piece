package payment_usecase

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

type PaymentUsecaseImpl struct {
	paymentRepo repository.PaymentRepository
	minioRepo   repository.MinioRepository
}

func NewPaymentUsecaseImpl(
	paymentRepo repository.PaymentRepository,
	minioRepo repository.MinioRepository,
) usecase.PaymentUsecase {
	return &PaymentUsecaseImpl{
		paymentRepo: paymentRepo,
		minioRepo:   minioRepo,
	}
}
