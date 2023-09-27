package _usecase

import (
	"context"
	"database/sql"
	"errors"

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

func (i *IncomeTypeUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateIncomeType) (*domain.ResponseIncomeType, error) {
	if err := i.incomeTypeRepo.OpenConn(ctx); err != nil {
		return nil, err
	}
	defer i.incomeTypeRepo.CloseConn()

	incomeType, err := i.incomeTypeRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, IncomeTypeIsNotExist
		}
		return nil, err
	}

	if incomeType.Name != req.Name {
		exist, err := i.incomeTypeRepo.CheckByNameAndProfileID(ctx, req.ProfileID, req.Name)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, NameIncomeTypeIsExist
		}
	}

	incomeTypeConv := converter.RequestUpdateIncomeTypeToModel(req)

	err = i.incomeTypeRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		err = i.incomeTypeRepo.Update(ctx, incomeTypeConv)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	resp := converter.IncomeTypeModelToResp(incomeTypeConv)

	return resp, nil
}

func (i *IncomeTypeUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	if err := i.incomeTypeRepo.OpenConn(ctx); err != nil {
		return err
	}
	defer i.incomeTypeRepo.CloseConn()

	err := i.incomeTypeRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		err := i.incomeTypeRepo.Delete(ctx, id, profileID)
		if err != nil {
			return nil
		}

		return nil
	})

	return err
}

func (i *IncomeTypeUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.ResponseIncomeType, error) {
	if err := i.incomeTypeRepo.OpenConn(ctx); err != nil {
		return nil, err
	}
	defer i.incomeTypeRepo.CloseConn()

	incomeType, err := i.incomeTypeRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, IncomeTypeIsNotExist
		}
		return nil, err
	}

	resp := converter.IncomeTypeModelToResp(incomeType)

	return resp, nil
}

func (i *IncomeTypeUsecaseImpl) GetAllByProfileID(ctx context.Context, profileID string) (*[]domain.ResponseIncomeType, error) {
	if err := i.incomeTypeRepo.OpenConn(ctx); err != nil {
		return nil, err
	}
	defer i.incomeTypeRepo.CloseConn()

	incomeTypes, err := i.incomeTypeRepo.GetAllByProfileID(ctx, profileID)
	if err != nil {
		return nil, err
	}

	var resps []domain.ResponseIncomeType

	for _, incomeType := range *incomeTypes {
		resp := converter.IncomeTypeModelToResp(&incomeType)
		resps = append(resps, *resp)
	}

	return &resps, nil
}
