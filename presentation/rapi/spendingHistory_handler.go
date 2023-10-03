package rapi

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/validation"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

type SpendingHistoryHandlerImpl struct {
	spendingHistoryUsecase domain.SpendingHistoryUsecase
}

func NewSpendingHistoryHandlerImpl(
	spendingHistoryUsecase domain.SpendingHistoryUsecase,
) *SpendingHistoryHandlerImpl {
	return &SpendingHistoryHandlerImpl{
		spendingHistoryUsecase: spendingHistoryUsecase,
	}
}

func (h *SpendingHistoryHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreateSpendingHistory)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}
	req.ProfileID = r.Header.Get("Profile-ID")

	err = validation.CreateSpendingHistory(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	spendingHistory, err := h.spendingHistoryUsecase.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.SpendingHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		if errors.Is(err, usecase.InvalidSpendingTypeID) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"spending_type_id": {
					err.Error(),
				},
			}, response.CM06)
		}
		if errors.Is(err, usecase.InvalidPaymentMethodID) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"payment_method_id": {
					err.Error(),
				},
			}, response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingHistory, "create spending history berhasil")
}

func (h *SpendingHistoryHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestUpdateSpendingHistory)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	req.ProfileID = r.Header.Get("Profile-ID")
	req.ID = chi.URLParam(r, "id")

	err = validation.UpdateSpendingHistory(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	spendingHistory, err := h.spendingHistoryUsecase.Update(r.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.SpendingHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		if errors.Is(err, usecase.InvalidSpendingTypeID) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"spending_type_id": {
					err.Error(),
				},
			}, response.CM06)
		}
		if errors.Is(err, usecase.InvalidPaymentMethodID) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"payment_method_id": {
					err.Error(),
				},
			}, response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingHistory, "update spending history berhasil")
}

func (h *SpendingHistoryHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")
	id := chi.URLParam(r, "id")

	err := h.spendingHistoryUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.SpendingHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		if errors.Is(err, usecase.ProfileIDNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "delete spending history berhasil")
}

func (h *SpendingHistoryHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	typeQuery := r.URL.Query().Get("type")
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	var startTime time.Time
	var endTime time.Time

	startTime, _ = time.Parse("2006-01-02", start)
	endTime, _ = time.Parse("2006-01-02", end)

	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")
	order, operation := helper.GetOrder(order)

	req := &domain.GetFilteredDataSpendingHistory{
		ProfileID: r.Header.Get("Profile-ID"),
		StartTime: startTime,
		EndTime:   endTime,
		Type:      typeQuery,
		RequestGetAllPaginate: domain.RequestGetAllPaginate{
			ProfileID: r.Header.Get("Profile-ID"),
			ID:        cursor,
			Operation: operation,
			Order:     order,
		},
	}

	spendingHistories, cursorResp, err := h.spendingHistoryUsecase.GetAllByTimeAndProfileID(r.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.InvalidTimestamp) {
			err = _error.HttpErrString(err.Error(), response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := map[string]any{
		"cursor":           cursorResp,
		"spending_history": spendingHistories,
	}
	helper.SuccessResponseEncode(w, resp, "data spending history")
}

func (h *SpendingHistoryHandlerImpl) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	spendingHistory, err := h.spendingHistoryUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.SpendingHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingHistory, "data spending history")
}
