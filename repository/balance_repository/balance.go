package balance_repository

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
)

type BalanceRepositoryImpl struct {
	repository.UnitOfWorkRepository
}

func NewBalanceRepositoryImpl(
	uow repository.UnitOfWorkRepository,
) repository.BalanceRepository {
	return &BalanceRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}
