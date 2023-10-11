package payment_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
)

type PaymentRepositoryImpl struct {
	repository.UnitOfWorkRepository
}

func NewPaymentRepositoryImpl(
	uow repository.UnitOfWorkRepository,
) repository.PaymentRepository {
	return &PaymentRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}
