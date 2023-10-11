package spendingType_usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (s *SpendingTypeUsecaseImpl) Create(ctx context.Context, req *usecase.RequestCreateSpendingType) (*usecase.ResponseSpendingType, error) {
	exist, err := s.spendingTypeRepo.CheckByTitleAndProfileID(ctx, req.ProfileID, req.Title)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, usecase.TitleSpendingTypeExist
	}

	spendingType := req.ToModel()
	err = s.spendingTypeRepo.StartTx(ctx, repository.LevelReadCommitted(), func() error {
		err = s.spendingTypeRepo.Create(ctx, spendingType)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	formatMaximumLimit := usecase.FormatRupiah(spendingType.MaximumLimit)
	resp := &usecase.ResponseSpendingType{
		ID:                 spendingType.ID,
		ProfileID:          spendingType.ProfileID,
		Title:              spendingType.Title,
		MaximumLimit:       spendingType.MaximumLimit,
		FormatMaximumLimit: formatMaximumLimit,
		Icon:               spendingType.Icon,
	}

	return resp, nil
}
