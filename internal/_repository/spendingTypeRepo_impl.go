package _repository

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type SpendingTypeRepositoryImpl struct {
	domain.UnitOfWorkRepository
}

func NewSpendingTypeRepositoryImpl(
	uow domain.UnitOfWorkRepository,
) domain.SpendingTypeRepository {
	return &SpendingTypeRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}

func (s *SpendingTypeRepositoryImpl) Create(ctx context.Context, spendingType *domain.SpendingType) error {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingTypeRepositoryImpl) Update(ctx context.Context, spendingType *domain.SpendingType) error {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingTypeRepositoryImpl) Delete(ctx context.Context, id string, profileID string) error {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingTypeRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.SpendingType, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingTypeRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.SpendingType, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingTypeRepositoryImpl) GetAllByProfileID(ctx context.Context, profileID string) (*domain.SpendingType, error) {
	// TODO implement me
	panic("implement me")
}
