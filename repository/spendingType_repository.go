package repository

import (
	"context"
	"time"
)

type SpendingTypeRepository interface {
	Create(ctx context.Context, spendingType *SpendingType) error
	Update(ctx context.Context, spendingType *SpendingType) error
	Delete(ctx context.Context, id string, profileID string) error
	CheckData(ctx context.Context, profileID string) (bool, error)
	CheckByTitleAndProfileID(ctx context.Context, profileID string, title string) (bool, error)
	GetDefault(ctx context.Context) (*[]SpendingType, error)
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*SpendingType, error)
	GetAllByTimeAndProfileID(ctx context.Context, req *GetAllSpendingTypeByFilterWithISD) (*[]SpendingTypeJoinTable, error)
	GetAllByProfileID(ctx context.Context, req *GetAllSpendingTypeWithISD) (*[]SpendingType, error)
	UnitOfWorkRepository
}

type SpendingType struct {
	ID           string
	ProfileID    string
	Title        string
	MaximumLimit int
	Icon         string
	AuditInfo
}

type SpendingTypeJoinTable struct {
	ID           string
	ProfileID    string
	Title        string
	MaximumLimit int
	Icon         string
	Used         int
	AuditInfo
}

type GetAllSpendingTypeByFilterWithISD struct {
	ProfileID string
	StartTime time.Time
	EndTime   time.Time
	InfiniteScrollData
}

type GetAllSpendingTypeWithISD struct {
	ProfileID string
	InfiniteScrollData
}
