package converter

import (
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase/helper"
)

func CreateIncomeHistoryToModel(req *domain.RequestCreateIncomeHistory) *domain.IncomeHistory {
	id := ulid.Make().String()
	timeIncomeHistory, _ := time.Parse("2006-01-02", req.TimeIncomeHistory)
	spendingHistory := &domain.IncomeHistory{
		ID:                    id,
		ProfileID:             req.ProfileID,
		IncomeTypeID:          req.IncomeTypeID,
		PaymentMethodID:       helper.NewNullString(req.PaymentMethodID),
		PaymentName:           helper.NewNullString(req.PaymentName),
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

func UpdateIncomeHistoryToModel(req *domain.RequestUpdateIncomeHistory) *domain.IncomeHistory {
	timeIncomeHistory, _ := time.Parse("2006-01-02", req.TimeIncomeHistory)
	spendingHistory := &domain.IncomeHistory{
		ID:                    req.ID,
		ProfileID:             req.ProfileID,
		IncomeTypeID:          req.IncomeTypeID,
		PaymentMethodID:       helper.NewNullString(req.PaymentMethodID),
		PaymentName:           helper.NewNullString(req.PaymentName),
		IncomeAmount:          req.IncomeAmount,
		Description:           req.Description,
		TimeIncomeHistory:     timeIncomeHistory.UTC(),
		ShowTimeIncomeHistory: req.ShowTimeIncomeHistory,
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: helper.NewNullString(req.ProfileID),
		},
	}
	return spendingHistory
}

func IncomeHistoryJoinModelToResponse(incomeHistory *domain.IncomeHistoryJoin) *domain.ResponseIncomeHistory {
	resp := &domain.ResponseIncomeHistory{
		ID:                    incomeHistory.ID,
		ProfileID:             incomeHistory.ProfileID,
		IncomeTypeID:          incomeHistory.IncomeTypeID,
		IncomeTypeTitle:       incomeHistory.IncomeTypeTitle,
		PaymentMethodID:       helper.GetNullString(incomeHistory.PaymentMethodID),
		PaymentMethodName:     helper.GetNullString(incomeHistory.PaymentMethodName),
		PaymentName:           helper.GetNullString(incomeHistory.PaymentName),
		IncomeAmount:          incomeHistory.IncomeAmount,
		FormatIncomeAmount:    helper.FormatRupiah(incomeHistory.IncomeAmount),
		Description:           incomeHistory.Description,
		TimeIncomeHistory:     incomeHistory.TimeIncomeHistory.UTC(),
		ShowTimeIncomeHistory: incomeHistory.ShowTimeIncomeHistory,
	}

	return resp
}

func GetAllIncomeHistoryJoinModelToResponse(spendingHistory domain.IncomeHistoryJoin) domain.ResponseIncomeHistory {
	resp := domain.ResponseIncomeHistory{
		ID:                    spendingHistory.ID,
		ProfileID:             spendingHistory.ProfileID,
		IncomeTypeID:          spendingHistory.IncomeTypeID,
		IncomeTypeTitle:       spendingHistory.IncomeTypeTitle,
		PaymentMethodID:       helper.GetNullString(spendingHistory.PaymentMethodID),
		PaymentMethodName:     helper.GetNullString(spendingHistory.PaymentMethodName),
		PaymentName:           helper.GetNullString(spendingHistory.PaymentName),
		IncomeAmount:          spendingHistory.IncomeAmount,
		FormatIncomeAmount:    helper.FormatRupiah(spendingHistory.IncomeAmount),
		Description:           spendingHistory.Description,
		TimeIncomeHistory:     spendingHistory.TimeIncomeHistory.UTC(),
		ShowTimeIncomeHistory: spendingHistory.ShowTimeIncomeHistory,
	}

	return resp
}
