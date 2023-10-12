package rapi

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/schema"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (p *Presenter) CreateIncomeHistory(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")
	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	req := new(schema.RequestCreateOrUpdateIncomeHistory)

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

	incomeHistory, err := p.incomeHistoryUsecase.Create(r.Context(), &usecase.RequestCreateIncomeHistory{
		ProfileID:             profileID,
		IncomeTypeID:          req.IncomeTypeID,
		PaymentMethodID:       req.PaymentMethodID,
		PaymentName:           req.PaymentName,
		IncomeAmount:          req.IncomeAmount,
		Description:           req.Description,
		TimeIncomeHistory:     req.TimeIncomeHistory,
		ShowTimeIncomeHistory: req.ShowTimeIncomeHistory,
	})

	if err != nil {
		if errors.Is(err, usecase.IncomeHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
			log.Warn().Msgf("anomali create income history | data : %v | err : %v", req, err)
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

	resp := &schema.ResponseIncomeHistory{
		ID:                    incomeHistory.ID,
		ProfileID:             incomeHistory.ProfileID,
		IncomeTypeID:          incomeHistory.IncomeTypeID,
		IncomeTypeTitle:       incomeHistory.IncomeTypeTitle,
		PaymentMethodID:       incomeHistory.PaymentMethodID,
		PaymentMethodName:     incomeHistory.PaymentMethodName,
		PaymentName:           incomeHistory.PaymentName,
		IncomeAmount:          incomeHistory.IncomeAmount,
		FormatIncomeAmount:    incomeHistory.FormatIncomeAmount,
		Description:           incomeHistory.Description,
		TimeIncomeHistory:     incomeHistory.TimeIncomeHistory,
		ShowTimeIncomeHistory: incomeHistory.ShowTimeIncomeHistory,
	}
	helper.SuccessResponseEncode(w, resp, "create income history berhasil")
}

func (p *Presenter) UpdateIncomeHistory(w http.ResponseWriter, r *http.Request) {
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

	req := new(schema.RequestCreateOrUpdateIncomeHistory)

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

	incomeHistory, err := p.incomeHistoryUsecase.Update(r.Context(), &usecase.RequestUpdateIncomeHistory{
		ID:                    id,
		ProfileID:             profileID,
		IncomeTypeID:          req.IncomeTypeID,
		PaymentMethodID:       req.PaymentMethodID,
		PaymentName:           req.PaymentName,
		IncomeAmount:          req.IncomeAmount,
		Description:           req.Description,
		TimeIncomeHistory:     req.TimeIncomeHistory,
		ShowTimeIncomeHistory: req.ShowTimeIncomeHistory,
	})

	if err != nil {
		if errors.Is(err, usecase.IncomeHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
			log.Warn().Msgf("anomali update income history | data : %v | err : %v", req, err)
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

	resp := &schema.ResponseIncomeHistory{
		ID:                    incomeHistory.ID,
		ProfileID:             incomeHistory.ProfileID,
		IncomeTypeID:          incomeHistory.IncomeTypeID,
		IncomeTypeTitle:       incomeHistory.IncomeTypeTitle,
		PaymentMethodID:       incomeHistory.PaymentMethodID,
		PaymentMethodName:     incomeHistory.PaymentMethodName,
		PaymentName:           incomeHistory.PaymentName,
		IncomeAmount:          incomeHistory.IncomeAmount,
		FormatIncomeAmount:    incomeHistory.FormatIncomeAmount,
		Description:           incomeHistory.Description,
		TimeIncomeHistory:     incomeHistory.TimeIncomeHistory,
		ShowTimeIncomeHistory: incomeHistory.ShowTimeIncomeHistory,
	}
	helper.SuccessResponseEncode(w, resp, "update income history berhasil")
}

func (p *Presenter) DeleteIncomeHistory(w http.ResponseWriter, r *http.Request) {
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

	err := p.incomeHistoryUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.IncomeHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		if errors.Is(err, usecase.ProfileIDNotFound) {
			err = _error.HttpErrString("invalid profile id", response.CM04)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "delete spending history berhasil")
}

func (p *Presenter) GetAllIncomeHistoryByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	typeQuery := r.URL.Query().Get("type")
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")

	startTime, _ := time.Parse("2006-01-02", start)
	endTime, _ := time.Parse("2006-01-02", end)

	incomeHistories, cursorResp, err := p.incomeHistoryUsecase.GetAllByTimeAndProfileID(r.Context(), &usecase.RequestGetAllIncomeHistoryWithISD{
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

	if incomeHistories == nil {
		helper.SuccessResponseEncode(w, map[string]any{
			"cursor":         "",
			"income_history": nil,
		}, "data income history")
		return
	}

	var respIncomeHistories []schema.ResponseIncomeHistory
	var respIncomeHistory *schema.ResponseIncomeHistory

	for _, incomeHistory := range *incomeHistories {
		respIncomeHistory = &schema.ResponseIncomeHistory{
			ID:                    incomeHistory.ID,
			ProfileID:             incomeHistory.ProfileID,
			IncomeTypeID:          incomeHistory.IncomeTypeID,
			IncomeTypeTitle:       incomeHistory.IncomeTypeTitle,
			PaymentMethodID:       incomeHistory.PaymentMethodID,
			PaymentMethodName:     incomeHistory.PaymentMethodName,
			PaymentName:           incomeHistory.PaymentName,
			IncomeAmount:          incomeHistory.IncomeAmount,
			FormatIncomeAmount:    incomeHistory.FormatIncomeAmount,
			Description:           incomeHistory.Description,
			TimeIncomeHistory:     incomeHistory.TimeIncomeHistory,
			ShowTimeIncomeHistory: incomeHistory.ShowTimeIncomeHistory,
		}
		respIncomeHistories = append(respIncomeHistories, *respIncomeHistory)
	}

	resp := map[string]any{
		"cursor":         cursorResp,
		"income_history": &respIncomeHistories,
	}
	helper.SuccessResponseEncode(w, resp, "data income history")
}

func (p *Presenter) GetIncomeHistoryByIDAndProfileID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	if errParse := util.ParseUlid(id); errParse != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01))
		return
	}

	incomeHistory, err := p.incomeHistoryUsecase.GetByIDAndProfileID(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.IncomeHistoryNotFound) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseIncomeHistory{
		ID:                    incomeHistory.ID,
		ProfileID:             incomeHistory.ProfileID,
		IncomeTypeID:          incomeHistory.IncomeTypeID,
		IncomeTypeTitle:       incomeHistory.IncomeTypeTitle,
		PaymentMethodID:       incomeHistory.PaymentMethodID,
		PaymentMethodName:     incomeHistory.PaymentMethodName,
		PaymentName:           incomeHistory.PaymentName,
		IncomeAmount:          incomeHistory.IncomeAmount,
		FormatIncomeAmount:    incomeHistory.FormatIncomeAmount,
		Description:           incomeHistory.Description,
		TimeIncomeHistory:     incomeHistory.TimeIncomeHistory,
		ShowTimeIncomeHistory: incomeHistory.ShowTimeIncomeHistory,
	}
	helper.SuccessResponseEncode(w, resp, "data income history")
}
