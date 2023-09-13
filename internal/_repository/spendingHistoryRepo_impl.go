package _repository

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type SpendingHistoryRepositoryImpl struct {
	domain.UnitOfWorkRepository
}

func NewSpendingHistoryRepositoryImpl(
	uow domain.UnitOfWorkRepository,
) domain.SpendingHistoryRepository {
	return &SpendingHistoryRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}

func (s *SpendingHistoryRepositoryImpl) Create(ctx context.Context, spendingHistory *domain.SpendingHistory) error {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryRepositoryImpl) Update(ctx context.Context, spendingHistory *domain.SpendingHistory) error {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryRepositoryImpl) Delete(ctx context.Context, id string, profileID string) error {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryRepositoryImpl) GetAllByProfileID(ctx context.Context, profileID string) (*[]domain.SpendingHistory, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.SpendingHistory, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.SpendingHistory, error) {
	// TODO implement me
	panic("implement me")
}
