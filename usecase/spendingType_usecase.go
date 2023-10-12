package usecase

import (
	"context"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
)

type SpendingTypeUsecase interface {
	Create(ctx context.Context, req *RequestCreateSpendingType) (*ResponseSpendingType, error)
	Update(ctx context.Context, req *RequestUpdateSpendingType) (*ResponseSpendingType, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseSpendingType, error)
	GetAllByPeriodeAndProfileID(ctx context.Context, req *RequestGetAllSpendingTypeByPeriodeWithISD) (*ResponseAllSpendingType, string, error)
	GetAllByProfileID(ctx context.Context, req *RequestGetAllByProfileIDWithISD) (*[]ResponseSpendingType, string, error)
}

type RequestGetAllSpendingTypeByPeriodeWithISD struct {
	Periode int
	RequestGetAllByProfileIDWithISD
}

type RequestCreateSpendingType struct {
	ProfileID    string
	Title        string
	MaximumLimit int
	Icon         string
}

type RequestUpdateSpendingType struct {
	ID           string
	ProfileID    string
	Title        string
	MaximumLimit int
	Icon         string
}

type ResponseSpendingType struct {
	ID                 string
	ProfileID          string
	Title              string
	MaximumLimit       int
	FormatMaximumLimit string
	Icon               string
}

type ResponseAllSpendingType struct {
	ResponseSpendingType *[]ResponseSpendingTypeJoinTable
	BudgetAmount         int
	FormatBudgetAmount   string
}

type ResponseSpendingTypeJoinTable struct {
	ID                 string
	ProfileID          string
	Title              string
	MaximumLimit       int
	FormatMaximumLimit string
	Icon               string
	Used               int
	FormatUsed         string
	PersentaseUsed     string
}

func (req *RequestCreateSpendingType) ToModel() *repository.SpendingType {
	id := ulid.Make().String()
	spendingType := &repository.SpendingType{
		ID:           id,
		ProfileID:    req.ProfileID,
		Title:        req.Title,
		MaximumLimit: req.MaximumLimit,
		Icon:         req.Icon,
		AuditInfo: repository.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}

	return spendingType
}

func (req *RequestUpdateSpendingType) ToModel() *repository.SpendingType {
	spendingType := &repository.SpendingType{
		ID:           req.ID,
		ProfileID:    req.ProfileID,
		Title:        req.Title,
		MaximumLimit: req.MaximumLimit,
		Icon:         req.Icon,
		AuditInfo: repository.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: repository.NewNullString(req.ProfileID),
		},
	}

	return spendingType
}
