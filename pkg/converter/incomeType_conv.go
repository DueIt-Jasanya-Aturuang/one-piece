package converter

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/helper"
)

func RequestCreateIncomeTypeToModel(i *domain.RequestCreateIncomeType) *domain.IncomeType {
	id := uuid.NewV4().String()
	return &domain.IncomeType{
		ID:          id,
		ProfileID:   i.ProfileID,
		Name:        i.Name,
		Description: helper.NewNullString(i.Description),
		Icon:        i.Icon,
		IncomeType:  "lainnya",
		FixedIncome: helper.NewNullBool(false),
		Periode:     helper.NewNullString(""),
		Amount:      helper.NewNullInt64(0),
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: i.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
}

func RequestUpdateIncomeTypeToModel(i *domain.RequestUpdateIncomeType) *domain.IncomeType {
	return &domain.IncomeType{
		ID:          i.ID,
		ProfileID:   i.ProfileID,
		Name:        i.Name,
		Description: helper.NewNullString(i.Description),
		Icon:        i.Icon,
		IncomeType:  "lainnya",
		FixedIncome: helper.NewNullBool(false),
		Periode:     helper.NewNullString(""),
		Amount:      helper.NewNullInt64(0),
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper.NewNullString(i.ProfileID),
		},
	}
}

func IncomeTypeModelToResp(i *domain.IncomeType) *domain.ResponseIncomeType {
	return &domain.ResponseIncomeType{
		ID:          i.ID,
		ProfileID:   i.ProfileID,
		Name:        i.Name,
		Description: helper.GetNullString(i.Description),
		Icon:        i.Icon,
	}
}
