package spendingHistory_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (s *SpendingHistoryUsecaseImpl) GetAllByTimeAndProfileID(ctx context.Context, req *usecase.RequestGetAllSpendingHistoryWithISD) (*[]usecase.ResponseSpendingHistory, string, error) {
	var err error
	if req.Type != "" {
		req.StartTime, req.EndTime, _ = usecase.TimeDateByTypeFilter(req.Type)
	} else {
		req.StartTime, req.EndTime, err = usecase.FormatDate(req.StartTime, req.EndTime)
		if err != nil {
			return nil, "", usecase.InvalidTimestamp
		}
	}

	order, operation := usecase.GetOrder(req.Order)
	spendingHistories, err := s.spendingHistoryRepo.GetAllByTimeAndProfileID(ctx, &repository.GetAllSpendingHistoryByFilterWithISD{
		ProfileID: req.ProfileID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Type:      req.Type,
		InfiniteScrollData: repository.InfiniteScrollData{
			ID:        req.ID,
			Order:     order,
			Operation: operation,
		},
	})
	if err != nil {
		return nil, "", err
	}

	if len(*spendingHistories) < 1 {
		return nil, "", nil
	}

	var spendingHistoryJoinResponses []usecase.ResponseSpendingHistory
	var spendingHistoryJoinResponse *usecase.ResponseSpendingHistory

	for _, spendingHistory := range *spendingHistories {
		spendingHistoryJoinResponse = &usecase.ResponseSpendingHistory{
			ID:                      spendingHistory.ID,
			ProfileID:               spendingHistory.ProfileID,
			SpendingTypeID:          spendingHistory.SpendingTypeID,
			SpendingTypeTitle:       spendingHistory.SpendingTypeTitle,
			PaymentMethodID:         repository.GetNullString(spendingHistory.PaymentMethodID),
			PaymentMethodName:       repository.GetNullString(spendingHistory.PaymentMethodName),
			PaymentName:             repository.GetNullString(spendingHistory.PaymentName),
			BeforeBalance:           spendingHistory.BeforeBalance,
			SpendingAmount:          spendingHistory.SpendingAmount,
			FormatSpendingAmount:    usecase.FormatRupiah(spendingHistory.SpendingAmount),
			AfterBalance:            spendingHistory.AfterBalance,
			Description:             spendingHistory.Description,
			TimeSpendingHistory:     spendingHistory.TimeSpendingHistory,
			ShowTimeSpendingHistory: spendingHistory.ShowTimeSpendingHistory,
		}
		spendingHistoryJoinResponses = append(spendingHistoryJoinResponses, *spendingHistoryJoinResponse)

	}

	cursor := (*spendingHistories)[len(*spendingHistories)-1].ID
	return &spendingHistoryJoinResponses, cursor, nil
}

func (s *SpendingHistoryUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*usecase.ResponseSpendingHistory, error) {
	spendingHistory, err := s.spendingHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.SpendingHistoryNotFound
		}
		return nil, err
	}

	resp := &usecase.ResponseSpendingHistory{
		ID:                      spendingHistory.ID,
		ProfileID:               spendingHistory.ProfileID,
		SpendingTypeID:          spendingHistory.SpendingTypeID,
		SpendingTypeTitle:       spendingHistory.SpendingTypeTitle,
		PaymentMethodID:         repository.GetNullString(spendingHistory.PaymentMethodID),
		PaymentMethodName:       repository.GetNullString(spendingHistory.PaymentMethodName),
		PaymentName:             repository.GetNullString(spendingHistory.PaymentName),
		BeforeBalance:           spendingHistory.BeforeBalance,
		SpendingAmount:          spendingHistory.SpendingAmount,
		FormatSpendingAmount:    usecase.FormatRupiah(spendingHistory.SpendingAmount),
		AfterBalance:            spendingHistory.AfterBalance,
		Description:             spendingHistory.Description,
		TimeSpendingHistory:     spendingHistory.TimeSpendingHistory,
		ShowTimeSpendingHistory: spendingHistory.ShowTimeSpendingHistory,
	}

	return resp, nil
}
