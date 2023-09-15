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

func CreateSpendingHistory(t *testing.T) {
	spendingHistory := &domain.SpendingHistory{
		ID:                      "spendingHistory",
		ProfileID:               "profileID1",
		SpendingTypeID:          "spendingType1",
		PaymentMethodID:         sql.NullString{},
		PaymentName:             sql.NullString{String: "dijajanin", Valid: true},
		BeforeBalance:           0,
		SpendingAmount:          15000,
		AfterBalance:            15000,
		Description:             "dijajanin",
		Location:                "depok",
		TimeSpendingHistory:     time.Now().UTC(),
		ShowTimeSpendingHistory: "depok jawa barat",
		AuditInfo: domain.AuditInfo{
			CreatedAt: 0,
			CreatedBy: "profileID1",
			UpdatedAt: 0,
		},
	}

	for i := 0; i < 5; i++ {
		if i == 4 {
			spendingHistory.TimeSpendingHistory = time.Now().Add(-24 * time.Hour).UTC()
		}
		if i == 3 {
			spendingHistory.PaymentMethodID = sql.NullString{String: "payment1", Valid: true}
			spendingHistory.PaymentName = sql.NullString{}
		}
		spendingHistory.ID = fmt.Sprintf("%s%d", spendingHistory.ID, i+1)
		err := SpendingHistoryRepo.OpenConn(context.TODO())
		assert.NoError(t, err)

		err = SpendingHistoryRepo.StartTx(context.TODO(), &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		}, func() error {
			err = SpendingHistoryRepo.Create(context.TODO(), spendingHistory)
			if err != nil {
				return err
			}
			return nil
		})
		SpendingTypeRepo.CloseConn()
		assert.NoError(t, err)
	}
}

func UpdateSpendingHistory(t *testing.T) {
	spendingHistory := &domain.SpendingHistory{
		ID:                      "spendingHistory1",
		ProfileID:               "profileID1",
		SpendingTypeID:          "spendingType1",
		PaymentMethodID:         sql.NullString{},
		PaymentName:             sql.NullString{String: "dijajanin", Valid: true},
		BeforeBalance:           0,
		SpendingAmount:          20000,
		AfterBalance:            20000,
		Description:             "dijajanin",
		Location:                "depok",
		TimeSpendingHistory:     time.Now().UTC(),
		ShowTimeSpendingHistory: "depok jawa barat",
		AuditInfo: domain.AuditInfo{
			UpdatedAt: 0,
			UpdatedBy: sql.NullString{String: "profileID1", Valid: true},
		},
	}

	err := SpendingHistoryRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	err = SpendingHistoryRepo.StartTx(context.TODO(), &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = SpendingHistoryRepo.Update(context.TODO(), spendingHistory)
		if err != nil {
			return err
		}
		return nil
	})
	assert.NoError(t, err)
}

func DeleteSpendingHistory(t *testing.T) {
	err := SpendingHistoryRepo.OpenConn(context.TODO())
	assert.NoError(t, err)
	defer SpendingTypeRepo.CloseConn()

	err = SpendingHistoryRepo.StartTx(context.TODO(), &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = SpendingHistoryRepo.Delete(context.TODO(), "spendingHistory1", "profileID1")
		if err != nil {
			return err
		}
		return nil
	})
	assert.NoError(t, err)
}

func GetAllByTimeAndProfileIDSpendingHistory(t *testing.T) {

	t.Run("time now", func(t *testing.T) {
		err := SpendingHistoryRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer SpendingTypeRepo.CloseConn()

		startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().UTC().Day(), 0, 0, 0, 0, time.UTC)
		endTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().UTC().Day(), 23, 59, 59, 0, time.UTC)
		spendingHistories, err := SpendingHistoryRepo.GetAllByTimeAndProfileID(context.TODO(), &domain.RequestGetFilteredDataSpendingHistory{
			ProfileID: "profileID1",
			StartTime: startTime,
			EndTime:   endTime,
		})
		assert.NoError(t, err)
		assert.NotNil(t, spendingHistories)
		assert.Equal(t, 3, len(*spendingHistories))
		t.Log(spendingHistories)
	})

	t.Run("time before one day", func(t *testing.T) {
		err := SpendingHistoryRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer SpendingTypeRepo.CloseConn()

		startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().UTC().Day(), 0, 0, 0, 0, time.UTC)
		endTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().UTC().Day(), 23, 59, 59, 0, time.UTC)
		startTime = startTime.Add(-24 * time.Hour)
		endTime = endTime.Add(-24 * time.Hour)
		spendingHistories, err := SpendingHistoryRepo.GetAllByTimeAndProfileID(context.TODO(), &domain.RequestGetFilteredDataSpendingHistory{
			ProfileID: "profileID1",
			StartTime: startTime,
			EndTime:   endTime,
		})
		assert.NoError(t, err)
		assert.NotNil(t, spendingHistories)
		assert.Equal(t, 1, len(*spendingHistories))

	})
}

func GetByIDAndProfileIDSpendingHistory(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		err := SpendingHistoryRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer SpendingTypeRepo.CloseConn()

		spendingHistories, err := SpendingHistoryRepo.GetByIDAndProfileID(context.TODO(), "spendingHistory123", "profileID1")
		assert.NoError(t, err)
		assert.NotNil(t, spendingHistories)
	})

	t.Run("ERROR-norows", func(t *testing.T) {
		err := SpendingHistoryRepo.OpenConn(context.TODO())
		assert.NoError(t, err)
		defer SpendingTypeRepo.CloseConn()

		spendingHistories, err := SpendingHistoryRepo.GetByIDAndProfileID(context.TODO(), "nil", "profileID1")
		assert.Error(t, err)
		assert.Nil(t, spendingHistories)
		assert.Equal(t, sql.ErrNoRows, err)
	})
}
