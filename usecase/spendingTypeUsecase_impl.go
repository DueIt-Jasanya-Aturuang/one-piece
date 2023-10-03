package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/oklog/ulid/v2"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/converter"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/helper"
)

type SpendingTypeUsecaseImpl struct {
	spendingTypeRepo    domain.SpendingTypeRepository
	spendingHistoryRepo domain.SpendingHistoryRepository
}

func NewSpendingTypeUsecaseImpl(
	spendingTypeRepo domain.SpendingTypeRepository,
	spendingHistoryRepo domain.SpendingHistoryRepository,
) domain.SpendingTypeUsecase {
	return &SpendingTypeUsecaseImpl{
		spendingTypeRepo:    spendingTypeRepo,
		spendingHistoryRepo: spendingHistoryRepo,
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
	defer s.spendingTypeRepo.CloseConn()

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

func (s *SpendingTypeUsecaseImpl) GetAllByPeriodeAndProfileID(ctx context.Context, req *domain.RequestGetAllSpendingTypeByTime) (*domain.ResponseAllSpendingType, string, error) {
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return nil, "", err
	}
	defer s.spendingTypeRepo.CloseConn()

	exist, err := s.spendingTypeRepo.CheckData(ctx, req.ProfileID)
	if err != nil {
		return nil, "", err
	}

	if !exist {
		err = s.createDefaultSpendingType(ctx, req.ProfileID)
		if err != nil {
			return nil, "", err
		}
	}

	req.StartTime, req.EndTime, err = helper.TimeDate(req.Periode)
	if err != nil {
		return nil, "", err
	}

	budgetAmount, err := s.spendingHistoryRepo.GetAllAmountByTimeAndProfileID(ctx, &domain.GetSpendingHistoryByTimeAndProfileID{
		ProfileID: req.ProfileID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		return nil, "", err
	}

	spendingTypes, err := s.spendingTypeRepo.GetAllByTimeAndProfileID(ctx, req)
	if err != nil {
		return nil, "", err
	}

	var spendingTypeJoinResponses []domain.ResponseSpendingTypeJoin
	var spendingTypeJoinResponse domain.ResponseSpendingTypeJoin
	var cursor string

	for _, spendingType := range *spendingTypes {
		formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)
		formatUsed := helper.FormatRupiah(spendingType.Used)
		persentaseMaximumLimit := helper.Persentase(spendingType.Used, spendingType.MaximumLimit)

		spendingTypeJoinResponse = converter.SpendingTypeModelJoinToResponse(spendingType, persentaseMaximumLimit, formatMaximumLimit, formatUsed)
		spendingTypeJoinResponses = append(spendingTypeJoinResponses, spendingTypeJoinResponse)

		cursor = spendingType.ID
	}

	respAll := &domain.ResponseAllSpendingType{
		ResponseSpendingType: &spendingTypeJoinResponses,
		BudgetAmount:         budgetAmount,
		FormatBudgetAmount:   helper.FormatRupiah(budgetAmount),
	}
	return respAll, cursor, nil
}

func (s *SpendingTypeUsecaseImpl) GetAllByProfileID(ctx context.Context, req *domain.RequestGetAllPaginate) (*[]domain.ResponseSpendingType, string, error) {
	err := s.spendingTypeRepo.OpenConn(ctx)
	if err != nil {
		return nil, "", err
	}
	defer s.spendingTypeRepo.CloseConn()

	exist, err := s.spendingTypeRepo.CheckData(ctx, req.ProfileID)
	if err != nil {
		return nil, "", err
	}

	if !exist {
		err = s.createDefaultSpendingType(ctx, req.ProfileID)
		if err != nil {
			return nil, "", err
		}
	}

	spendingTypes, err := s.spendingTypeRepo.GetAllByProfileID(ctx, req)
	if err != nil {
		return nil, "", err
	}

	var spendingTypeResponses []domain.ResponseSpendingType
	var spendingTypeResponse domain.ResponseSpendingType
	var cursor string
	for _, spendingType := range *spendingTypes {
		formatMaximumLimit := helper.FormatRupiah(spendingType.MaximumLimit)

		spendingTypeResponse = *converter.SpendingTypeModelToResponse(&spendingType, formatMaximumLimit)
		spendingTypeResponses = append(spendingTypeResponses, spendingTypeResponse)

		cursor = spendingType.ID
	}

	return &spendingTypeResponses, cursor, nil
}

func (s *SpendingTypeUsecaseImpl) createDefaultSpendingType(ctx context.Context, profileID string) error {
	err := s.spendingTypeRepo.StartTx(ctx, helper.LevelReadCommitted(), func() error {
		spendingTypes, err := s.spendingTypeRepo.GetDefault(ctx)
		if err != nil {
			return err
		}

		for _, spendingType := range *spendingTypes {
			spendingType.ProfileID = profileID
			spendingType.ID = ulid.Make().String()
			err = s.spendingTypeRepo.Create(ctx, &spendingType)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
