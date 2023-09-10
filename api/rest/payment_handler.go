package rest

import (
	"net/http"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
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
}

func (h *PaymentHandlerImpl) UpdatePayment(w http.ResponseWriter, r *http.Request) {}
