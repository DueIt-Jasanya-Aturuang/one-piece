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
		AuditInfo: domain.AuditInfo{
			CreatedAt: time.Now().Unix(),
			CreatedBy: "profileID1",
			UpdatedAt: time.Now().Unix(),
		},
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
		AuditInfo: domain.AuditInfo{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: sql.NullString{String: "profileID1", Valid: true},
		},
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

func GetByIDSpendingTypeERRORDeletedAtNull(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	spendingType, err := SpendingTypeRepo.GetByID(context.TODO(), "spendingType1")
	assert.Error(t, err)
	assert.Nil(t, spendingType)
	assert.Equal(t, sql.ErrNoRows, err)
}

func GetByIDSpendingTypeERRORInvalidID(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	spendingType, err := SpendingTypeRepo.GetByID(context.TODO(), "spendingType1123")
	assert.Error(t, err)
	assert.Nil(t, spendingType)
	assert.Equal(t, sql.ErrNoRows, err)
}

func GetByIDAndProfileIDSpendingType(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	spendingType, err := SpendingTypeRepo.GetByIDAndProfileID(context.TODO(), "spendingType2", "profileID1")
	assert.NoError(t, err)
	assert.NotNil(t, spendingType)
	t.Log(spendingType)
}

func GetByIDAndProfileIDSpendingTypeERRORDeletedAtNull(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	spendingType, err := SpendingTypeRepo.GetByIDAndProfileID(context.TODO(), "spendingType1", "profileID1")
	assert.Error(t, err)
	assert.Nil(t, spendingType)
	assert.Equal(t, sql.ErrNoRows, err)
}

func GetByIDAndProfileIDSpendingTypeERRORInvalidID(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	spendingType, err := SpendingTypeRepo.GetByIDAndProfileID(context.TODO(), "spendingType1123", "profileID1")
	assert.Error(t, err)
	assert.Nil(t, spendingType)
	assert.Equal(t, sql.ErrNoRows, err)
}

func GetByIDAndProfileIDSpendingTypeERRORInvalidProfileID(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	spendingType, err := SpendingTypeRepo.GetByIDAndProfileID(context.TODO(), "spendingType2", "profileID1123")
	assert.Error(t, err)
	assert.Nil(t, spendingType)
	assert.Equal(t, sql.ErrNoRows, err)
}

func GetAllByProfileIDSpendingType(t *testing.T) {
	err := SpendingTypeRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	spendingTypes, err := SpendingTypeRepo.GetAllByProfileID(context.TODO(), "profileID1")
	assert.NoError(t, err)
	assert.NotNil(t, spendingTypes)
	t.Log(spendingTypes)
	assert.Equal(t, 4, len(*spendingTypes))
}
