package _usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type SpendingHistoryUsecaseImpl struct {
	spendingHistoryRepo domain.SpendingHistoryRepository
	spendingTypeRepo    domain.SpendingTypeRepository
}

func NewSpendingHistoryUsecaseImpl(
	spendingHistoryRepo domain.SpendingHistoryRepository,
	spendingTypeRepo domain.SpendingTypeRepository,
) domain.SpendingHistoryUsecase {
	return &SpendingHistoryUsecaseImpl{
		spendingHistoryRepo: spendingHistoryRepo,
		spendingTypeRepo:    spendingTypeRepo,
	}
}

func (s *SpendingHistoryUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateSpendingHistory) (*domain.ResponseSpendingHistory, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateSpendingHistory) (*domain.ResponseSpendingHistory, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryUsecaseImpl) GetAllByProfileID(ctx context.Context, profileID string) (*[]domain.ResponseSpendingHistory, error) {
	// TODO implement me
	panic("implement me")
}

func (s *SpendingHistoryUsecaseImpl) GetByIDAndProfileID(
	ctx context.Context, id string, profileID string,
) (*domain.ResponseSpendingHistory, error) {
	// TODO implement me
	panic("implement me")
}
