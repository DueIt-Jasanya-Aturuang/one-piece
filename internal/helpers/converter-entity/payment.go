package converterentity

import (
	"time"

	sqlHelper "github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/internal/utils/sql-helper"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/dto"
	"github.com/DueIt-Jasanya-Aturuang/DueIt-Payment-Service/src/modules/entities"
)

func (convert *ConvertImpl) CreatePaymentToEntity(id string) entities.Payment {
	return entities.Payment{
		Id:          id,
		Name:        convert.PaymentCreateRequest.Name,
		Description: sqlHelper.NewNullString(convert.PaymentCreateRequest.Description),
		Image:       convert.PaymentCreateRequest.Image.Filename,
		CreatedAt:   time.Now().Unix(),
		CreatedBy:   id,
		UpdatedAt:   time.Now().Unix(),
	}
}

func (convert *ConvertImpl) UpdatePaymentToEntity() entities.Payment {
	return entities.Payment{
		Name:        convert.PaymentUpdateRequest.Name,
		Description: sqlHelper.NewNullString(convert.PaymentUpdateRequest.Description),
		Image:       convert.PaymentEntity.Image,
		UpdatedAt:   time.Now().Unix(),
		UpdatedBy:   sqlHelper.NewNullString(convert.PaymentEntity.Id),
	}
}

func (convert *ConvertImpl) PaymentEntityToResponse() *dto.PaymentResponse {
	return &dto.PaymentResponse{
		Id:          convert.PaymentEntity.Id,
		Name:        convert.PaymentEntity.Name,
		Description: sqlHelper.GetNullString(convert.PaymentEntity.Description),
		Image:       convert.PaymentEntity.Image,
	}
}
