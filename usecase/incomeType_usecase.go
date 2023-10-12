package usecase

import (
	"context"
	"time"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type IncomeTypeUsecase interface {
	Create(ctx context.Context, req *RequestCreateIncomeType) (*ResponseIncomeType, error)
	Update(ctx context.Context, req *RequestUpdateIncomeType) (*ResponseIncomeType, error)
	Delete(ctx context.Context, id string, profileID string) error
	GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*ResponseIncomeType, error)
	GetAllByProfileID(ctx context.Context, req *RequestGetAllByProfileIDWithISD) (*[]ResponseIncomeType, string, error)
}

type RequestCreateIncomeType struct {
	ProfileID   string
	Name        string
	Description string
	Icon        string
}

type RequestUpdateIncomeType struct {
	ID          string
	ProfileID   string
	Name        string
	Description string
	Icon        string
}

type ResponseIncomeType struct {
	ID          string
	ProfileID   string
	Name        string
	Description *string
	Icon        string
}

func (req *RequestCreateIncomeType) ToModel() *repository.IncomeType {
	id := util.NewUlid()
	return &repository.IncomeType{
		ID:          id,
		ProfileID:   req.ProfileID,
		Name:        req.Name,
		Description: repository.NewNullString(req.Description),
		Icon:        req.Icon,
		IncomeType:  "lainnya",
		FixedIncome: repository.NewNullBool(false),
		Periode:     repository.NewNullString(""),
		Amount:      repository.NewNullInt64(0),
		AuditInfo: repository.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
}

func (req *RequestUpdateIncomeType) ToModel() *repository.IncomeType {
	return &repository.IncomeType{
		ID:          req.ID,
		ProfileID:   req.ProfileID,
		Name:        req.Name,
		Description: repository.NewNullString(req.Description),
		Icon:        req.Icon,
		IncomeType:  "lainnya",
		FixedIncome: repository.NewNullBool(false),
		Periode:     repository.NewNullString(""),
		Amount:      repository.NewNullInt64(0),
		AuditInfo: repository.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: repository.NewNullString(req.ProfileID),
		},
	}
}
