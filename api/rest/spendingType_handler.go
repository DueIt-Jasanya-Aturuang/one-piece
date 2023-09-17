package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	resp "github.com/jasanya-tech/jasanya-response-backend-golang"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/validation"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

type SpendingTypeHandlerImpl struct {
	spendingTypeUsecase domain.SpendingTypeUsecase
}

func NewSpendingTypeHandlerImpl(
	spendingTypeUsecase domain.SpendingTypeUsecase,
) *SpendingTypeHandlerImpl {
	return &SpendingTypeHandlerImpl{
		spendingTypeUsecase: spendingTypeUsecase,
	}
}

func (h *SpendingTypeHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreateSpendingType)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	err = validation.CreateSpendingType(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	spendingType, err := h.spendingTypeUsecase.Create(r.Context(), req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingType, "crete spending type berhasil")
}

func (h *SpendingTypeHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestUpdateSpendingType)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}
	req.ID = chi.URLParam(r, "id")

	err = validation.UpdateSpendingType(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	spendingType, err := h.spendingTypeUsecase.Update(r.Context(), req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingType, "update spending type berhasil")
}

func (h *SpendingTypeHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile-id")
	id := chi.URLParam(r, "id")

	_, err := uuid.Parse(profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, resp.HttpErrString(string(resp.S404), resp.S404))
		return
	}
	_, err = uuid.Parse(id)
	if err != nil {
		helper.ErrorResponseEncode(w, resp.HttpErrString(string(resp.S404), resp.S404))
		return
	}

	err = h.spendingTypeUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "deleted spending type berhasil")
}

func (h *SpendingTypeHandlerImpl) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile-id")
	id := chi.URLParam(r, "id")

	_, err := uuid.Parse(profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, resp.HttpErrString(string(resp.S404), resp.S404))
		return
	}
	_, err = uuid.Parse(id)
	if err != nil {
		helper.ErrorResponseEncode(w, resp.HttpErrString(string(resp.S404), resp.S404))
		return
	}

	spendingType, err := h.spendingTypeUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingType, "data spending type")
}

func (h *SpendingTypeHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := chi.URLParam(r, "profile-id")

	_, err := uuid.Parse(profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, resp.HttpErrString(string(resp.S404), resp.S404))
		return
	}
	periode := r.URL.Query().Get("periode")
	periodInt, err := strconv.Atoi(periode)
	if err != nil {
		helper.ErrorResponseEncode(w, resp.HttpErrString(string(resp.S404), resp.S404))
		return
	}

	spendingTypes, err := h.spendingTypeUsecase.GetAllByProfileID(r.Context(), profileID, periodInt)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingTypes, "data spending types")
}
