package converter

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/helper"
)

func SpendingTypeRequestCreateToModel(req *domain.RequestCreateSpendingType) *domain.SpendingType {
	id := uuid.NewV4().String()
	spendingType := &domain.SpendingType{
		ID:           id,
		ProfileID:    req.ProfileID,
		Title:        req.Title,
		MaximumLimit: req.MaximumLimit,
		Icon:         req.Icon,
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}

	return spendingType
}

func SpendingTypeRequestUpdateToModel(req *domain.RequestUpdateSpendingType) *domain.SpendingType {
	spendingType := &domain.SpendingType{
		ID:           req.ID,
		ProfileID:    req.ProfileID,
		Title:        req.Title,
		MaximumLimit: req.MaximumLimit,
		Icon:         req.Icon,
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper.NewNullString(req.ProfileID),
		},
	}

	return spendingType
}

func SpendingTypeModelToResponse(spendingType *domain.SpendingType, formatMaximumLimit string) *domain.ResponseSpendingType {
	resp := &domain.ResponseSpendingType{
		ID:                 spendingType.ID,
		ProfileID:          spendingType.ProfileID,
		Title:              spendingType.Title,
		MaximumLimit:       spendingType.MaximumLimit,
		FormatMaximumLimit: formatMaximumLimit,
		Icon:               spendingType.Icon,
	}

	return resp
}

func SpendingTypeModelJoinToResponse(spendingType domain.SpendingTypeJoin, usedPersentase string, formatMaximumLimit string, formatUsed string) domain.ResponseSpendingTypeJoin {
	resp := domain.ResponseSpendingTypeJoin{
		ID:                 spendingType.ID,
		ProfileID:          spendingType.ProfileID,
		Title:              spendingType.Title,
		MaximumLimit:       spendingType.MaximumLimit,
		FormatMaximumLimit: formatMaximumLimit,
		Icon:               spendingType.Icon,
		Used:               spendingType.Used,
		FormatUsed:         formatUsed,
		PersentaseUsed:     usedPersentase,
	}

	return resp
}
