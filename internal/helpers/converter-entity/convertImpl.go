package converterentity

import (
	"github.com/DueIt-Jasanya-Aturuang/one-piece/src/modules/dto"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/src/modules/entities"
)

type ConvertImpl struct {
	PaymentEntity        *entities.Payment
	PaymentCreateRequest *dto.PaymentCreateRequest
	PaymentUpdateRequest *dto.PaymentUpdateRequest
}

func NewConvertImpl() *ConvertImpl {
	return &ConvertImpl{}
}
