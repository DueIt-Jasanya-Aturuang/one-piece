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
	req.Image = fileHeader

	err = validation.CreatePayment(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	payment, err := h.paymentUsecase.Create(r.Context(), req)
	if err != nil {
		if errors.Is(err, _usecase.NamePaymentExist) {
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
	if id == "" {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01))
		return
	}
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
		if errors.Is(err, _usecase.NamePaymentExist) {
			err = _error.HttpErrMapOfSlices(map[string][]string{
				"name": {
					err.Error(),
				},
			}, response.CM06)
		} else if errors.Is(err, _usecase.PaymentNotExist) {
			err = _error.HttpErrString(err.Error(), response.CM01)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, payment, "update payment berhasil")
}

func (h *PaymentHandlerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	payments, err := h.paymentUsecase.GetAll(r.Context())
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, payments, "data payment")
}
