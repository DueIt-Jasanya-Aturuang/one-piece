package domain

import (
	"context"
	"database/sql"
	"time"
)

type IncomeHistoryRepository interface {
	Create(ctx context.Context, income *IncomeHistory) error
	Update(ctx context.Context, income *IncomeHistory) error
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *GetIncomeHistoryByTimeAndProfileID) (*[]IncomeHistoryJoin, error)
	GetTotalIncomeByPeriode(ctx context.Context, req *GetIncomeHistoryByTimeAndProfileID) (int, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*IncomeHistoryJoin, error)
	UnitOfWorkRepository
}

type IncomeHistoryUsecase interface {
	Create(ctx context.Context, req *RequestCreateIncomeHistory) (*ResponseIncomeHistory, error)
	Update(ctx context.Context, req *RequestUpdateIncomeHistory) (*ResponseIncomeHistory, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *GetFilteredDataIncomeHistory) (*[]ResponseIncomeHistory, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseIncomeHistory, error)
}

// GetFilteredDataIncomeHistory for usecase layer param
type GetFilteredDataIncomeHistory struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
	Type      string
}

type ValidatePaymentAndIncomeTypeID struct {
	ProfileID       string
	SpendingTypeID  string
	PaymentMethodID string
}

// GetIncomeHistoryByTimeAndProfileID for repo layer param
type GetIncomeHistoryByTimeAndProfileID struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
}

// IncomeHistory model
type IncomeHistory struct {
	ID                    string
	ProfileID             string
	IncomeTypeID          string
	PaymentMethodID       sql.NullString
	PaymentName           sql.NullString
	IncomeAmount          int
	Description           string
	TimeIncomeHistory     time.Time
	ShowTimeIncomeHistory string
	AuditInfo
}

// IncomeHistoryJoin model join table
type IncomeHistoryJoin struct {
	ID                    string
	ProfileID             string
	IncomeTypeID          string
	IncomeTypeTitle       string
	PaymentMethodID       sql.NullString
	PaymentMethodName     sql.NullString
	PaymentName           sql.NullString
	IncomeAmount          int
	Description           string
	TimeIncomeHistory     time.Time
	ShowTimeIncomeHistory string
	AuditInfo
}

// RequestCreateIncomeHistory request schema api
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

// ResponseIncomeHistory response schema api
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
