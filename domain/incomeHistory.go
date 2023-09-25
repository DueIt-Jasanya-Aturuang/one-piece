package domain

import (
	"context"
	"time"
)

type IncomeHistory struct {
	ID string
	AuditInfo
}

type IncomeHistoryJoin struct {
	ID string
	AuditInfo
}

type RequestCreateIncomeHistory struct {
	Name string `json:"name"`
}

type RequestUpdateIncomeHistory struct {
	ProfileID string
	ID        string `json:"id"`
	Name      string `json:"name"`
}

type RequestGetFilteredDataIncomeHistory struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
	Type      string
}

type ResponseIncomeHistory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IncomeHistoryRepository interface {
	Create(ctx context.Context, income *IncomeHistory) error
	Update(ctx context.Context, income *IncomeHistory) error
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *RequestGetFilteredDataIncomeHistory) (*[]IncomeHistoryJoin, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*IncomeHistoryJoin, error)
	UnitOfWorkRepository
}

type IncomeHistoryUsecase interface {
	Create(ctx context.Context, req *RequestCreateIncomeHistory) (*ResponseIncomeHistory, error)
	Update(ctx context.Context, req *RequestUpdateIncomeHistory) (*ResponseIncomeHistory, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetAllByTimeAndProfileID(ctx context.Context, req *RequestGetFilteredDataIncomeHistory) (*[]ResponseIncomeHistory, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseIncomeHistory, error)
}
