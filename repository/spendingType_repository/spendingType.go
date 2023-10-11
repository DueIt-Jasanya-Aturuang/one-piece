package spendingType_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
)

type SpendingTypeRepositoryImpl struct {
	repository.UnitOfWorkRepository
}

func NewSpendingTypeRepositoryImpl(
	uow repository.UnitOfWorkRepository,
) repository.SpendingTypeRepository {
	return &SpendingTypeRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}
