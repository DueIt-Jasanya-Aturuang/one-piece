package rapi

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/schema"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (p *Presenter) CreateSpendingType(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	req := new(schema.RequestCreateOrUpdateSpendingType)

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

	spendingType, err := p.spendingTypeUsecase.Create(r.Context(), &usecase.RequestCreateSpendingType{
		ProfileID:    profileID,
		Title:        req.Title,
		MaximumLimit: req.MaximumLimit,
		Icon:         req.Icon,
	})
	if err != nil {
		if errors.Is(err, usecase.TitleSpendingTypeExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"title": {
					err.Error(),
				},
			}, response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseSpendingType{
		ID:                 spendingType.ID,
		ProfileID:          spendingType.ProfileID,
		Title:              spendingType.Title,
		MaximumLimit:       spendingType.MaximumLimit,
		FormatMaximumLimit: spendingType.FormatMaximumLimit,
		Icon:               spendingType.Icon,
	}
	helper.SuccessResponseEncode(w, resp, "crete spending type berhasil")
}

func (p *Presenter) UpdateSpendingType(w http.ResponseWriter, r *http.Request) {
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

	req := new(schema.RequestCreateOrUpdateSpendingType)

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

	spendingType, err := p.spendingTypeUsecase.Update(r.Context(), &usecase.RequestUpdateSpendingType{
		ID:           id,
		ProfileID:    profileID,
		Title:        req.Title,
		MaximumLimit: req.MaximumLimit,
		Icon:         req.Icon,
	})
	if err != nil {
		if errors.Is(err, usecase.TitleSpendingTypeExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"title": {
					err.Error(),
				},
			}, response.CM06)
		}
		if errors.Is(err, usecase.SpendingTypeNotFound) {
			err = _error.HttpErrString("not found", response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseSpendingType{
		ID:                 spendingType.ID,
		ProfileID:          spendingType.ProfileID,
		Title:              spendingType.Title,
		MaximumLimit:       spendingType.MaximumLimit,
		FormatMaximumLimit: spendingType.FormatMaximumLimit,
		Icon:               spendingType.Icon,
	}
	helper.SuccessResponseEncode(w, resp, "update spending type berhasil")
}

func (p *Presenter) DeleteSpendingType(w http.ResponseWriter, r *http.Request) {
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

	err := p.spendingTypeUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "deleted spending type berhasil")
}

func (p *Presenter) GetSpendingTypeByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
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

	spendingType, err := p.spendingTypeUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.SpendingTypeNotFound) {
			err = _error.HttpErrString("not found", response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseSpendingType{
		ID:                 spendingType.ID,
		ProfileID:          spendingType.ProfileID,
		Title:              spendingType.Title,
		MaximumLimit:       spendingType.MaximumLimit,
		FormatMaximumLimit: spendingType.FormatMaximumLimit,
		Icon:               spendingType.Icon,
	}
	helper.SuccessResponseEncode(w, resp, "data spending type")
}

func (p *Presenter) GetAllSpendingTypeByPeriodeAndProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
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

	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")

	spendingTypes, cursorResp, err := p.spendingTypeUsecase.GetAllByPeriodeAndProfileID(r.Context(), &usecase.RequestGetAllSpendingTypeByPeriodeWithISD{
		Periode: periodInt,
		RequestGetAllByProfileIDWithISD: usecase.RequestGetAllByProfileIDWithISD{
			ID:        cursor,
			ProfileID: profileID,
			Order:     order,
		},
	})
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	if spendingTypes == nil {
		helper.SuccessResponseEncode(w, map[string]any{
			"cursor":               "",
			"spending_type":        nil,
			"budget_amount":        0,
			"format_budget_amount": "0",
		}, "data spending type")
		return
	}

	var responseSpendingTypes []schema.ResponseSpendingTypeJoinTable
	var responseSpendingType *schema.ResponseSpendingTypeJoinTable

	for _, spendingType := range *spendingTypes.ResponseSpendingType {
		responseSpendingType = &schema.ResponseSpendingTypeJoinTable{
			ID:                 spendingType.ID,
			ProfileID:          spendingType.ProfileID,
			Title:              spendingType.Title,
			MaximumLimit:       spendingType.MaximumLimit,
			FormatMaximumLimit: spendingType.FormatMaximumLimit,
			Icon:               spendingType.Icon,
			Used:               spendingType.Used,
			FormatUsed:         spendingType.FormatUsed,
			PersentaseUsed:     spendingType.PersentaseUsed,
		}

		responseSpendingTypes = append(responseSpendingTypes, *responseSpendingType)
	}

	resp := map[string]any{
		"cursor":               cursorResp,
		"spending_type":        &responseSpendingTypes,
		"budget_amount":        spendingTypes.BudgetAmount,
		"format_budget_amount": spendingTypes.FormatBudgetAmount,
	}
	helper.SuccessResponseEncode(w, resp, "data spending types")
}

func (p *Presenter) GetAllSpendingTypeByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	_, err := uuid.Parse(profileID)
	if err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("not found", response.CM01))
		return
	}

	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")

	spendingTypes, cursorResp, err := p.spendingTypeUsecase.GetAllByProfileID(r.Context(), &usecase.RequestGetAllByProfileIDWithISD{
		ID:        cursor,
		ProfileID: profileID,
		Order:     order,
	})
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	if spendingTypes == nil {
		helper.SuccessResponseEncode(w, map[string]any{
			"cursor":        "",
			"spending_type": nil,
		}, "data spending type")
		return
	}

	var responseSpendingTypes []schema.ResponseSpendingType
	var responseSpendingType *schema.ResponseSpendingType

	for _, spendingType := range *spendingTypes {
		responseSpendingType = &schema.ResponseSpendingType{
			ID:                 spendingType.ID,
			ProfileID:          spendingType.ProfileID,
			Title:              spendingType.Title,
			MaximumLimit:       spendingType.MaximumLimit,
			FormatMaximumLimit: spendingType.FormatMaximumLimit,
			Icon:               spendingType.Icon,
		}

		responseSpendingTypes = append(responseSpendingTypes, *responseSpendingType)
	}

	resp := map[string]any{
		"cursor":        cursorResp,
		"spending_type": &responseSpendingTypes,
	}

	helper.SuccessResponseEncode(w, resp, "data spending types")
}
