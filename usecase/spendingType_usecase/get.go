package spendingType_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (s *SpendingTypeUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*usecase.ResponseSpendingType, error) {
	spendingType, err := s.spendingTypeRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.SpendingTypeNotFound
		}
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

func (s *SpendingTypeUsecaseImpl) GetAllByPeriodeAndProfileID(ctx context.Context, req *usecase.RequestGetAllSpendingTypeByPeriodeWithISD) (*usecase.ResponseAllSpendingType, string, error) {
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

	startTime, endTime, err := usecase.TimeDate(req.Periode)
	if err != nil {
		return nil, "", err
	}

	budgetAmount, err := s.spendingHistoryRepo.GetTotalAmountByPeriode(ctx, &repository.GetTotalSpendingHistoryByPeriode{
		ProfileID: req.ProfileID,
		StartTime: startTime,
		EndTime:   endTime,
	})
	if err != nil {
		return nil, "", err
	}

	order, operation := usecase.GetOrder(req.Order)
	spendingTypes, err := s.spendingTypeRepo.GetAllByTimeAndProfileID(ctx, &repository.GetAllSpendingTypeByFilterWithISD{
		ProfileID: req.ProfileID,
		StartTime: startTime,
		EndTime:   endTime,
		InfiniteScrollData: repository.InfiniteScrollData{
			ID:        req.ID,
			Order:     order,
			Operation: operation,
		},
	})
	if err != nil {
		return nil, "", err
	}

	if len(*spendingTypes) < 1 {
		return nil, "", nil
	}
	var spendingTypeJoinResponses []usecase.ResponseSpendingTypeJoinTable
	var spendingTypeJoinResponse *usecase.ResponseSpendingTypeJoinTable

	for _, spendingType := range *spendingTypes {
		spendingTypeJoinResponse = &usecase.ResponseSpendingTypeJoinTable{
			ID:                 spendingType.ID,
			ProfileID:          spendingType.ProfileID,
			Title:              spendingType.Title,
			MaximumLimit:       spendingType.MaximumLimit,
			FormatMaximumLimit: usecase.FormatRupiah(spendingType.MaximumLimit),
			Icon:               spendingType.Icon,
			Used:               spendingType.Used,
			FormatUsed:         usecase.FormatRupiah(spendingType.Used),
			PersentaseUsed:     usecase.Persentase(spendingType.Used, spendingType.MaximumLimit),
		}
		spendingTypeJoinResponses = append(spendingTypeJoinResponses, *spendingTypeJoinResponse)

	}

	respAll := &usecase.ResponseAllSpendingType{
		ResponseSpendingType: &spendingTypeJoinResponses,
		BudgetAmount:         budgetAmount,
		FormatBudgetAmount:   usecase.FormatRupiah(budgetAmount),
	}
	cursor := (*spendingTypes)[len(*spendingTypes)-1].ID
	return respAll, cursor, nil
}

func (s *SpendingTypeUsecaseImpl) GetAllByProfileID(ctx context.Context, req *usecase.RequestGetAllByProfileIDWithISD) (*[]usecase.ResponseSpendingType, string, error) {
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

	order, operation := usecase.GetOrder(req.Order)
	spendingTypes, err := s.spendingTypeRepo.GetAllByProfileID(ctx, &repository.GetAllSpendingTypeWithISD{
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

	if len(*spendingTypes) < 1 {
		return nil, "", nil
	}

	var spendingTypeResponses []usecase.ResponseSpendingType
	var spendingTypeResponse *usecase.ResponseSpendingType

	for _, spendingType := range *spendingTypes {
		spendingTypeResponse = &usecase.ResponseSpendingType{
			ID:                 spendingType.ID,
			ProfileID:          spendingType.ProfileID,
			Title:              spendingType.Title,
			MaximumLimit:       spendingType.MaximumLimit,
			FormatMaximumLimit: usecase.FormatRupiah(spendingType.MaximumLimit),
			Icon:               spendingType.Icon,
		}
		spendingTypeResponses = append(spendingTypeResponses, *spendingTypeResponse)

	}

	cursor := (*spendingTypes)[len(*spendingTypes)-1].ID
	return &spendingTypeResponses, cursor, nil
}
