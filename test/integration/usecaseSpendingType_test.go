package integration

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreateSpendingTypeUsecase(t *testing.T) {
	req := &domain.RequestCreateSpendingType{
		ProfileID:    "profileID1",
		Title:        "jajan",
		MaximumLimit: 100000,
		Icon:         "icon.png",
	}

	resp, err := SpendingTypeUsecase.Create(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func CreateSpendingTypeUsecaseERRORNameAlready(t *testing.T) {
	req := &domain.RequestCreateSpendingType{
		ProfileID:    "profileID1",
		Title:        "jajan",
		MaximumLimit: 100000,
		Icon:         "icon.png",
	}

	resp, err := SpendingTypeUsecase.Create(context.TODO(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	var errHTTP *domain.ErrHTTP
	assert.Equal(t, true, errors.As(err, &errHTTP))
}

func UpdateSpendingTypeUsecase(t *testing.T) {
	req := &domain.RequestUpdateSpendingType{
		ID:           "spendingType2",
		ProfileID:    "profileID1",
		Title:        "jajan2",
		MaximumLimit: 100000,
		Icon:         "icon.png",
	}

	resp, err := SpendingTypeUsecase.Update(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func UpdateSpendingTypeUsecaseERRORNameAlready(t *testing.T) {
	req := &domain.RequestUpdateSpendingType{
		ID:           "spendingType3",
		ProfileID:    "profileID1",
		Title:        "jajan",
		MaximumLimit: 100000,
		Icon:         "icon.png",
	}

	resp, err := SpendingTypeUsecase.Update(context.TODO(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	var errHTTP *domain.ErrHTTP
	assert.Equal(t, true, errors.As(err, &errHTTP))
}

func DeleteSpendingTypeUsecase(t *testing.T) {
	err := SpendingTypeUsecase.Delete(context.TODO(), "spendingType3", "profileID1")
	assert.NoError(t, err)
}

func GetByIDAndProfileIDSpendingTypeUsecase(t *testing.T) {
	resp, err := SpendingTypeUsecase.GetByIDAndProfileID(context.TODO(), "spendingType2", "profileID1")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Log(resp)
}

func GetByIDAndProfileIDSpendingTypeUsecaseERRORNoRow(t *testing.T) {
	resp, err := SpendingTypeUsecase.GetByIDAndProfileID(context.TODO(), "spendingType3", "profileID1")
	assert.Error(t, err)
	assert.Nil(t, resp)
	var errHTTP *domain.ErrHTTP
	assert.Equal(t, true, errors.As(err, &errHTTP))
	assert.Equal(t, 404, errHTTP.Code)
}

func GetAllByProfileIDSpendingTypeUsecase(t *testing.T) {
	resp, err := SpendingTypeUsecase.GetAllByProfileID(context.TODO(), "profileID1", 14)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Log(resp.ResponseSpendingType)
}

func GetAllByProfileIDSpendingTypeUsecaseWithCreateDefaultType(t *testing.T) {
	resp, err := SpendingTypeUsecase.GetAllByProfileID(context.TODO(), "profileID2", 14)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 3, len(*resp.ResponseSpendingType))
}
