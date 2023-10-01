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

type IncomeHistoryHandlerImpl struct {
	incomeHistoryUsecase domain.IncomeHistoryUsecase
}

func NewIncomeHistoryHandlerImpl(
	incomeHistoryUsecase domain.IncomeHistoryUsecase,
) *IncomeHistoryHandlerImpl {
	return &IncomeHistoryHandlerImpl{
		incomeHistoryUsecase: incomeHistoryUsecase,
	}
}

func (h *IncomeHistoryHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreateIncomeHistory)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}
	req.ProfileID = r.Header.Get("Profile-ID")

	err = validation.CreateIncomeHistory(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	incomeHistory, err := h.incomeHistoryUsecase.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.IncomeHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		if errors.Is(err, usecase.InvalidIncomeTypeID) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"income_type_id": {
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

	helper.SuccessResponseEncode(w, incomeHistory, "create income history berhasil")
}

func (h *IncomeHistoryHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestUpdateIncomeHistory)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	req.ProfileID = r.Header.Get("Profile-ID")
	req.ID = chi.URLParam(r, "id")

	err = validation.UpdateIncomeHistory(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	incomeHistory, err := h.incomeHistoryUsecase.Update(r.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.IncomeHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		if errors.Is(err, usecase.InvalidIncomeTypeID) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"income_type_id": {
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

	helper.SuccessResponseEncode(w, incomeHistory, "update income history berhasil")
}

func (h *IncomeHistoryHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")
	id := chi.URLParam(r, "id")

	err := h.incomeHistoryUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.IncomeHistoryNotFound) {
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

func (h *IncomeHistoryHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	typeQuery := r.URL.Query().Get("type")
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	var startTime time.Time
	var endTime time.Time

	startTime, _ = time.Parse("2006-01-02", start)
	endTime, _ = time.Parse("2006-01-02", end)

	req := &domain.GetFilteredDataIncomeHistory{
		ProfileID: r.Header.Get("Profile-ID"),
		StartTime: startTime,
		EndTime:   endTime,
		Type:      typeQuery,
	}

	incomeHistories, err := h.incomeHistoryUsecase.GetAllByTimeAndProfileID(r.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.InvalidTimestamp) {
			err = _error.HttpErrString(err.Error(), response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, incomeHistories, "data income history")
}

func (h *IncomeHistoryHandlerImpl) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	incomeHistory, err := h.incomeHistoryUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.IncomeHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, incomeHistory, "data income history")
}