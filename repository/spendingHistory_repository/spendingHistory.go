package spendingHistory_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
)

type SpendingHistoryRepositoryImpl struct {
	repository.UnitOfWorkRepository
}

func NewSpendingHistoryRepositoryImpl(
	uow repository.UnitOfWorkRepository,
) repository.SpendingHistoryRepository {
	return &SpendingHistoryRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}
