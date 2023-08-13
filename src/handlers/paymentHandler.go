package handlers

import (
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/http-protocol/exception"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (h *HttpHandlerImpl) CreatePayment(c *fiber.Ctx) error {
	payload := new(dto.PaymentCreateRequest)
	if err := c.BodyParser(payload); err != nil {
		log.Warn().Msgf("cannnot parser | %v", err.Error())
		return exception.Err(c, err)
	}
	image, _ := c.FormFile("image")
	if image != nil {
		payload.Image = image
	}

	payment, err := h.PaymentService.CreatePayment(c.Context(), payload)
	if err != nil {
		return exception.Err(c, err)
	}

	return c.Status(200).JSON(exception.ResponseSuccess{
		Data: map[string]any{
			"payment": payment,
		},
	})
}

func (h *HttpHandlerImpl) UpdatePayment(c *fiber.Ctx) error {
	paymentId := c.Get("id", "")
	payload := new(dto.PaymentUpdateRequest)
	if err := c.BodyParser(payload); err != nil {
		log.Warn().Msgf("cannnot parser | %v", err.Error())
		return exception.Err(c, err)
	}

	image, _ := c.FormFile("image")
	if image != nil {
		payload.Image = image
	}

	payment, err := h.PaymentService.UpdatePayment(c.Context(), payload, paymentId)
	if err != nil {
		return exception.Err(c, err)
	}

	return c.Status(200).JSON(exception.ResponseSuccess{
		Data: map[string]any{
			"payment": payment,
		},
	})
}
