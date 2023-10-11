package rapi

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"
	"github.com/oklog/ulid/v2"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/validation"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase_old"
)

type PaymentHandlerImpl struct {
	paymentUsecase domain.PaymentUsecase
}

func NewPaymentHandlerImpl(
	paymentUsecase domain.PaymentUsecase,
) *PaymentHandlerImpl {
	return &PaymentHandlerImpl{
		paymentUsecase: paymentUsecase,
	}
}

func (h *PaymentHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreatePayment)

	err := helper.ParseForm(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	_, fileHeader, _ := r.FormFile("image")
	profileID := r.Header.Get("Profile-ID")

	req.ProfileID = profileID
	req.Image = fileHeader

	err = validation.CreatePayment(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	payment, err := h.paymentUsecase.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, usecase_old.NamePaymentExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"name": {
					err.Error(),
				},
			}, response.CM06)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, payment, "created payment berhasil")
}

func (h *PaymentHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestUpdatePayment)

	err := helper.ParseForm(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")
	req.ProfileID = profileID
	req.ID = id

	_, fileHeader, _ := r.FormFile("image")
	req.Image = fileHeader

	err = validation.UpdatePayment(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	payment, err := h.paymentUsecase.Update(r.Context(), req)
	if err != nil {
		if errors.Is(err, usecase_old.NamePaymentExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"name": {
					err.Error(),
				},
			}, response.CM06)
		} else if errors.Is(err, usecase_old.PaymentNotExist) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, payment, "update payment berhasil")
}

func (h *PaymentHandlerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")
	cursor := r.URL.Query().Get("cursor")
	order := r.URL.Query().Get("order")
	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM05))
		return
	}

	order, operation := helper.GetOrder(order)

	req := &domain.RequestGetAllPaginate{
		ProfileID: profileID,
		ID:        cursor,
		Operation: operation,
		Order:     order,
	}
	payments, cursorResp, err := h.paymentUsecase.GetAllByProfileID(r.Context(), req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := map[string]any{
		"payments": payments,
		"cursor":   cursorResp,
	}
	helper.SuccessResponseEncode(w, resp, "data payment")
}

func (h *PaymentHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString("invalid profile id", response.CM05))
		return
	}

	if _, err := ulid.Parse(id); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01))
		return
	}

	err := h.paymentUsecase.Delete(r.Context(), id, profileID)
	if err != nil {
		if errors.Is(err, usecase_old.PaymentNotExist) {
			helper.SuccessResponseEncode(w, nil, "deleted payment successfully")
			return
		}
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, nil, "deleted payment successfully")
}
