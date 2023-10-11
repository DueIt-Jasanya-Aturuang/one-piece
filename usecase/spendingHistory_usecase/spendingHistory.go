package spendingHistory_usecase

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

type SpendingHistoryUsecaseImpl struct {
	spendingHistoryRepo repository.SpendingHistoryRepository
	spendingTypeRepo    repository.SpendingTypeRepository
	balanceUsecase      usecase.BalanceUsecase
	paymentRepo         repository.PaymentRepository
}

func NewSpendingHistoryUsecaseImpl(
	spendingHistoryRepo repository.SpendingHistoryRepository,
	spendingTypeRepo repository.SpendingTypeRepository,
	balanceUsecase usecase.BalanceUsecase,
	paymentRepo repository.PaymentRepository,
) usecase.SpendingHistoryUsecase {
	return &SpendingHistoryUsecaseImpl{
		spendingHistoryRepo: spendingHistoryRepo,
		spendingTypeRepo:    spendingTypeRepo,
		balanceUsecase:      balanceUsecase,
		paymentRepo:         paymentRepo,
	}
}
