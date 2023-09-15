package converter

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/helper"
)

func CreatePaymentReqToModel(req *domain.RequestCreatePayment, fileName string) *domain.Payment {
	id := uuid.NewV4().String()
	payment := &domain.Payment{
		ID:          id,
		Name:        req.Name,
		Description: helper.NewNullString(req.Description),
		Image:       fileName,
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: "",
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{},
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		},
	}

	return payment
}

func PaymentModelToResp(payment *domain.Payment) *domain.ResponsePayment {
	resp := &domain.ResponsePayment{
		ID:          payment.ID,
		Name:        payment.Name,
		Description: helper.GetNullString(payment.Description),
		Image:       payment.Image,
	}

	return resp
}

func PaymentGetAllPaymentModelToResp(payment domain.Payment) domain.ResponsePayment {
	resp := domain.ResponsePayment{
		ID:          payment.ID,
		Name:        payment.Name,
		Description: helper.GetNullString(payment.Description),
		Image:       payment.Image,
	}

	return resp
}

func UpdatePaymentReqToModel(req *domain.RequestUpdatePayment, fileName string) *domain.Payment {
	payment := &domain.Payment{
		ID:          req.ID,
		Name:        req.Name,
		Description: helper.NewNullString(req.Description),
		Image:       fileName,
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper.NewNullString(""),
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		},
	}

	return payment
}
