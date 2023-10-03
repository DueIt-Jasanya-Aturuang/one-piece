package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type IncomeHistoryRepositoryImpl struct {
	domain.UnitOfWorkRepository
}

func NewIncomeHistoryRepositoryImpl(uow domain.UnitOfWorkRepository) domain.IncomeHistoryRepository {
	return &IncomeHistoryRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}

func (i *IncomeHistoryRepositoryImpl) Create(ctx context.Context, income *domain.IncomeHistory) error {
	query := `INSERT INTO t_income_history(id, profile_id, income_type_id, payment_method_id, payment_name, income_amount,
                               description, time_income_history, show_time_income_history, created_at, created_by, updated_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	tx, err := i.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		income.ID,
		income.ProfileID,
		income.IncomeTypeID,
		income.PaymentMethodID,
		income.PaymentName,
		income.IncomeAmount,
		income.Description,
		income.TimeIncomeHistory,
		income.ShowTimeIncomeHistory,
		income.CreatedAt,
		income.CreatedBy,
		income.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (i *IncomeHistoryRepositoryImpl) Update(ctx context.Context, income *domain.IncomeHistory) error {
	query := `UPDATE t_income_history SET income_type_id=$1, payment_method_id=$2, payment_name=$3, income_amount=$4, description=$5,
                              time_income_history=$6, show_time_income_history=$7, updated_at=$8, updated_by=$9 
                          WHERE id =$10 AND profile_id = $11 AND deleted_at IS NULL`

	tx, err := i.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		income.IncomeTypeID,
		income.PaymentMethodID,
		income.PaymentName,
		income.IncomeAmount,
		income.Description,
		income.TimeIncomeHistory,
		income.ShowTimeIncomeHistory,
		income.UpdatedAt,
		income.UpdatedBy,
		income.ID,
		income.ProfileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (i *IncomeHistoryRepositoryImpl) Delete(ctx context.Context, id string, profileID string) error {
	query := `UPDATE t_income_history SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND profile_id=$4 AND deleted_at IS NULL`

	tx, err := i.GetTx()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		time.Now().Unix(),
		profileID,
		id,
		profileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (i *IncomeHistoryRepositoryImpl) GetAllByTimeAndProfileID(ctx context.Context, req *domain.GetFilteredDataIncomeHistory) (*[]domain.IncomeHistoryJoin, error) {
	query := `SELECT tih.id, tih.profile_id, tih.income_type_id, tih.payment_method_id, tih.payment_name, tih.income_amount, tih.description, 
       					tih.time_income_history, tih.show_time_income_history, tih.created_at, tih.created_by, tih.updated_at, tih.updated_by, 
       					tih.deleted_at, tih.deleted_by, mit.name, mpm.name
				FROM t_income_history tih
				JOIN m_income_type mit ON tih.income_type_id = mit.id
				LEFT JOIN m_payment_methods mpm ON tih.payment_method_id = mpm.id
				WHERE tih.profile_id=$1 AND tih.time_income_history BETWEEN $2 AND $3 AND tih.deleted_at IS NULL `
	if req.ID != "" {
		query += `AND tih.id ` + req.Operation + ` $4 `
	}
	query += `ORDER BY tih.id ` + req.Order + ` LIMIT 5`

	conn, err := i.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var rows *sql.Rows
	if req.ID != "" {
		rows, err = stmt.QueryContext(ctx, req.ProfileID, req.StartTime, req.EndTime, req.ID)
	} else {
		rows, err = stmt.QueryContext(ctx, req.ProfileID, req.StartTime, req.EndTime)
	}

	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
		return nil, err
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrQueryRowsClose, err)
		}
	}()

	var incomeTypes []domain.IncomeHistoryJoin
	var incomeType domain.IncomeHistoryJoin

	for rows.Next() {
		if err = rows.Scan(
			&incomeType.ID,
			&incomeType.ProfileID,
			&incomeType.IncomeTypeID,
			&incomeType.PaymentMethodID,
			&incomeType.PaymentName,
			&incomeType.IncomeAmount,
			&incomeType.Description,
			&incomeType.TimeIncomeHistory,
			&incomeType.ShowTimeIncomeHistory,
			&incomeType.CreatedAt,
			&incomeType.CreatedBy,
			&incomeType.UpdatedAt,
			&incomeType.UpdatedBy,
			&incomeType.DeletedAt,
			&incomeType.DeletedBy,
			&incomeType.IncomeTypeTitle,
			&incomeType.PaymentMethodName,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowsScan, err)
			return nil, err
		}

		incomeTypes = append(incomeTypes, incomeType)
	}

	return &incomeTypes, err
}

func (i *IncomeHistoryRepositoryImpl) GetTotalIncomeByPeriode(ctx context.Context, req *domain.GetIncomeHistoryByTimeAndProfileID) (int, error) {
	query := `SELECT COALESCE(
    					SUM(
    					    CASE WHEN tih.time_income_history BETWEEN $1 AND $2 AND tih.deleted_at IS NULL THEN tih.income_amount ELSE 0 END
    					), 0
    				  )
				FROM t_income_history tih
				WHERE profile_id=$3 AND deleted_at IS NULL`

	conn, err := i.GetConn()
	if err != nil {
		return 0, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return 0, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var result int
	if err = stmt.QueryRowContext(ctx, req.StartTime, req.EndTime, req.ProfileID).Scan(&result); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return 0, err
	}

	return result, nil
}

func (i *IncomeHistoryRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.IncomeHistoryJoin, error) {
	query := `SELECT tih.id, tih.profile_id, tih.income_type_id, tih.payment_method_id, tih.payment_name, tih.income_amount, tih.description, 
       					tih.time_income_history, tih.show_time_income_history, tih.created_at, tih.created_by, tih.updated_at, tih.updated_by, 
       					tih.deleted_at, tih.deleted_by, mit.name, mpm.name
				FROM t_income_history tih
				JOIN m_income_type mit ON tih.income_type_id = mit.id
				LEFT JOIN m_payment_methods mpm ON tih.payment_method_id = mpm.id
				WHERE tih.id=$1 AND tih.profile_id=$2 AND tih.deleted_at IS NULL`

	conn, err := i.GetConn()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var incomeType domain.IncomeHistoryJoin

	if err = stmt.QueryRowContext(ctx, id, profileID).Scan(
		&incomeType.ID,
		&incomeType.ProfileID,
		&incomeType.IncomeTypeID,
		&incomeType.PaymentMethodID,
		&incomeType.PaymentName,
		&incomeType.IncomeAmount,
		&incomeType.Description,
		&incomeType.TimeIncomeHistory,
		&incomeType.ShowTimeIncomeHistory,
		&incomeType.CreatedAt,
		&incomeType.CreatedBy,
		&incomeType.UpdatedAt,
		&incomeType.UpdatedBy,
		&incomeType.DeletedAt,
		&incomeType.DeletedBy,
		&incomeType.IncomeTypeTitle,
		&incomeType.PaymentMethodName,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &incomeType, err
}
