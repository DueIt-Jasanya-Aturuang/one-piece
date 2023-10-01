package rapi

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/presentation/rapi/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/usecase"
)

type BalanceHandlerImpl struct {
	balanceUsecase domain.BalanceUsecase
}

func NewBalanceHandlerImpl(balanceUsecase domain.BalanceUsecase) *BalanceHandlerImpl {
	return &BalanceHandlerImpl{
		balanceUsecase: balanceUsecase,
	}
}

func (b BalanceHandlerImpl) GetByProfileID(w http.ResponseWriter, r *http.Request) {
	profileID := r.Header.Get("Profile-ID")

	if _, err := uuid.Parse(profileID); err != nil {
		err := _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp, err := b.balanceUsecase.GetByProfileID(r.Context(), profileID)
	if err != nil {
		if errors.Is(err, usecase.BalanceNotExist) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, resp, "data balance profil")
}