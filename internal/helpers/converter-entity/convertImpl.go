package converterentity

import (
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/dto"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/entities"
)

type ConvertImpl struct {
	PaymentEntity        *entities.Payment
	PaymentCreateRequest *dto.PaymentCreateRequest
	PaymentUpdateRequest *dto.PaymentUpdateRequest
}

func NewConvertImpl() *ConvertImpl {
	return &ConvertImpl{}
}
