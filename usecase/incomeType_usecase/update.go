package incomeType_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (i *IncomeTypeUsecaseImpl) Update(ctx context.Context, req *usecase.RequestUpdateIncomeType) (*usecase.ResponseIncomeType, error) {
	incomeType, err := i.incomeTypeRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.IncomeTypeIsNotExist
		}
		return nil, err
	}

	if incomeType.Name != req.Name {
		exist, err := i.incomeTypeRepo.CheckByNameAndProfileID(ctx, req.ProfileID, req.Name)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, usecase.NameIncomeTypeIsExist
		}
	}

	incomeTypeConv := req.ToModel()

	err = i.incomeTypeRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		err = i.incomeTypeRepo.Update(ctx, incomeTypeConv)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	resp := &usecase.ResponseIncomeType{
		ID:          incomeTypeConv.ID,
		ProfileID:   incomeTypeConv.ProfileID,
		Name:        incomeTypeConv.Name,
		Description: repository.GetNullString(incomeTypeConv.Description),
		Icon:        incomeTypeConv.Icon,
	}

	return resp, nil
}
