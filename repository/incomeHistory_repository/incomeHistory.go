package incomeHistory_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
)

type IncomeHistoryRepositoryImpl struct {
	repository.UnitOfWorkRepository
}

func NewIncomeHistoryRepositoryImpl(uow repository.UnitOfWorkRepository) repository.IncomeHistoryRepository {
	return &IncomeHistoryRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}
