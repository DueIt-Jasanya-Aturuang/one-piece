package converter

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	helper2 "github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/helper"
)

func CreateSpendingHistoryToModel(req *domain.RequestCreateSpendingHistory, balance int) *domain.SpendingHistory {
	id := uuid.NewV4().String()
	timeSpendingHistory, _ := time.Parse("2006-01-02", req.TimeSpendingHistory)
	spendingHistory := &domain.SpendingHistory{
		ID:                      id,
		ProfileID:               req.ProfileID,
		SpendingTypeID:          req.SpendingTypeID,
		PaymentMethodID:         helper2.NewNullString(req.PaymentMethodID),
		PaymentName:             helper2.NewNullString(req.PaymentName),
		BeforeBalance:           balance,
		SpendingAmount:          req.SpendingAmount,
		AfterBalance:            balance - req.SpendingAmount,
		Description:             req.Description,
		TimeSpendingHistory:     timeSpendingHistory.UTC(),
		ShowTimeSpendingHistory: req.ShowTimeSpendingHistory,
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
	return spendingHistory
}

func UpdateSpendingHistoryToModel(req *domain.RequestUpdateSpendingHistory, balance int) *domain.SpendingHistory {
	timeSpendingHistory, _ := time.Parse("2006-01-02", req.TimeSpendingHistory)
	spendingHistory := &domain.SpendingHistory{
		ID:                      req.ID,
		ProfileID:               req.ProfileID,
		SpendingTypeID:          req.SpendingTypeID,
		PaymentMethodID:         helper2.NewNullString(req.PaymentMethodID),
		PaymentName:             helper2.NewNullString(req.PaymentName),
		BeforeBalance:           balance,
		SpendingAmount:          req.SpendingAmount,
		AfterBalance:            balance - req.SpendingAmount,
		Description:             req.Description,
		TimeSpendingHistory:     timeSpendingHistory.UTC(),
		ShowTimeSpendingHistory: req.ShowTimeSpendingHistory,
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper2.NewNullString(req.ProfileID),
		},
	}
	return spendingHistory
}

func SpendingHistoryJoinModelToResponse(spendingHistory *domain.SpendingHistoryJoin) *domain.ResponseSpendingHistory {
	resp := &domain.ResponseSpendingHistory{
		ID:                      spendingHistory.ID,
		ProfileID:               spendingHistory.ProfileID,
		SpendingTypeID:          spendingHistory.SpendingTypeID,
		SpendingTypeTitle:       spendingHistory.SpendingTypeTitle,
		PaymentMethodID:         helper2.GetNullString(spendingHistory.PaymentMethodID),
		PaymentMethodName:       helper2.GetNullString(spendingHistory.PaymentMethodName),
		PaymentName:             helper2.GetNullString(spendingHistory.PaymentName),
		BeforeBalance:           spendingHistory.BeforeBalance,
		SpendingAmount:          spendingHistory.SpendingAmount,
		FormatSpendingAmount:    helper2.FormatRupiah(spendingHistory.SpendingAmount),
		AfterBalance:            spendingHistory.AfterBalance,
		Description:             spendingHistory.Description,
		TimeSpendingHistory:     spendingHistory.TimeSpendingHistory.UTC(),
		ShowTimeSpendingHistory: spendingHistory.ShowTimeSpendingHistory,
	}

	return resp
}

func GetAllSpendingHistoryJoinModelToResponse(spendingHistory domain.SpendingHistoryJoin) domain.ResponseSpendingHistory {
	resp := domain.ResponseSpendingHistory{
		ID:                      spendingHistory.ID,
		ProfileID:               spendingHistory.ProfileID,
		SpendingTypeID:          spendingHistory.SpendingTypeID,
		SpendingTypeTitle:       spendingHistory.SpendingTypeTitle,
		PaymentMethodID:         helper2.GetNullString(spendingHistory.PaymentMethodID),
		PaymentMethodName:       helper2.GetNullString(spendingHistory.PaymentMethodName),
		PaymentName:             helper2.GetNullString(spendingHistory.PaymentName),
		BeforeBalance:           spendingHistory.BeforeBalance,
		SpendingAmount:          spendingHistory.SpendingAmount,
		FormatSpendingAmount:    helper2.FormatRupiah(spendingHistory.SpendingAmount),
		AfterBalance:            spendingHistory.AfterBalance,
		Description:             spendingHistory.Description,
		TimeSpendingHistory:     spendingHistory.TimeSpendingHistory.UTC(),
		ShowTimeSpendingHistory: spendingHistory.ShowTimeSpendingHistory,
	}

	return resp
}
