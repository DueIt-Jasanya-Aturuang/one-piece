package integration

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreateSpendingType(t *testing.T) {
	spendingType := &domain.SpendingType{
		ID:           "spendingType1",
		ProfileID:    "profileID1",
		Title:        "title",
		MaximumLimit: 15000000,
		CreatedAt:    time.Now().Unix(),
		CreatedBy:    "profileID1",
		UpdatedAt:    time.Now().Unix(),
	}

	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	err = SpendingTypeRepo.StartTx(context.TODO(), &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = SpendingTypeRepo.Create(context.TODO(), spendingType)
		if err != nil {
			return err
		}
		return nil
	})
	assert.NoError(t, err)
}

func UpdateSpendingType(t *testing.T) {
	spendingType := &domain.SpendingType{
		ID:           "spendingType1",
		ProfileID:    "profileID1",
		Title:        "title",
		MaximumLimit: 15000000,
		UpdatedAt:    time.Now().Unix(),
		UpdatedBy:    sql.NullString{String: "profileID1", Valid: true},
	}

	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	err = SpendingTypeRepo.StartTx(context.TODO(), &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = SpendingTypeRepo.Update(context.TODO(), spendingType)
		if err != nil {
			return err
		}
		return nil
	})
	assert.NoError(t, err)
}
