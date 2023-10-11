package incomeHistory_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (i *IncomeHistoryRepositoryImpl) GetAllByTimeAndProfileID(
	ctx context.Context, req *repository.GetAllIncomeHistoryByTimeFilterWithISD,
) (*[]repository.IncomeHistoryJoinTable, error) {
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

	db, err := i.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
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

	var incomeTypes []repository.IncomeHistoryJoinTable
	var incomeType repository.IncomeHistoryJoinTable

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

func (i *IncomeHistoryRepositoryImpl) GetTotalAmountByPeriode(ctx context.Context, req *repository.GetTotalIncomeHistoryByPeriode) (int, error) {
	query := `SELECT COALESCE(
    					SUM(
    					    CASE WHEN tih.time_income_history BETWEEN $1 AND $2 AND tih.deleted_at IS NULL THEN tih.income_amount ELSE 0 END
    					), 0
    				  )
				FROM t_income_history tih
				WHERE profile_id=$3 AND deleted_at IS NULL`

	db, err := i.GetDB()
	if err != nil {
		return 0, err
	}

	stmt, err := db.PrepareContext(ctx, query)
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

func (i *IncomeHistoryRepositoryImpl) GetByIDAndProfileID(
	ctx context.Context, id string, profileID string,
) (*repository.IncomeHistoryJoinTable, error) {
	query := `SELECT tih.id, tih.profile_id, tih.income_type_id, tih.payment_method_id, tih.payment_name, tih.income_amount, tih.description, 
       					tih.time_income_history, tih.show_time_income_history, tih.created_at, tih.created_by, tih.updated_at, tih.updated_by, 
       					tih.deleted_at, tih.deleted_by, mit.name, mpm.name
				FROM t_income_history tih
				JOIN m_income_type mit ON tih.income_type_id = mit.id
				LEFT JOIN m_payment_methods mpm ON tih.payment_method_id = mpm.id
				WHERE tih.id=$1 AND tih.profile_id=$2 AND tih.deleted_at IS NULL`

	db, err := i.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, errClose)
		}
	}()

	var incomeType repository.IncomeHistoryJoinTable

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
