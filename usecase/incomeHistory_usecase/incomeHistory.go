package incomeHistory_usecase

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

type IncomeHistoryUsecaseImpl struct {
	incomeTypeRepo    repository.IncomeTypeRepository
	incomeHistoryRepo repository.IncomeHistoryRepository
	paymentRepo       repository.PaymentRepository
	balanceUsecase    usecase.BalanceUsecase
}

func NewIncomeHistoryUsecaseImpl(
	incomeTypeRepo repository.IncomeTypeRepository,
	incomeHistoryRepo repository.IncomeHistoryRepository,
	paymentRepo repository.PaymentRepository,
	balanceUsecase usecase.BalanceUsecase,
) usecase.IncomeHistoryUsecase {
	return &IncomeHistoryUsecaseImpl{
		incomeTypeRepo:    incomeTypeRepo,
		incomeHistoryRepo: incomeHistoryRepo,
		paymentRepo:       paymentRepo,
		balanceUsecase:    balanceUsecase,
	}
}
