package spendingType_usecase

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

type SpendingTypeUsecaseImpl struct {
	spendingTypeRepo    repository.SpendingTypeRepository
	spendingHistoryRepo repository.SpendingHistoryRepository
}

func NewSpendingTypeUsecaseImpl(
	spendingTypeRepo repository.SpendingTypeRepository,
	spendingHistoryRepo repository.SpendingHistoryRepository,
) usecase.SpendingTypeUsecase {
	return &SpendingTypeUsecaseImpl{
		spendingTypeRepo:    spendingTypeRepo,
		spendingHistoryRepo: spendingHistoryRepo,
	}
}
