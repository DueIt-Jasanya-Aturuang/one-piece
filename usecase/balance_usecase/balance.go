package balance_usecase

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

type BalanceUsecaseImpl struct {
	balanceRepo repository.BalanceRepository
}

func NewBalanceUsecaseImpl(balanceRepo repository.BalanceRepository) usecase.BalanceUsecase {
	return &BalanceUsecaseImpl{
		balanceRepo: balanceRepo,
	}
}
