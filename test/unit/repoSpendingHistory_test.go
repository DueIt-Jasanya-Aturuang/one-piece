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

func TestSpendingHistoryGetAllByTimeAndProfileID(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingHistoryRepo := _repository.NewSpendingHistoryRepositoryImpl(uow)

	query := regexp.QuoteMeta(`SELECT tsh.id, tsh.profile_id, tsh.spending_type_id, tsh.payment_method_id, tsh.payment_name, tsh.before_balance, 
       				tsh.spending_amount, tsh.after_balance, tsh.description, tsh.location, tsh.time_spending_history, tsh.show_time_spending_history, 
       				tsh.created_at, tsh.created_by, tsh.updated_at, tsh.updated_by, tsh.deleted_at, tsh.deleted_by,
       				mst.title, mpm.name
				FROM t_spending_history tsh 
				JOIN m_spending_type mst ON tsh.spending_type_id = mst.id
				LEFT JOIN m_payment_methods mpm ON tsh.payment_method_id = mpm.id
				WHERE tsh.profile_id = $1 AND tsh.time_spending_history BETWEEN $2 AND $3 AND tsh.deleted_at IS NULL`)

	rows := sqlmock.NewRows([]string{"id", "profile_id", "spending_type_id", "payment_method_id", "payment_name",
		"before_balance", "spending_amount", "after_balance", "description", "location", "time_spending_history", "show_time_spending_history",
		"created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by", "title", "name",
	})
	start := time.Now()
	end := time.Now()
	t.Run("SUCCESS", func(t *testing.T) {
		rows.AddRow("test", "test", "test", "test", nil, 0, 123, 123, "test", "depok", time.Now(), "test", 0, "test", 0, nil, nil, nil, "test", "test")
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			"profileID", start, end,
		).WillReturnRows(rows)

		err = spendingHistoryRepo.OpenConn(context.TODO())
		defer spendingHistoryRepo.CloseConn()

		spendingHistories, err := spendingHistoryRepo.GetAllByTimeAndProfileID(context.TODO(), &domain.RequestGetFilteredDataSpendingHistory{
			ProfileID: "profileID",
			StartTime: start,
			EndTime:   end,
		})
		assert.NoError(t, err)
		assert.NotNil(t, spendingHistories)
		assert.Equal(t, 1, len(*spendingHistories))
		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("SUCCESS_nil", func(t *testing.T) {
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			"profileID", start, end,
		).WillReturnRows(rows)

		err = spendingHistoryRepo.OpenConn(context.TODO())
		defer spendingHistoryRepo.CloseConn()

		spendingHistories, err := spendingHistoryRepo.GetAllByTimeAndProfileID(context.TODO(), &domain.RequestGetFilteredDataSpendingHistory{
			ProfileID: "profileID",
			StartTime: start,
			EndTime:   end,
		})
		assert.NoError(t, err)
		assert.NotNil(t, spendingHistories)
		assert.Equal(t, 0, len(*spendingHistories))
		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestSpendingHistoryGetByIDAndProfileID(t *testing.T) {
	db, mocksql, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	uow := _repository.NewUnitOfWorkRepositoryImpl(db)
	spendingHistoryRepo := _repository.NewSpendingHistoryRepositoryImpl(uow)

	query := regexp.QuoteMeta(`SELECT tsh.id, tsh.profile_id, tsh.spending_type_id, tsh.payment_method_id, tsh.payment_name, tsh.before_balance, 
       				tsh.spending_amount, tsh.after_balance, tsh.description, tsh.location, tsh.time_spending_history, tsh.show_time_spending_history, 
       				tsh.created_at, tsh.created_by, tsh.updated_at, tsh.updated_by, tsh.deleted_at, tsh.deleted_by,
       				mst.title, mpm.name
				FROM t_spending_history tsh 
				JOIN m_spending_type mst ON tsh.spending_type_id = mst.id
				LEFT JOIN m_payment_methods mpm ON tsh.payment_method_id = mpm.id
				WHERE tsh.profile_id = $1 AND tsh.id = $2 AND tsh.deleted_at IS NULL`)

	rows := sqlmock.NewRows([]string{"id", "profile_id", "spending_type_id", "payment_method_id", "payment_name",
		"before_balance", "spending_amount", "after_balance", "description", "location", "time_spending_history", "show_time_spending_history",
		"created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by", "title", "name",
	})
	t.Run("SUCCESS", func(t *testing.T) {
		rows.AddRow("test", "test", "test", "test", nil, 0, 123, 123, "test", "depok", time.Now(), "test", 0, "test", 0, nil, nil, nil, "test", "test")
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			"profileID", "id",
		).WillReturnRows(rows)

		err = spendingHistoryRepo.OpenConn(context.TODO())
		defer spendingHistoryRepo.CloseConn()

		spendingHistory, err := spendingHistoryRepo.GetByIDAndProfileID(context.TODO(), "id", "profileID")
		assert.NoError(t, err)
		assert.NotNil(t, spendingHistory)
		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ERROR_sqlnorows", func(t *testing.T) {
		mocksql.ExpectPrepare(query)
		mocksql.ExpectQuery(query).WithArgs(
			"profileID", "nil",
		).WillReturnError(sql.ErrNoRows)

		err = spendingHistoryRepo.OpenConn(context.TODO())
		defer spendingHistoryRepo.CloseConn()

		spendingHistory, err := spendingHistoryRepo.GetByIDAndProfileID(context.TODO(), "nil", "profileID")
		assert.Error(t, err)
		assert.Nil(t, spendingHistory)
		assert.Equal(t, sql.ErrNoRows, err)
		err = mocksql.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
