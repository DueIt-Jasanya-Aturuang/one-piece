package usecase

import (
	"context"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type SpendingHistoryUsecase interface {
	Create(ctx context.Context, req *RequestCreateSpendingHistory) (*ResponseSpendingHistory, error)
	Update(ctx context.Context, req *RequestUpdateSpendingHistory) (*ResponseSpendingHistory, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *RequestGetAllSpendingHistoryWithISD) (*[]ResponseSpendingHistory, string, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseSpendingHistory, error)
}

type ValidatePaymentIDAndSpendingTypeID struct {
	ProfileID       string
	SpendingTypeID  string
	PaymentMethodID string
}

type RequestGetAllSpendingHistoryWithISD struct {
	Type      string
	StartTime time.Time
	EndTime   time.Time
	RequestGetAllByProfileIDWithISD
}

type RequestCreateSpendingHistory struct {
	ProfileID               string
	SpendingTypeID          string `json:"spending_type_id"`
	PaymentMethodID         string `json:"payment_method_id"`
	PaymentName             string `json:"payment_name"`
	SpendingAmount          int    `json:"spending_amount"`
	Description             string `json:"description"`
	TimeSpendingHistory     string `json:"time_spending_history"`
	ShowTimeSpendingHistory string `json:"show_time_spending_history"`
}

type RequestUpdateSpendingHistory struct {
	ID                      string
	ProfileID               string
	SpendingTypeID          string `json:"spending_type_id"`
	PaymentMethodID         string `json:"payment_method_id"`
	PaymentName             string `json:"payment_name"`
	SpendingAmount          int    `json:"spending_amount"`
	Description             string `json:"description"`
	TimeSpendingHistory     string `json:"time_spending_history"`
	ShowTimeSpendingHistory string `json:"show_time_spending_history"`
}

type ResponseSpendingHistory struct {
	ID                      string    `json:"id"`
	ProfileID               string    `json:"profile_id"`
	SpendingTypeID          string    `json:"spending_type_id"`
	SpendingTypeTitle       string    `json:"spending_type_title"`
	PaymentMethodID         *string   `json:"payment_method_id"`
	PaymentMethodName       *string   `json:"payment_method_name"`
	PaymentName             *string   `json:"payment_name"`
	BeforeBalance           int       `json:"before_balance"`
	SpendingAmount          int       `json:"spending_amount"`
	FormatSpendingAmount    string    `json:"format_spending_amount"`
	AfterBalance            int       `json:"after_balance"`
	Description             string    `json:"description"`
	TimeSpendingHistory     time.Time `json:"time_spending_history"`
	ShowTimeSpendingHistory string    `json:"show_time_spending_history"`
}

func (req *RequestCreateSpendingHistory) ToModel(balance int) *repository.SpendingHistory {
	id := util.NewUlid()
	timeSpendingHistory, _ := time.Parse("2006-01-02", req.TimeSpendingHistory)
	spendingHistory := &repository.SpendingHistory{
		ID:                      id,
		ProfileID:               req.ProfileID,
		SpendingTypeID:          req.SpendingTypeID,
		PaymentMethodID:         repository.NewNullString(req.PaymentMethodID),
		PaymentName:             repository.NewNullString(req.PaymentName),
		BeforeBalance:           balance,
		SpendingAmount:          req.SpendingAmount,
		AfterBalance:            balance - req.SpendingAmount,
		Description:             req.Description,
		TimeSpendingHistory:     timeSpendingHistory.UTC(),
		ShowTimeSpendingHistory: req.ShowTimeSpendingHistory,
		AuditInfo: repository.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
	return spendingHistory
}

func (req *RequestUpdateSpendingHistory) ToModel(balance int) *repository.SpendingHistory {
	timeSpendingHistory, _ := time.Parse("2006-01-02", req.TimeSpendingHistory)
	spendingHistory := &repository.SpendingHistory{
		ID:                      req.ID,
		ProfileID:               req.ProfileID,
		SpendingTypeID:          req.SpendingTypeID,
		PaymentMethodID:         repository.NewNullString(req.PaymentMethodID),
		PaymentName:             repository.NewNullString(req.PaymentName),
		BeforeBalance:           balance,
		SpendingAmount:          req.SpendingAmount,
		AfterBalance:            balance - req.SpendingAmount,
		Description:             req.Description,
		TimeSpendingHistory:     timeSpendingHistory.UTC(),
		ShowTimeSpendingHistory: req.ShowTimeSpendingHistory,
		AuditInfo: repository.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: repository.NewNullString(req.ProfileID),
		},
	}
	return spendingHistory
}
