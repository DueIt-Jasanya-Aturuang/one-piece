package handlers

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/src/modules/services"
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

func (handler *HttpHandlerImpl) HandlerRouter(app *fiber.App) {
	app.Route("/payment", func(router fiber.Router) {
		router.Post("/", handler.CreatePayment)
		router.Put("/:id", handler.UpdatePayment)
	})
}
