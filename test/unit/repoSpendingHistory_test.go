package unit

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/internal/_repository"
)

func TestSpendingHistoryCreate(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingHistoryRepo := _repository.NewSpendingHistoryRepositoryImpl(uow)
	spendingHistory := &domain.SpendingHistory{
		ID:                      "test",
		ProfileID:               "test",
		SpendingTypeID:          "test",
		PaymentMethodID:         sql.NullString{String: "test", Valid: true},
		PaymentName:             sql.NullString{},
		BeforeBalance:           0,
		SpendingAmount:          123,
		AfterBalance:            123,
		Description:             "jajan",
		Location:                "depok",
		TimeSpendingHistory:     time.Now().UTC(),
		ShowTimeSpendingHistory: "show time spending",
		AuditInfo: domain.AuditInfo{
			CreatedAt: 0,
			CreatedBy: "test",
			UpdatedAt: 0,
		},
	}
	query := regexp.QuoteMeta(`INSERT INTO t_spending_history (id, profile_id, spending_type_id, payment_method_id, payment_name, 
                        before_balance, spending_amount, after_balance, description, location, time_spending_history, show_time_spending_history, 
                        created_at, created_by, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`)

	mocksql.ExpectBegin()
	mocksql.ExpectPrepare(query)
	mocksql.ExpectExec(query).WithArgs(
		"test", "test", "test", "test", nil, 0, 123, 123, "jajan", "depok", spendingHistory.TimeSpendingHistory, "show time spending", 0, "test", 0,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mocksql.ExpectCommit()

	err = spendingHistoryRepo.OpenConn(context.TODO())
	defer spendingHistoryRepo.CloseConn()

	err = spendingHistoryRepo.StartTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false}, func() error {
		err = spendingHistoryRepo.Create(context.TODO(), spendingHistory)
		assert.NoError(t, err)
		return nil
	})

	err = mocksql.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestSpendingHistoryUpdate(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingHistoryRepo := _repository.NewSpendingHistoryRepositoryImpl(uow)
	spendingHistory := &domain.SpendingHistory{
		ID:                      "test",
		ProfileID:               "test",
		SpendingTypeID:          "test",
		PaymentMethodID:         sql.NullString{String: "test", Valid: true},
		PaymentName:             sql.NullString{},
		BeforeBalance:           0,
		SpendingAmount:          123,
		AfterBalance:            123,
		Description:             "jajan",
		Location:                "depok",
		TimeSpendingHistory:     time.Now().UTC(),
		ShowTimeSpendingHistory: "show time spending",
		AuditInfo: domain.AuditInfo{
			UpdatedAt: 0,
			UpdatedBy: sql.NullString{String: "test", Valid: true},
		},
	}
	query := regexp.QuoteMeta(`UPDATE t_spending_history SET spending_type_id = $1, payment_method_id = $2, payment_name = $3, before_balance = $4, 
                              spending_amount = $5, after_balance = $6, description = $7, location = $8, time_spending_history = $9, 
                              show_time_spending_history = $10, updated_at = $11, updated_by = $12 
                          WHERE id = $13 AND profile_id = $14 AND deleted_at IS NULL`)

	mocksql.ExpectBegin()
	mocksql.ExpectPrepare(query)
	mocksql.ExpectExec(query).WithArgs(
		"test", "test", nil, 0, 123, 123, "jajan", "depok", spendingHistory.TimeSpendingHistory, "show time spending", 0, "test", "test", "test",
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mocksql.ExpectCommit()

	err = spendingHistoryRepo.OpenConn(context.TODO())
	defer spendingHistoryRepo.CloseConn()

	err = spendingHistoryRepo.StartTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false}, func() error {
		err = spendingHistoryRepo.Update(context.TODO(), spendingHistory)
		assert.NoError(t, err)
		return nil
	})

	err = mocksql.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestSpendingHistoryDelete(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingHistoryRepo := _repository.NewSpendingHistoryRepositoryImpl(uow)

	query := regexp.QuoteMeta(`UPDATE t_spending_history SET deleted_at = $1, deleted_by = $2
                              WHERE id = $3 AND profile_id = $4`)

	mocksql.ExpectBegin()
	mocksql.ExpectPrepare(query)
	mocksql.ExpectExec(query).WithArgs(
		time.Now().Unix(), "profileID", "id", "profileID",
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mocksql.ExpectCommit()

	err = spendingHistoryRepo.OpenConn(context.TODO())
	defer spendingHistoryRepo.CloseConn()

	err = spendingHistoryRepo.StartTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false}, func() error {
		err = spendingHistoryRepo.Delete(context.TODO(), "id", "profileID")
		assert.NoError(t, err)
		return nil
	})

	err = mocksql.ExpectationsWereMet()
	assert.NoError(t, err)
}
