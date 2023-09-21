package _usecase

import (
	"context"
	"database/sql"
	"errors"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/helper"
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
	defer s.spendingTypeRepo.CloseConn()

	exist, err := s.spendingTypeRepo.CheckByTitleAndProfileID(ctx, req.ProfileID, req.Title)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, TitleSpendingTypeExist
	}

	spendingType := converter.SpendingTypeRequestCreateToModel(req)
	err = s.spendingTypeRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		err = s.spendingTypeRepo.Create(ctx, spendingType)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)
	spendingTypeResponse := converter.SpendingTypeModelToResponse(spendingType, formatMaximumLimit)

	return spendingTypeResponse, nil
}

func (s *SpendingTypeUsecaseImpl) Update(ctx context.Context, req *domain.RequestUpdateSpendingType) (*domain.ResponseSpendingType, error) {
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}
	defer s.spendingTypeRepo.CloseConn()

	spendingType, err := s.spendingTypeRepo.GetByIDAndProfileID(ctx, req.ID, req.ProfileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, SpendingTypeNotFound
		}
		return nil, err
	}

	if spendingType.Title != req.Title {
		exist, err := s.spendingTypeRepo.CheckByTitleAndProfileID(ctx, req.ProfileID, req.Title)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, TitleSpendingTypeExist
		}
	}

	spendingType = converter.SpendingTypeRequestUpdateToModel(req)
	err = s.spendingTypeRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		err = s.spendingTypeRepo.Update(ctx, spendingType)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)
	spendingTypeResponse := converter.SpendingTypeModelToResponse(spendingType, formatMaximumLimit)

	return spendingTypeResponse, nil
}

func (s *SpendingTypeUsecaseImpl) Delete(ctx context.Context, id string, profileID string) error {
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return err
	}
	defer s.spendingTypeRepo.CloseConn()

	err = s.spendingTypeRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		err = s.spendingTypeRepo.Delete(ctx, id, profileID)
		if err != nil {
			return err
		}

		return nil
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
			return nil, SpendingTypeNotFound
		}
		return nil, err
	}

	formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)
	spendingTypeResponse := converter.SpendingTypeModelToResponse(spendingType, formatMaximumLimit)

	return spendingTypeResponse, nil
}

func (s *SpendingTypeUsecaseImpl) GetAllByProfileID(ctx context.Context, profileID string, periode int) (*domain.ResponseAllSpendingType, error) {
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return nil, err
	}

	exist, err := s.spendingTypeRepo.CheckData(ctx, profileID)
	if err != nil {
		return nil, err
	}

	if !exist {
		err = s.spendingTypeRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
			spendingTypes, err := s.spendingTypeRepo.GetDefault(ctx)
			if err != nil {
				return err
			}

			for _, spendingType := range *spendingTypes {
				spendingType.ProfileID = profileID
				spendingType.ID = uuid.NewV4().String()
				err = s.spendingTypeRepo.Create(ctx, &spendingType)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	startTime, endTime, err := helper.TimeDate(periode)
	if err != nil {
		return nil, err
	}

	req := &domain.RequestGetAllSpendingType{
		ProfileID: profileID,
		StartTime: startTime,
		EndTime:   endTime,
	}

	spendingTypes, err := s.spendingTypeRepo.GetAllByProfileID(ctx, req)
	if err != nil {
		return nil, err
	}

	var spendingTypeResponses []domain.ResponseSpendingType
	var spendingTypeResponse domain.ResponseSpendingType
	var budgetAmount int

	for _, spendingType := range *spendingTypes {
		budgetAmount += spendingType.Used
		formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)
		formatUsed := helper.FormatRupiah(spendingType.Used)
		persentaseMaximumLimit := helper.Persentase(spendingType.Used, spendingType.MaximumLimit)

		spendingTypeResponse = converter.SpendingTypeModelJoinToResponse(spendingType, persentaseMaximumLimit, formatMaximumLimit, formatUsed)
		spendingTypeResponses = append(spendingTypeResponses, spendingTypeResponse)
	}

	respAll := &domain.ResponseAllSpendingType{
		ResponseSpendingType: &spendingTypeResponses,
		BudgetAmount:         budgetAmount,
		FormatBudgetAmount:   helper.FormatRupiah(budgetAmount),
	}
	return respAll, nil
}