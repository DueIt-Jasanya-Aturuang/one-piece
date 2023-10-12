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

func (p *Presenter) CreatePayment(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	req := new(schema.RequestCreateOrUpdatePayment)

	err := helper.ParseForm(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	_, fileHeader, _ := r.FormFile("image")
	req.Image = fileHeader

	err = req.ValidateCreate()
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	payment, err := p.paymentUsecase.Create(r.Context(), &usecase.RequestCreatePayment{
		ProfileID:   profileID,
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
	})
	if err != nil {
		if errors.Is(err, usecase.NamePaymentExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"name": {
					err.Error(),
				},
			}, response.CM06)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponsePayment{
		ID:          payment.ID,
		ProfileID:   payment.ProfileID,
		Name:        payment.Name,
		Description: payment.Description,
		Image:       payment.Image,
	}
	helper.SuccessResponseEncode(w, resp, "created payment berhasil")
}

func (p *Presenter) UpdatePayment(w http.ResponseWriter, r *http.Request) {
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

	req := new(schema.RequestCreateOrUpdatePayment)

	err := helper.ParseForm(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	_, fileHeader, _ := r.FormFile("image")
	req.Image = fileHeader

	err = req.ValidateUpdate()
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	payment, err := p.paymentUsecase.Update(r.Context(), &usecase.RequestUpdatePayment{
		ProfileID:   profileID,
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
	})
	if err != nil {
		if errors.Is(err, usecase.NamePaymentExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"name": {
					err.Error(),
				},
			}, response.CM06)
		} else if errors.Is(err, usecase.PaymentNotExist) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponsePayment{
		ID:          payment.ID,
		ProfileID:   payment.ProfileID,
		Name:        payment.Name,
		Description: payment.Description,
		Image:       payment.Image,
	}
	helper.SuccessResponseEncode(w, resp, "update payment berhasil")
}

func (p *Presenter) GetAllPayment(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM04))
		return
	}

	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")

	payments, cursorResp, err := p.paymentUsecase.GetAllByProfileID(r.Context(), &usecase.RequestGetAllByProfileIDWithISD{
		ID:        cursor,
		ProfileID: profileID,
		Order:     order,
	})
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	var responses []schema.ResponsePayment
	var responsePayment *schema.ResponsePayment

	for _, payment := range *payments {
		responsePayment = &schema.ResponsePayment{
			ID:          payment.ID,
			ProfileID:   payment.ProfileID,
			Name:        payment.Name,
			Description: payment.Description,
			Image:       payment.Image,
		}

		responses = append(responses, *responsePayment)
	}

	resp := map[string]any{
		"payments": &responses,
		"cursor":   cursorResp,
	}
	helper.SuccessResponseEncode(w, resp, "data payment")
}

func (p *Presenter) DeletePayment(w http.ResponseWriter, r *http.Request) {
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

	err := p.paymentUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase.PaymentNotExist) {
			helper.SuccessResponseEncode(w, nil, "deleted payment successfully")
			return
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "deleted payment successfully")
}
