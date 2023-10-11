package incomeType_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (i *IncomeTypeUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*usecase.ResponseIncomeType, error) {
	incomeType, err := i.incomeTypeRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.IncomeTypeIsNotExist
		}
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

func (i *IncomeTypeUsecaseImpl) GetAllByProfileID(ctx context.Context, req *usecase.RequestGetAllByProfileIDWithISD) (*[]usecase.ResponseIncomeType, string, error) {
	order, operation := usecase.GetOrder(req.Order)
	incomeTypes, err := i.incomeTypeRepo.GetAllByProfileID(ctx, &repository.GetAllIncomeTypeWithISD{
		ProfileID: req.ProfileID,
		InfiniteScrollData: repository.InfiniteScrollData{
			ID:        req.ID,
			Order:     order,
			Operation: operation,
		},
	})
	if err != nil {
		return nil, "", err
	}

	if len(*incomeTypes) < 1 {
		return nil, "", nil
	}

	var resps []usecase.ResponseIncomeType
	var resp *usecase.ResponseIncomeType

	for _, incomeType := range *incomeTypes {
		resp = &usecase.ResponseIncomeType{
			ID:          incomeType.ID,
			ProfileID:   incomeType.ProfileID,
			Name:        incomeType.Name,
			Description: repository.GetNullString(incomeType.Description),
			Icon:        incomeType.Icon,
		}

		resps = append(resps, *resp)
	}

	cursor := (*incomeTypes)[len(*incomeTypes)-1].ID
	return &resps, cursor, nil
}
