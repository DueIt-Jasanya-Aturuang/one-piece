package rapi

import (
	"errors"
	"net/http"

	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/schema"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (p *Presenter) GetBalanceByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if err := util.ParseUUID(profileID); err != nil {
		helper.ErrorResponseEncode(w, _error.HttpErrString(response.CodeCompanyName[response.CM04], response.CM04))
		return
	}

	balance, err := p.balanceUsecase.GetOrCreateByProfileID(r.Context(), profileID)
	if err != nil {
		if errors.Is(err, usecase.BalanceNotExist) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	resp := &schema.ResponseBalance{
		ID:                        balance.ID,
		ProfileID:                 balance.ProfileID,
		TotalIncomeAmount:         balance.TotalIncomeAmount,
		TotalIncomeAmountFormat:   balance.TotalIncomeAmountFormat,
		TotalSpendingAmount:       balance.TotalSpendingAmount,
		TotalSpendingAmountFormat: balance.TotalSpendingAmountFormat,
		Balance:                   balance.Balance,
		BalanceFormat:             balance.BalanceFormat,
	}
	helper.SuccessResponseEncode(w, resp, "data balance profil")
}
