package rapi

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/schema"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (p *Presenter) CreateSpendingHistory(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")
	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	req := new(schema.RequestCreateOrUpdateSpendingHistory)

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

	spendingHistory, err := p.spendingHistoryUsecase.Create(r.Context(), &usecase.RequestCreateSpendingHistory{
		ProfileID:               profileID,
		SpendingTypeID:          req.SpendingTypeID,
		PaymentMethodID:         req.PaymentMethodID,
		PaymentName:             req.PaymentName,
		SpendingAmount:          req.SpendingAmount,
		Description:             req.Description,
		TimeSpendingHistory:     req.TimeSpendingHistory,
		ShowTimeSpendingHistory: req.ShowTimeSpendingHistory,
	})
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

	resp := &schema.ResponseSpendingHistory{
		ID:                      spendingHistory.ID,
		ProfileID:               spendingHistory.ProfileID,
		SpendingTypeID:          spendingHistory.SpendingTypeID,
		SpendingTypeTitle:       spendingHistory.SpendingTypeTitle,
		PaymentMethodID:         spendingHistory.PaymentMethodID,
		PaymentMethodName:       spendingHistory.PaymentMethodName,
		PaymentName:             spendingHistory.PaymentName,
		BeforeBalance:           spendingHistory.BeforeBalance,
		SpendingAmount:          spendingHistory.SpendingAmount,
		FormatSpendingAmount:    spendingHistory.FormatSpendingAmount,
		AfterBalance:            spendingHistory.AfterBalance,
		Description:             spendingHistory.Description,
		TimeSpendingHistory:     spendingHistory.TimeSpendingHistory,
		ShowTimeSpendingHistory: spendingHistory.ShowTimeSpendingHistory,
	}
	helper.SuccessResponseEncode(w, resp, "create spending history berhasil")
}

func (p *Presenter) UpdateSpendingHistory(w http.ResponseWriter, r *http.Request) {
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

	req := new(schema.RequestCreateOrUpdateSpendingHistory)

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

	spendingHistory, err := p.spendingHistoryUsecase.Update(r.Context(), &usecase.RequestUpdateSpendingHistory{
		ID:                      id,
		ProfileID:               profileID,
		SpendingTypeID:          req.SpendingTypeID,
		PaymentMethodID:         req.PaymentMethodID,
		PaymentName:             req.PaymentName,
		SpendingAmount:          req.SpendingAmount,
		Description:             req.Description,
		TimeSpendingHistory:     req.TimeSpendingHistory,
		ShowTimeSpendingHistory: req.ShowTimeSpendingHistory,
	})
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

	resp := &schema.ResponseSpendingHistory{
		ID:                      spendingHistory.ID,
		ProfileID:               spendingHistory.ProfileID,
		SpendingTypeID:          spendingHistory.SpendingTypeID,
		SpendingTypeTitle:       spendingHistory.SpendingTypeTitle,
		PaymentMethodID:         spendingHistory.PaymentMethodID,
		PaymentMethodName:       spendingHistory.PaymentMethodName,
		PaymentName:             spendingHistory.PaymentName,
		BeforeBalance:           spendingHistory.BeforeBalance,
		SpendingAmount:          spendingHistory.SpendingAmount,
		FormatSpendingAmount:    spendingHistory.FormatSpendingAmount,
		AfterBalance:            spendingHistory.AfterBalance,
		Description:             spendingHistory.Description,
		TimeSpendingHistory:     spendingHistory.TimeSpendingHistory,
		ShowTimeSpendingHistory: spendingHistory.ShowTimeSpendingHistory,
	}
	helper.SuccessResponseEncode(w, resp, "update spending history berhasil")
}

func (p *Presenter) DeleteSpendingHistory(w http.ResponseWriter, r *http.Request) {
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

	err := p.spendingHistoryUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.SpendingHistoryNotFound) {
			helper.SuccessResponseEncode(w, nil, "delete spending history berhasil")
			return
		}
		if errors.Is(err, usecase.ProfileIDNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "delete spending history berhasil")
}

func (p *Presenter) GetAllSpendingHistoryByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	typeQuery := r.URL.Query().Get("type")
	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")
	startTime, _ := time.Parse("2006-01-02", r.URL.Query().Get("start"))
	endTime, _ := time.Parse("2006-01-02", r.URL.Query().Get("end"))

	spendingHistories, cursorResp, err := p.spendingHistoryUsecase.GetAllByTimeAndProfileID(r.Context(), &usecase.RequestGetAllSpendingHistoryWithISD{
		Type:      typeQuery,
		StartTime: startTime,
		EndTime:   endTime,
		RequestGetAllByProfileIDWithISD: usecase.RequestGetAllByProfileIDWithISD{
			ID:        cursor,
			ProfileID: profileID,
			Order:     order,
		},
	})
	if err != nil {
		if errors.Is(err, usecase.InvalidTimestamp) {
			err = _error.HttpErrString(err.Error(), response.CM06)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	if spendingHistories == nil {
		helper.SuccessResponseEncode(w, map[string]any{
			"cursor":           "",
			"spending_history": nil,
		}, "data spending history")
		return
	}

	var responses []schema.ResponseSpendingHistory
	var responseSpendingHistory *schema.ResponseSpendingHistory

	for _, spendingHistory := range *spendingHistories {
		responseSpendingHistory = &schema.ResponseSpendingHistory{
			ID:                      spendingHistory.ID,
			ProfileID:               spendingHistory.ProfileID,
			SpendingTypeID:          spendingHistory.SpendingTypeID,
			SpendingTypeTitle:       spendingHistory.SpendingTypeTitle,
			PaymentMethodID:         spendingHistory.PaymentMethodID,
			PaymentMethodName:       spendingHistory.PaymentMethodName,
			PaymentName:             spendingHistory.PaymentName,
			BeforeBalance:           spendingHistory.BeforeBalance,
			SpendingAmount:          spendingHistory.SpendingAmount,
			FormatSpendingAmount:    spendingHistory.FormatSpendingAmount,
			AfterBalance:            spendingHistory.AfterBalance,
			Description:             spendingHistory.Description,
			TimeSpendingHistory:     spendingHistory.TimeSpendingHistory,
			ShowTimeSpendingHistory: spendingHistory.ShowTimeSpendingHistory,
		}

		responses = append(responses, *responseSpendingHistory)
	}

	resp := map[string]any{
		"cursor":           cursorResp,
		"spending_history": &responses,
	}
	helper.SuccessResponseEncode(w, resp, "data spending history")
}

func (p *Presenter) GetSpendingHistoryByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
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

	spendingHistory, err := p.spendingHistoryUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.SpendingHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseSpendingHistory{
		ID:                      spendingHistory.ID,
		ProfileID:               spendingHistory.ProfileID,
		SpendingTypeID:          spendingHistory.SpendingTypeID,
		SpendingTypeTitle:       spendingHistory.SpendingTypeTitle,
		PaymentMethodID:         spendingHistory.PaymentMethodID,
		PaymentMethodName:       spendingHistory.PaymentMethodName,
		PaymentName:             spendingHistory.PaymentName,
		BeforeBalance:           spendingHistory.BeforeBalance,
		SpendingAmount:          spendingHistory.SpendingAmount,
		FormatSpendingAmount:    spendingHistory.FormatSpendingAmount,
		AfterBalance:            spendingHistory.AfterBalance,
		Description:             spendingHistory.Description,
		TimeSpendingHistory:     spendingHistory.TimeSpendingHistory,
		ShowTimeSpendingHistory: spendingHistory.ShowTimeSpendingHistory,
	}

	helper.SuccessResponseEncode(w, resp, "data spending history")
}
