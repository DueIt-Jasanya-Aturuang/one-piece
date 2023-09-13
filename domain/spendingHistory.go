package domain

import (
	"context"
	"database/sql"
)

// SpendingHistory spending history entity
type SpendingHistory struct {
	ID                  string
	ProfileID           string
	SpendingTypeID      string
	PaymentMethodID     sql.NullString
	PaymentName         sql.NullString
	BeforeBalance       int
	SpendingAmount      int
	AfterBalance        int
	Description         string
	TimeSpendingHistory string
	CreatedAt           int64
	CreatedBy           string
	UpdatedAt           int64
	UpdatedBy           sql.NullString
	DeletedAt           sql.NullInt64
	DeletedBy           sql.NullString
}

// RequestCreateSpendingHistory request create spending history
type RequestCreateSpendingHistory struct {
	ProfileID           string `json:"profile_id"`
	SpendingTypeID      string `json:"spending_type_id"`
	PaymentMethodID     string `json:"payment_method_id"`
	PaymentName         string `json:"payment_name"`
	SpendingAmount      int    `json:"spending_amount"`
	Description         string `json:"description"`
	TimeSpendingHistory string `json:"time_spending_history"`
}

// RequestUpdateSpendingHistory request update spending history
type RequestUpdateSpendingHistory struct {
	ID                  string
	ProfileID           string
	SpendingTypeID      string `json:"spending_type_id"`
	PaymentMethodID     string `json:"payment_method_id"`
	PaymentName         string `json:"payment_name"`
	SpendingAmount      int    `json:"spending_amount"`
	Description         string `json:"description"`
	TimeSpendingHistory string `json:"time_spending_history"`
}

// ResponseSpendingHistory response spending history
type ResponseSpendingHistory struct {
	ID                  string  `json:"id"`
	ProfileID           string  `json:"profile_id"`
	SpendingTypeID      string  `json:"spending_type_id"`
	PaymentMethodID     *string `json:"payment_method_id"`
	PaymentName         *string `json:"payment_name"`
	BeforeBalance       int     `json:"before_balance"`
	SpendingAmount      int     `json:"spending_amount"`
	AfterBalance        int     `json:"after_balance"`
	Description         string  `json:"description"`
	TimeSpendingHistory string  `json:"time_spending_history"`
}

// SpendingHistoryRepository spending history repository interface
type SpendingHistoryRepository interface {
	Create(ctx context.Context, spendingHistory *SpendingHistory) error
	Update(ctx context.Context, spendingHistory *SpendingHistory) error
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByProfileID(ctx context.Context, profileID string) (*[]SpendingHistory, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*SpendingHistory, error)
	UnitOfWorkRepository
}

// SpendingHistoryUsecase spending history usecase interface
type SpendingHistoryUsecase interface {
	Create(ctx context.Context, req *RequestCreateSpendingHistory) (*ResponseSpendingHistory, error)
	Update(ctx context.Context, req *RequestUpdateSpendingHistory) (*ResponseSpendingHistory, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByProfileID(ctx context.Context, profileID string) (*[]ResponseSpendingHistory, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseSpendingHistory, error)
}
