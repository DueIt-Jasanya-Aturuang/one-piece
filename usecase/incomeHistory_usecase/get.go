package incomeHistory_usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

func (i *IncomeHistoryUsecaseImpl) GetAllByTimeAndProfileID(ctx context.Context, req *usecase.RequestGetAllIncomeHistoryWithISD) (*[]usecase.ResponseIncomeHistory, string, error) {
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

	incomeHistories, err := i.incomeHistoryRepo.GetAllByTimeAndProfileID(ctx, &repository.GetAllIncomeHistoryByTimeFilterWithISD{
		ProfileID: req.ProfileID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		InfiniteScrollData: repository.InfiniteScrollData{
			ID:        req.ID,
			Order:     order,
			Operation: operation,
		},
	})
	if err != nil {
		return nil, "", err
	}

	if len(*incomeHistories) < 1 {
		return nil, "", nil
	}

	var incomeHistoryJoinResponses []usecase.ResponseIncomeHistory
	var incomeHistoryJoinResponse *usecase.ResponseIncomeHistory

	for _, incomeHistory := range *incomeHistories {
		incomeHistoryJoinResponse = &usecase.ResponseIncomeHistory{
			ID:                    incomeHistory.ID,
			ProfileID:             incomeHistory.ProfileID,
			IncomeTypeID:          incomeHistory.IncomeTypeID,
			IncomeTypeTitle:       incomeHistory.IncomeTypeTitle,
			PaymentMethodID:       repository.GetNullString(incomeHistory.PaymentMethodID),
			PaymentMethodName:     repository.GetNullString(incomeHistory.PaymentMethodName),
			PaymentName:           repository.GetNullString(incomeHistory.PaymentName),
			IncomeAmount:          incomeHistory.IncomeAmount,
			FormatIncomeAmount:    usecase.FormatRupiah(incomeHistory.IncomeAmount),
			Description:           incomeHistory.Description,
			TimeIncomeHistory:     incomeHistory.TimeIncomeHistory,
			ShowTimeIncomeHistory: incomeHistory.ShowTimeIncomeHistory,
		}
		incomeHistoryJoinResponses = append(incomeHistoryJoinResponses, *incomeHistoryJoinResponse)

	}

	cursor := (*incomeHistories)[len(*incomeHistories)-1].ID
	return &incomeHistoryJoinResponses, cursor, nil
}

func (i *IncomeHistoryUsecaseImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*usecase.ResponseIncomeHistory, error) {
	incomeHistory, err := i.incomeHistoryRepo.GetByIDAndProfileID(ctx, id, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.IncomeHistoryNotFound
		}
		return nil, err
	}

	resp := &usecase.ResponseIncomeHistory{
		ID:                    incomeHistory.ID,
		ProfileID:             incomeHistory.ProfileID,
		IncomeTypeID:          incomeHistory.IncomeTypeID,
		IncomeTypeTitle:       incomeHistory.IncomeTypeTitle,
		PaymentMethodID:       repository.GetNullString(incomeHistory.PaymentMethodID),
		PaymentMethodName:     repository.GetNullString(incomeHistory.PaymentMethodName),
		PaymentName:           repository.GetNullString(incomeHistory.PaymentName),
		IncomeAmount:          incomeHistory.IncomeAmount,
		FormatIncomeAmount:    usecase.FormatRupiah(incomeHistory.IncomeAmount),
		Description:           incomeHistory.Description,
		TimeIncomeHistory:     incomeHistory.TimeIncomeHistory,
		ShowTimeIncomeHistory: incomeHistory.ShowTimeIncomeHistory,
	}

	return resp, nil
}
