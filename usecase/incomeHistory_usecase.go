package usecase

import (
	"context"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type IncomeHistoryUsecase interface {
	Create(ctx context.Context, req *RequestCreateIncomeHistory) (*ResponseIncomeHistory, error)
	Update(ctx context.Context, req *RequestUpdateIncomeHistory) (*ResponseIncomeHistory, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *RequestGetAllIncomeHistoryWithISD) (*[]ResponseIncomeHistory, string, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseIncomeHistory, error)
}

type ValidatePaymentIDAndIncomeTypeID struct {
	ProfileID       string
	IncomeTypeID    string
	PaymentMethodID string
}

type RequestGetAllIncomeHistoryWithISD struct {
	Type      string
	StartTime time.Time
	EndTime   time.Time
	RequestGetAllByProfileIDWithISD
}

type RequestCreateIncomeHistory struct {
	ProfileID             string
	IncomeTypeID          string `json:"income_type_id"`
	PaymentMethodID       string `json:"payment_method_id"`
	PaymentName           string `json:"payment_name"`
	IncomeAmount          int    `json:"income_amount"`
	Description           string `json:"description"`
	TimeIncomeHistory     string `json:"time_income_history"`
	ShowTimeIncomeHistory string `json:"show_time_income_history"`
}

type RequestUpdateIncomeHistory struct {
	ID                    string
	ProfileID             string
	IncomeTypeID          string `json:"income_type_id"`
	PaymentMethodID       string `json:"payment_method_id"`
	PaymentName           string `json:"payment_name"`
	IncomeAmount          int    `json:"income_amount"`
	Description           string `json:"description"`
	TimeIncomeHistory     string `json:"time_income_history"`
	ShowTimeIncomeHistory string `json:"show_time_income_history"`
}

type ResponseIncomeHistory struct {
	ID                    string    `json:"id"`
	ProfileID             string    `json:"profile_id"`
	IncomeTypeID          string    `json:"income_type_id"`
	IncomeTypeTitle       string    `json:"income_type_title"`
	PaymentMethodID       *string   `json:"payment_method_id"`
	PaymentMethodName     *string   `json:"payment_method_name"`
	PaymentName           *string   `json:"payment_name"`
	IncomeAmount          int       `json:"income_amount"`
	FormatIncomeAmount    string    `json:"format_income_amount"`
	Description           string    `json:"description"`
	TimeIncomeHistory     time.Time `json:"time_income_history"`
	ShowTimeIncomeHistory string    `json:"show_time_income_history"`
}

func (req *RequestCreateIncomeHistory) ToModel() *repository.IncomeHistory {
	id := util.NewUlid()
	timeIncomeHistory, _ := time.Parse("2006-01-02", req.TimeIncomeHistory)
	return &repository.IncomeHistory{
		ID:                    id,
		ProfileID:             req.ProfileID,
		IncomeTypeID:          req.IncomeTypeID,
		PaymentMethodID:       repository.NewNullString(req.PaymentMethodID),
		PaymentName:           repository.NewNullString(req.PaymentName),
		IncomeAmount:          req.IncomeAmount,
		Description:           req.Description,
		TimeIncomeHistory:     timeIncomeHistory.UTC(),
		ShowTimeIncomeHistory: req.ShowTimeIncomeHistory,
		AuditInfo: repository.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
}

func (req *RequestUpdateIncomeHistory) ToModel() *repository.IncomeHistory {
	timeIncomeHistory, _ := time.Parse("2006-01-02", req.TimeIncomeHistory)
	return &repository.IncomeHistory{
		ID:                    req.ID,
		ProfileID:             req.ProfileID,
		IncomeTypeID:          req.IncomeTypeID,
		PaymentMethodID:       repository.NewNullString(req.PaymentMethodID),
		PaymentName:           repository.NewNullString(req.PaymentName),
		IncomeAmount:          req.IncomeAmount,
		Description:           req.Description,
		TimeIncomeHistory:     timeIncomeHistory.UTC(),
		ShowTimeIncomeHistory: req.ShowTimeIncomeHistory,
		AuditInfo: repository.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: repository.NewNullString(req.ProfileID),
		},
	}
}
