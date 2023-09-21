package rest

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jasanya-tech/jasanya-response-backend-golang/_error"
	"github.com/jasanya-tech/jasanya-response-backend-golang/response"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/api/rest/helper"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/pkg/_usecase"
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
	profileID := chi.URLParam(r, "profile-id")

	if _, err := uuid.Parse(profileID); err != nil {
		err := _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		helper.ErrorResponseEncode(w, err)
		return
	}

	resp, err := b.balanceUsecase.GetByProfileID(r.Context(), profileID)
	if err != nil {
		if errors.Is(err, _usecase.BalanceNotExist) {
			err = _error.HttpErrString(response.CodeCompanyName[response.CM01], response.CM01)
		}

		helper.ErrorResponseEncode(w, err)
		return
	}

	helper.SuccessResponseEncode(w, resp, "data balance profil")
}
