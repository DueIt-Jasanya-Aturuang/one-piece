package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"

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
	// TODO implement me
	panic("implement me")
}

func (h *SpendingHistoryHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (h *SpendingHistoryHandlerImpl) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}
