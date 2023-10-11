package usecase

import (
	"context"
	"database/sql"
	"mime/multipart"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type PaymentUsecase interface {
	Create(ctx context.Context, req *RequestCreatePayment) (*ResponsePayment, error)
	Update(ctx context.Context, req *RequestUpdatePayment) (*ResponsePayment, error)
	GetAllByProfileID(ctx context.Context, req *RequestGetAllByProfileIDWithISD) (*[]ResponsePayment, string, error)
	Delete(ctx context.Context, id string, profileID string) error
}

type RequestCreatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
	ProfileID   string
}

type RequestUpdatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
	ProfileID   string
	ID          string
}

type ResponsePayment struct {
	ID          string  `json:"id"`
	ProfileID   string  `json:"profile_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Image       string  `json:"image"`
}

func (req *RequestCreatePayment) ToModel(fileName string) *repository.Payment {
	id := util.NewUlid()
	payment := &repository.Payment{
		ID:          id,
		ProfileID:   req.ProfileID,
		Name:        req.Name,
		Description: repository.NewNullString(req.Description),
		Image:       fileName,
		AuditInfo: repository.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}

	return payment
}

func (req *RequestUpdatePayment) ToModel(fileName string) *repository.Payment {
	payment := &repository.Payment{
		ID:          req.ID,
		ProfileID:   req.ProfileID,
		Name:        req.Name,
		Description: repository.NewNullString(req.Description),
		Image:       fileName,
		AuditInfo: repository.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: repository.NewNullString(req.ProfileID),
			DeletedAt: sql.NullInt64{},
			DeletedBy: sql.NullString{},
		},
	}

	return payment
}
