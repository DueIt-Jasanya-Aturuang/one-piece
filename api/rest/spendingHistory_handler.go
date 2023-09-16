package rest

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/validation"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
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

	err = validation.CreateSpendingHistory(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	spendingHistory, err := h.spendingHistoryUsecase.Create(r.Context(), req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := domain.ResponseSuccessHTTP{
		Data: spendingHistory,
		Code: 201,
	}

	helper.SuccessResponseEncode(w, resp)
}

func (h *SpendingHistoryHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestUpdateSpendingHistory)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}
	req.ID = chi.URLParam(r, "id")

	err = validation.UpdateSpendingHistory(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	spendingHistory, err := h.spendingHistoryUsecase.Update(r.Context(), req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := domain.ResponseSuccessHTTP{
		Data: spendingHistory,
		Code: 201,
	}

	helper.SuccessResponseEncode(w, resp)
}

func (h *SpendingHistoryHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile-id")
	id := chi.URLParam(r, "id")

	err := h.spendingHistoryUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := domain.ResponseSuccessHTTP{
		Message: "successfully deleted data",
		Code:    200,
	}

	helper.SuccessResponseEncode(w, resp)
}

func (h *SpendingHistoryHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	typeQuery := r.URL.Query().Get("type")
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")

	var startTime time.Time
	var endTime time.Time

	startTime, _ = time.Parse("2006-01-02", start)
	endTime, _ = time.Parse("2006-01-02", end)

	req := &domain.RequestGetFilteredDataSpendingHistory{
		ProfileID: chi.URLParam(r, "profile-id"),
		StartTime: startTime,
		EndTime:   endTime,
		Type:      typeQuery,
	}

	log.Info().Msgf("%v", req)

	spendingHistories, err := h.spendingHistoryUsecase.GetAllByTimeAndProfileID(r.Context(), req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := domain.ResponseSuccessHTTP{
		Data: spendingHistories,
		Code: 200,
	}

	helper.SuccessResponseEncode(w, resp)
}

func (h *SpendingHistoryHandlerImpl) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := chi.URLParam(r, "profile-id")

	spendingHistory, err := h.spendingHistoryUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := domain.ResponseSuccessHTTP{
		Data: spendingHistory,
		Code: 200,
	}

	helper.SuccessResponseEncode(w, resp)
}
