package integration

import (
	"context"
	"database/sql"
	"fmt"
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

	for i := 0; i < 5; i++ {
		spendingType.ID = fmt.Sprintf("spendingType%d", i+1)
		err := SpendingTypeRepo.OpenConn(context.TODO())
		assert.NoError(t, err)

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
		SpendingTypeRepo.CloseConn()
	}

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

func DeleteSpendingType(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	err = SpendingTypeRepo.StartTx(context.TODO(), &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = SpendingTypeRepo.Delete(context.TODO(), "spendingType1", "profileID1")
		if err != nil {
			return err
		}
		return nil
	})
	assert.NoError(t, err)
}

func GetByIDSpendingType(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	spendingType, err := SpendingTypeRepo.GetByID(context.TODO(), "spendingType2")
	assert.NoError(t, err)
	assert.NotNil(t, spendingType)
	t.Log(spendingType)
}

func GetByIDSpendingTypeERROR(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	spendingType, err := SpendingTypeRepo.GetByID(context.TODO(), "spendingType1")
	assert.Error(t, err)
	assert.Nil(t, spendingType)
	assert.Equal(t, sql.ErrNoRows, err)
}
