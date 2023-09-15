package _repository

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type BalanceRepositoryImpl struct {
	domain.UnitOfWorkRepository
}

func NewBalanceRepositoryImpl(
	uow domain.UnitOfWorkRepository,
) domain.BalanceRepository {
	return &BalanceRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}

func (b *BalanceRepositoryImpl) Create(ctx context.Context, balance *domain.Balance) error {
	// TODO implement me
	panic("implement me")
}

func (b *BalanceRepositoryImpl) UpdateByProfileID(ctx context.Context, balance *domain.Balance) error {
	// TODO implement me
	panic("implement me")
}

func (b *BalanceRepositoryImpl) GetByProfileID(ctx context.Context, profileID string) (*domain.Balance, error) {
	// TODO implement me
	panic("implement me")
}
