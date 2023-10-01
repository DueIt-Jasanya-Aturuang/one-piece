package converter

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	helper2 "github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/helper"
)

func CreateIncomeHistoryToModel(req *domain.RequestCreateIncomeHistory, balance int) *domain.IncomeHistory {
	id := uuid.NewV4().String()
	timeIncomeHistory, _ := time.Parse("2006-01-02", req.TimeIncomeHistory)
	spendingHistory := &domain.IncomeHistory{
		ID:                    id,
		ProfileID:             req.ProfileID,
		IncomeTypeID:          req.IncomeTypeID,
		PaymentMethodID:       helper2.NewNullString(req.PaymentMethodID),
		PaymentName:           helper2.NewNullString(req.PaymentName),
		IncomeAmount:          req.IncomeAmount,
		Description:           req.Description,
		TimeIncomeHistory:     timeIncomeHistory.UTC(),
		ShowTimeIncomeHistory: req.ShowTimeIncomeHistory,
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: req.ProfileID,
			UpdatedAt: time.Now().Unix(),
		},
	}
	return spendingHistory
}

func UpdateIncomeHistoryToModel(req *domain.RequestUpdateIncomeHistory, balance int) *domain.IncomeHistory {
	timeIncomeHistory, _ := time.Parse("2006-01-02", req.TimeIncomeHistory)
	spendingHistory := &domain.IncomeHistory{
		ID:                    req.ID,
		ProfileID:             req.ProfileID,
		IncomeTypeID:          req.IncomeTypeID,
		PaymentMethodID:       helper2.NewNullString(req.PaymentMethodID),
		PaymentName:           helper2.NewNullString(req.PaymentName),
		IncomeAmount:          req.IncomeAmount,
		Description:           req.Description,
		TimeIncomeHistory:     timeIncomeHistory.UTC(),
		ShowTimeIncomeHistory: req.ShowTimeIncomeHistory,
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper2.NewNullString(req.ProfileID),
		},
	}
	return spendingHistory
}

func IncomeHistoryJoinModelToResponse(spendingHistory *domain.IncomeHistoryJoin) *domain.ResponseIncomeHistory {
	resp := &domain.ResponseIncomeHistory{
		ID:                    spendingHistory.ID,
		ProfileID:             spendingHistory.ProfileID,
		IncomeTypeID:          spendingHistory.IncomeTypeID,
		IncomeTypeTitle:       spendingHistory.IncomeTypeTitle,
		PaymentMethodID:       helper2.GetNullString(spendingHistory.PaymentMethodID),
		PaymentMethodName:     helper2.GetNullString(spendingHistory.PaymentMethodName),
		PaymentName:           helper2.GetNullString(spendingHistory.PaymentName),
		IncomeAmount:          spendingHistory.IncomeAmount,
		FormatIncomeAmount:    helper2.FormatRupiah(spendingHistory.IncomeAmount),
		Description:           spendingHistory.Description,
		TimeIncomeHistory:     spendingHistory.TimeIncomeHistory.UTC(),
		ShowTimeIncomeHistory: spendingHistory.ShowTimeIncomeHistory,
	}

	return resp
}

func GetAllIncomeHistoryJoinModelToResponse(spendingHistory domain.IncomeHistoryJoin) domain.ResponseIncomeHistory {
	resp := domain.ResponseIncomeHistory{
		ID:                    spendingHistory.ID,
		ProfileID:             spendingHistory.ProfileID,
		IncomeTypeID:          spendingHistory.IncomeTypeID,
		IncomeTypeTitle:       spendingHistory.IncomeTypeTitle,
		PaymentMethodID:       helper2.GetNullString(spendingHistory.PaymentMethodID),
		PaymentMethodName:     helper2.GetNullString(spendingHistory.PaymentMethodName),
		PaymentName:           helper2.GetNullString(spendingHistory.PaymentName),
		IncomeAmount:          spendingHistory.IncomeAmount,
		FormatIncomeAmount:    helper2.FormatRupiah(spendingHistory.IncomeAmount),
		Description:           spendingHistory.Description,
		TimeIncomeHistory:     spendingHistory.TimeIncomeHistory.UTC(),
		ShowTimeIncomeHistory: spendingHistory.ShowTimeIncomeHistory,
	}

	return resp
}
