package _usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type SpendingTypeUsecaseImpl struct {
	spendingTypeRepo domain.SpendingTypeRepository
}

func NewSpendingTypeUsecaseImpl(
	spendingTypeRepo domain.SpendingTypeRepository,
) domain.SpendingTypeUsecase {
	return &SpendingTypeUsecaseImpl{
		spendingTypeRepo: spendingTypeRepo,
	}
}

func (s *SpendingTypeUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateSpendingType) (*domain.ResponseSpendingType, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingTypeUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateSpendingType) (*domain.ResponseSpendingType, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingTypeUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingTypeUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.ResponseSpendingType, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingTypeUsecaseImpl) GetAllByProfileID(ctx context.Context, profileID string) (*[]domain.ResponseSpendingType, error) {
	// TODO implement me
	panic("implement me")
}
