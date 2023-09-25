package rest

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/validation"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/_usecase"
)

type IncomeTypeHandlerImpl struct {
	incomeTypeUsecase domain.IncomeTypeUsecase
}

func NewIncomeTypeHandlerImpl(incomeTypeUsecase domain.IncomeTypeUsecase) *IncomeTypeHandlerImpl {
	return &IncomeTypeHandlerImpl{
		incomeTypeUsecase: incomeTypeUsecase,
	}
}

func (i *IncomeTypeHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreateIncomeType)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}
	req.ProfileID = r.Header.Get("Profile-ID")

	err = validation.CreateIncomeType(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	incomeType, err := i.incomeTypeUsecase.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.NameIncomeTypeIsExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"name": {
					"name pemasukan kategori sudah tersedia",
				},
			}, response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, incomeType, "created pemasukan kategori sukses")
}

func (i *IncomeTypeHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestUpdateIncomeType)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}
	req.ID = chi.URLParam(r, "id")
	req.ProfileID = r.Header.Get("Profile-ID")

	err = validation.UpdateIncomeType(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	incomeType, err := i.incomeTypeUsecase.Update(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.NameIncomeTypeIsExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"name": {
					"name pemasukan kategori sudah tersedia",
				},
			}, response.CM06)
		}
		if errors.Is(err, _usecase.IncomeTypeIsNotExist) {
			err = _error.HttpErrString("pemasukan kategori tidak ditemukan", response.CM01)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, incomeType, "update pemasukan kategori sukses")
}

func (i *IncomeTypeHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(id); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}
	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM05))
		return
	}

	err := i.incomeTypeUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "deleted pemasukan kategori sukses")
}

func (i *IncomeTypeHandlerImpl) GetByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(id); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}
	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM05))
		return
	}

	resp, err := i.incomeTypeUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, _usecase.IncomeTypeIsNotExist) {
			err = _error.HttpErrString("data pemasukan kategori tidak ditemukan", response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, resp, "data pemasukan kategori")
}

func (i *IncomeTypeHandlerImpl) GetAllByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM05))
		return
	}

	resps, err := i.incomeTypeUsecase.GetAllByProfileID(r.Context(), profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, resps, "data pemasukan kategori")
}
