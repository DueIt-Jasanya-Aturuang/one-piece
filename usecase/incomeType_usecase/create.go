package incomeType_usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (i *IncomeTypeUsecaseImpl) Create(ctx context.Context, req *usecase.RequestCreateIncomeType) (*usecase.ResponseIncomeType, error) {
	exist, err := i.incomeTypeRepo.CheckByNameAndProfileID(ctx, req.ProfileID, req.Name)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, usecase.NameIncomeTypeIsExist
	}

	incomeType := req.ToModel()
	err = i.incomeTypeRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		err = i.incomeTypeRepo.Create(ctx, incomeType)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	resp := &usecase.ResponseIncomeType{
		ID:          incomeType.ID,
		ProfileID:   incomeType.ProfileID,
		Name:        incomeType.Name,
		Description: repository.GetNullString(incomeType.Description),
		Icon:        incomeType.Icon,
	}
	return resp, nil
}
