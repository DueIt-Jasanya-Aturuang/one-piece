package incomeType_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
)

type IncomeTypeRepositoryImpl struct {
	repository.UnitOfWorkRepository
}

func NewIncomeTypeRepositoryImpl(uow repository.UnitOfWorkRepository) repository.IncomeTypeRepository {
	return &IncomeTypeRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}
