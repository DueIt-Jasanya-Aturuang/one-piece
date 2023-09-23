package domain

import (
	"context"
	"database/sql"
	"time"
)

// SpendingHistory spending history entity
type SpendingHistory struct {
	ID                      string
	ProfileID               string
	SpendingTypeID          string
	PaymentMethodID         sql.NullString
	PaymentName             sql.NullString
	BeforeBalance           int
	SpendingAmount          int
	AfterBalance            int
	Description             string
	TimeSpendingHistory     time.Time
	ShowTimeSpendingHistory string
	AuditInfo
}

// SpendingHistoryJoin query inner join
type SpendingHistoryJoin struct {
	ID                      string
	ProfileID               string
	SpendingTypeID          string
	SpendingTypeTitle       string
	PaymentMethodID         sql.NullString
	PaymentMethodName       sql.NullString
	PaymentName             sql.NullString
	BeforeBalance           int
	SpendingAmount          int
	AfterBalance            int
	Description             string
	TimeSpendingHistory     time.Time
	ShowTimeSpendingHistory string
	AuditInfo
}

// RequestCreateSpendingHistory request create spending history
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

// RequestUpdateSpendingHistory request update spending history
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

// RequestGetFilteredDataSpendingHistory request get filtered data spending history
type RequestGetFilteredDataSpendingHistory struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
	Type      string
}

// RequestValidatePaymentAndSpendingTypeID untuk validasi
type RequestValidatePaymentAndSpendingTypeID struct {
	ProfileID       string
	SpendingTypeID  string
	PaymentMethodID string
}

// ResponseSpendingHistory response spending history
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

// SpendingHistoryRepository spending history repository interface
type SpendingHistoryRepository interface {
	Create(ctx context.Context, spendingHistory *SpendingHistory) error
	Update(ctx context.Context, spendingHistory *SpendingHistory) error
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *RequestGetFilteredDataSpendingHistory) (*[]SpendingHistoryJoin, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*SpendingHistoryJoin, error)
	UnitOfWorkRepository
}

// SpendingHistoryUsecase spending history usecase interface
type SpendingHistoryUsecase interface {
	Create(ctx context.Context, req *RequestCreateSpendingHistory) (*ResponseSpendingHistory, error)
	Update(ctx context.Context, req *RequestUpdateSpendingHistory) (*ResponseSpendingHistory, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *RequestGetFilteredDataSpendingHistory) (*[]ResponseSpendingHistory, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseSpendingHistory, error)
}
