package integration

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
)

func CreatePayment(t *testing.T) {
	payment := &domain.Payment{
		ID:          "payment1",
		Name:        "bca",
		Description: sql.NullString{},
		Image:       "bca.png",
		CreatedAt:   0,
		CreatedBy:   "userID1",
		UpdatedAt:   0,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}

	ctx := context.Background()
	err := PaymentRepo.OpenConn(ctx)
	assert.NoError(t, err)
	defer PaymentRepo.CloseConn()

	err = PaymentRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = PaymentRepo.CreatePayment(context.TODO(), payment)
		if err != nil {
			return err
		}

		return nil
	})

}

func GetPaymentById(t *testing.T) {
	ctx := context.Background()
	err := PaymentRepo.OpenConn(ctx)
	assert.NoError(t, err)
	defer PaymentRepo.CloseConn()

	payment, err := PaymentRepo.GetPaymentByID(ctx, "payment1")
	assert.NoError(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, "bca", payment.Name)
}

func GetPaymentByIdERROR(t *testing.T) {
	ctx := context.Background()
	err := PaymentRepo.OpenConn(ctx)
	assert.NoError(t, err)
	defer PaymentRepo.CloseConn()

	payment, err := PaymentRepo.GetPaymentByID(ctx, "payment1nill")
	assert.Error(t, err)
	assert.Nil(t, payment)
	assert.Equal(t, sql.ErrNoRows, err)
}

func UpdatePayment(t *testing.T) {
	payment := &domain.Payment{
		ID:          "payment1",
		Name:        "bca",
		Description: sql.NullString{String: "payment bca", Valid: true},
		Image:       "bca2.png",
		CreatedAt:   0,
		CreatedBy:   "userID1",
		UpdatedAt:   0,
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}

	ctx := context.Background()
	err := PaymentRepo.OpenConn(ctx)
	assert.NoError(t, err)
	defer PaymentRepo.CloseConn()

	err = PaymentRepo.StartTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}, func() error {
		err = PaymentRepo.UpdatePayment(context.TODO(), payment)
		if err != nil {
			return err
		}

		return nil
	})

}

func GetPaymentByName(t *testing.T) {
	ctx := context.Background()
	err := PaymentRepo.OpenConn(ctx)
	assert.NoError(t, err)
	defer PaymentRepo.CloseConn()

	payment, err := PaymentRepo.GetPaymentByName(ctx, "bca")
	assert.NoError(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, "payment bca", payment.Description.String)
	assert.Equal(t, true, payment.Description.Valid)
}

func GetPaymentByNameERROR(t *testing.T) {
	ctx := context.Background()
	err := PaymentRepo.OpenConn(ctx)
	assert.NoError(t, err)
	defer PaymentRepo.CloseConn()

	payment, err := PaymentRepo.GetPaymentByName(ctx, "namepaymentnil")
	assert.Error(t, err)
	assert.Nil(t, payment)
	assert.Equal(t, sql.ErrNoRows, err)
}
