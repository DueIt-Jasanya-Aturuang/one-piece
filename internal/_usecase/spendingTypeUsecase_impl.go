package _usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
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
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}

	spendingType := converter.SpendingTypeRequestCreateToModel(req)
	err = s.spendingTypeRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = s.spendingTypeRepo.Create(ctx, spendingType)
		return err
	})

	if err != nil {
		return nil, err
	}

	formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)
	resp := converter.SpendingTypeModelToResponse(spendingType, formatMaximumLimit)

	return resp, nil
}

func (s *SpendingTypeUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateSpendingType) (*domain.ResponseSpendingType, error) {
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}

	spendingType := converter.SpendingTypeRequestUpdateToModel(req)
	err = s.spendingTypeRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = s.spendingTypeRepo.Update(ctx, spendingType)
		return err
	})

	if err != nil {
		return nil, err
	}

	formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)
	resp := converter.SpendingTypeModelToResponse(spendingType, formatMaximumLimit)

	return resp, nil
}

func (s *SpendingTypeUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return err
	}

	err = s.spendingTypeRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = s.spendingTypeRepo.Delete(ctx, id, profileID)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *SpendingTypeUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.ResponseSpendingType, error) {
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}

	spendingType, err := s.spendingTypeRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrHTTPString("data not found", 404)
		}
		return nil, err
	}

	formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)
	resp := converter.SpendingTypeModelToResponse(spendingType, formatMaximumLimit)

	return resp, nil
}

func (s *SpendingTypeUsecaseImpl) GetAllByProfileID(ctx context.Context, profileID string, periode int) (*[]domain.ResponseSpendingType, error) {
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}

	err = s.spendingTypeRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		exist, err := s.spendingTypeRepo.CheckData(ctx, profileID)
		if err != nil {
			return err
		}

		if !exist {
			spendingTypes, err := s.spendingTypeRepo.GetDefault(ctx)
			if err != nil {
				return err
			}

			for _, spendingType := range *spendingTypes {
				err = s.spendingTypeRepo.Create(ctx, &spendingType)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	req := &domain.RequestGetAllSpendingType{
		ProfileID: profileID,
		StartTime: time.Time{},
		EndTime:   time.Time{},
	}
	spendingTypes, err := s.spendingTypeRepo.GetAllByProfileID(ctx, req)

	var resps []domain.ResponseSpendingType
	var resp domain.ResponseSpendingType
	for _, spendingType := range *spendingTypes {
		formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)
		persentaseMaximumLimit := helper.Persentase(spendingType.Used, spendingType.MaximumLimit)
		resp = converter.SpendingTypeModelJoinToResponse(spendingType, persentaseMaximumLimit, formatMaximumLimit)
		resps = append(resps, resp)
	}

	return &resps, nil
}
