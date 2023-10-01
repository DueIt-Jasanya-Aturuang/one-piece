package usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type IncomeHistoryUsecaseImpl struct {
	incomeTypeRepo    domain.IncomeTypeRepository
	incomeHistoryRepo domain.IncomeHistoryRepository
	paymentRepo       domain.PaymentRepository
	balanceRepo       domain.BalanceRepository
}

func NewIncomeHistoryUsecaseImpl(
	incomeTypeRepo domain.IncomeTypeRepository,
	incomeHistoryRepo domain.IncomeHistoryRepository,
	paymentRepo domain.PaymentRepository,
	balanceRepo domain.BalanceRepository,
) domain.IncomeHistoryUsecase {
	return &IncomeHistoryUsecaseImpl{
		incomeTypeRepo:    incomeTypeRepo,
		incomeHistoryRepo: incomeHistoryRepo,
		paymentRepo:       paymentRepo,
		balanceRepo:       balanceRepo,
	}
}

func (i *IncomeHistoryUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateIncomeHistory) (*domain.ResponseIncomeHistory, error) {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	// TODO implement me
	panic("implement me")
}

func (i *IncomeHistoryUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateIncomeHistory) (*domain.ResponseIncomeHistory, error) {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	// TODO implement me
	panic("implement me")
}

func (i *IncomeHistoryUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return err
	}
	// TODO implement me
	panic("implement me")
}

func (i *IncomeHistoryUsecaseImpl) GetAllByTimeAndProfileID(ctx context.Context, req *domain.GetFilteredDataIncomeHistory) (*[]domain.ResponseIncomeHistory, error) {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	// TODO implement me
	panic("implement me")
}

func (i *IncomeHistoryUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.ResponseIncomeHistory, error) {
	err := i.incomeHistoryRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	// TODO implement me
	panic("implement me")
}
