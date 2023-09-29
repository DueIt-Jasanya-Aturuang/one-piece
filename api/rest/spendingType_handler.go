package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/validation"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/_usecase"
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

	req.ProfileID = r.Header.Get("Profile-ID")
	err = validation.CreateSpendingType(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	spendingType, err := h.spendingTypeUsecase.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.TitleSpendingTypeExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"title": {
					err.Error(),
				},
			}, response.CM06)
		}
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

	req.ProfileID = r.Header.Get("Profile-ID")
	req.ID = chi.URLParam(r, "id")

	err = validation.UpdateSpendingType(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	spendingType, err := h.spendingTypeUsecase.Update(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.TitleSpendingTypeExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"title": {
					err.Error(),
				},
			}, response.CM06)
		}
		if errors.Is(err, _usecase.SpendingTypeNotFound) {
			err = _error.HttpErrString("not found", response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingType, "update spending type berhasil")
}

func (h *SpendingTypeHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")
	id := chi.URLParam(r, "id")

	_, err := uuid.Parse(profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}
	_, err = uuid.Parse(id)
	if err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
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
	profileID := r.Header.Get("Profile-ID")
	id := chi.URLParam(r, "id")

	log.Info().Msgf("%s | %s", profileID, id)
	_, err := uuid.Parse(profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}
	_, err = uuid.Parse(id)
	if err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}

	spendingType, err := h.spendingTypeUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, _usecase.SpendingTypeNotFound) {
			err = _error.HttpErrString("not found", response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingType, "data spending type")
}

func (h *SpendingTypeHandlerImpl) GetAllByPeriodeAndProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	_, err := uuid.Parse(profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}
	periode := chi.URLParam(r, "periode")
	periodInt, err := strconv.Atoi(periode)
	if err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}

	if periodInt > 29 || periodInt < 1 {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}

	spendingTypes, err := h.spendingTypeUsecase.GetAllByPeriodeAndProfileID(r.Context(), profileID, periodInt)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingTypes, "data spending types")
}

func (h *SpendingTypeHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	_, err := uuid.Parse(profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}

	spendingTypes, err := h.spendingTypeUsecase.GetAllByProfileID(r.Context(), profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, spendingTypes, "data spending types")
}
