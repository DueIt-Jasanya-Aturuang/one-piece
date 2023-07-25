package handlers

import (
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/services"
	"github.com/gofiber/fiber/v2"
)

type HttpHandlerImpl struct {
	PaymentService services.PaymentService
}

func NewHttpHandler(
	paymentService services.PaymentService,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		PaymentService: paymentService,
	}
}

func (hand *HttpHandlerImpl) HandlerRouter(app *fiber.App) {
}
