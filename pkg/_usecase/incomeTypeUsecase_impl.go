package _usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/helper"
)

type IncomeTypeUsecaseImpl struct {
	incomeTypeRepo domain.IncomeTypeRepository
}

func NewIncomeTypeUsecaseImpl(
	incomeTypeRepo domain.IncomeTypeRepository,
) domain.IncomeTypeUsecase {
	return &IncomeTypeUsecaseImpl{
		incomeTypeRepo: incomeTypeRepo,
	}
}

func (i *IncomeTypeUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateIncomeType) (*domain.ResponseIncomeType, error) {
	if err := i.incomeTypeRepo.OpenConn(ctx); err != nil {
		return nil, err
	}
	defer i.incomeTypeRepo.CloseConn()

	exist, err := i.incomeTypeRepo.CheckByNameAndProfileID(ctx, req.ProfileID, req.Name)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, NameIncomeTypeIsExist
	}

	incomeType := converter.RequestCreateIncomeTypeToModel(req)
	err = i.incomeTypeRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		err = i.incomeTypeRepo.Create(ctx, incomeType)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	resp := converter.IncomeTypeModelToResp(incomeType)
	return resp, nil
}

func (i *IncomeTypeUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateIncomeHistory) (*domain.ResponseIncomeType, error) {
	// TODO implement me
	panic("implement me")
}

func (i *IncomeTypeUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	// TODO implement me
	panic("implement me")
}

func (i *IncomeTypeUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.ResponseIncomeType, error) {
	// TODO implement me
	panic("implement me")
}

func (i *IncomeTypeUsecaseImpl) GetAllByProfileID(ctx context.Context, profileID string) (*[]domain.ResponseIncomeType, error) {
	// TODO implement me
	panic("implement me")
}
