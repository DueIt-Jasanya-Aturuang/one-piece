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
	SpendingTypeID          string
	PaymentMethodID         string
	PaymentName             string
	SpendingAmount          int
	Description             string
	TimeSpendingHistory     string
	ShowTimeSpendingHistory string
}

type RequestUpdateSpendingHistory struct {
	ID                      string
	ProfileID               string
	SpendingTypeID          string
	PaymentMethodID         string
	PaymentName             string
	SpendingAmount          int
	Description             string
	TimeSpendingHistory     string
	ShowTimeSpendingHistory string
}

type ResponseSpendingHistory struct {
	ID                      string
	ProfileID               string
	SpendingTypeID          string
	SpendingTypeTitle       string
	PaymentMethodID         *string
	PaymentMethodName       *string
	PaymentName             *string
	BeforeBalance           int
	SpendingAmount          int
	FormatSpendingAmount    string
	AfterBalance            int
	Description             string
	TimeSpendingHistory     time.Time
	ShowTimeSpendingHistory string
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
