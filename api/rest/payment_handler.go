package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/validation"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
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

func (h *PaymentHandlerImpl) CreatePayment(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestCreatePayment)

	err := helper.ParseForm(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	_, fileHeader, _ := r.FormFile("image")
	req.Image = fileHeader

	err = validation.CreatePaymentValidation(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	payment, err := h.paymentUsecase.CreatePayment(r.Context(), req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := domain.ResponseSuccessHTTP{
		Data: payment,
		Code: 201,
	}

	helper.SuccessResponseEncode(w, resp)
}

func (h *PaymentHandlerImpl) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	req := new(domain.RequestUpdatePayment)

	err := helper.ParseForm(r, req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		helper.ErrorResponseEncode(w, util.ErrHTTPString("not found", 404))
		return
	}
	req.ID = id

	_, fileHeader, _ := r.FormFile("image")
	req.Image = fileHeader
	
	err = validation.UpdatePaymentValidation(req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	payment, err := h.paymentUsecase.UpdatePayment(r.Context(), req)
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := domain.ResponseSuccessHTTP{
		Data: payment,
		Code: 200,
	}

	helper.SuccessResponseEncode(w, resp)
}

func (h *PaymentHandlerImpl) GetAllPayment(w http.ResponseWriter, r *http.Request) {
	payments, err := h.paymentUsecase.GetAllPayment(r.Context())
	if err != nil {
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := domain.ResponseSuccessHTTP{
		Data: payments,
		Code: 200,
	}

	helper.SuccessResponseEncode(w, resp)
}
