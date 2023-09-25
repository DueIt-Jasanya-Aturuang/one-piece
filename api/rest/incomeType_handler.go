package rest

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
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
