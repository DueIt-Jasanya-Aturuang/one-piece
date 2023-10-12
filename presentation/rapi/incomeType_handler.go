package rapi

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/schema"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (p *Presenter) CreateIncomeType(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	req := new(schema.RequestCreateOrUpdateIncomeType)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	err = req.Validate()
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	incomeType, err := p.incomeTypeUsecase.Create(r.Context(), &usecase.RequestCreateIncomeType{
		ProfileID:   profileID,
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	})
	if err != nil {
		if errors.Is(err, usecase.NameIncomeTypeIsExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"name": {
					"name pemasukan kategori sudah tersedia",
				},
			}, response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseIncomeType{
		ID:          incomeType.ID,
		ProfileID:   incomeType.ProfileID,
		Name:        incomeType.Name,
		Description: incomeType.Description,
		Icon:        incomeType.Icon,
	}
	helper.SuccessResponseEncode(w, resp, "created pemasukan kategori sukses")
}

func (p *Presenter) UpdateIncomeType(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")
	id := chi.URLParam(r, "id")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	if errParse := util.ParseUlid(id); errParse != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01))
		return
	}

	req := new(schema.RequestCreateOrUpdateIncomeType)

	err := helper.DecodeJson(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	err = req.Validate()
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	incomeType, err := p.incomeTypeUsecase.Update(r.Context(), &usecase.RequestUpdateIncomeType{
		ID:          id,
		ProfileID:   profileID,
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	})
	if err != nil {
		if errors.Is(err, usecase.NameIncomeTypeIsExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"name": {
					"name pemasukan kategori sudah tersedia",
				},
			}, response.CM06)
		}
		if errors.Is(err, usecase.IncomeTypeIsNotExist) {
			err = _error.HttpErrString("pemasukan kategori tidak ditemukan", response.CM01)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseIncomeType{
		ID:          incomeType.ID,
		ProfileID:   incomeType.ProfileID,
		Name:        incomeType.Name,
		Description: incomeType.Description,
		Icon:        incomeType.Icon,
	}
	helper.SuccessResponseEncode(w, resp, "update pemasukan kategori sukses")
}

func (p *Presenter) DeleteIncomeType(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")
	id := chi.URLParam(r, "id")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	if errParse := util.ParseUlid(id); errParse != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01))
		return
	}

	err := p.incomeTypeUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "deleted pemasukan kategori sukses")
}

func (p *Presenter) GetIncomeTypeByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")
	id := chi.URLParam(r, "id")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	if errParse := util.ParseUlid(id); errParse != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01))
		return
	}

	incomeType, err := p.incomeTypeUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.IncomeTypeIsNotExist) {
			err = _error.HttpErrString("data pemasukan kategori tidak ditemukan", response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseIncomeType{
		ID:          incomeType.ID,
		ProfileID:   incomeType.ProfileID,
		Name:        incomeType.Name,
		Description: incomeType.Description,
		Icon:        incomeType.Icon,
	}
	helper.SuccessResponseEncode(w, resp, "data pemasukan kategori")
}

func (p *Presenter) GetAllIncomeTypeByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")

	incomeTypes, cursorResp, err := p.incomeTypeUsecase.GetAllByProfileID(r.Context(), &usecase.RequestGetAllByProfileIDWithISD{
		ID:        cursor,
		ProfileID: profileID,
		Order:     order,
	})
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	if incomeTypes == nil {
		helper.SuccessResponseEncode(w, map[string]any{
			"cursor":      "",
			"income_type": nil,
		}, "data income type")
		return
	}

	var responses []schema.ResponseIncomeType
	var responseIncomeType *schema.ResponseIncomeType

	for _, incomeType := range *incomeTypes {
		responseIncomeType = &schema.ResponseIncomeType{
			ID:          incomeType.ID,
			ProfileID:   incomeType.ProfileID,
			Name:        incomeType.Name,
			Description: incomeType.Description,
			Icon:        incomeType.Icon,
		}

		responses = append(responses, *responseIncomeType)
	}
	resp := map[string]any{
		"cursor":      cursorResp,
		"income_type": &responses,
	}

	helper.SuccessResponseEncode(w, resp, "data pemasukan kategori")
}
