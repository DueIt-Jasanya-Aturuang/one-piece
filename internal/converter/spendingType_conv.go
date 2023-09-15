package converter

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/helper"
)

func RequestCreateSpendingTypeToModel(req *domain.RequestCreateSpendingType) *domain.SpendingType {
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

func RequestUpdateSpendingTypeToModel(req *domain.RequestUpdateSpendingType) *domain.SpendingType {
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

func ModelSpendingTypeToResponse(spendingType *domain.SpendingType) *domain.ResponseSpendingType {
	resp := &domain.ResponseSpendingType{
		ID:           spendingType.ID,
		ProfileID:    spendingType.ProfileID,
		Title:        spendingType.Title,
		MaximumLimit: spendingType.MaximumLimit,
		Icon:         spendingType.Icon,
	}

	return resp
}
