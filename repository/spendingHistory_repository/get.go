package spendingHistory_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (s *SpendingHistoryRepositoryImpl) GetAllByTimeAndProfileID(
	ctx context.Context, req *repository.GetAllSpendingHistoryByFilterWithISD,
) (*[]repository.SpendingHistoryJoinTable, error) {
	query := `SELECT tsh.id, tsh.profile_id, tsh.spending_type_id, tsh.payment_method_id, tsh.payment_name, tsh.before_balance, 
       				tsh.spending_amount, tsh.after_balance, tsh.description, tsh.time_spending_history, tsh.show_time_spending_history, 
       				tsh.created_at, tsh.created_by, tsh.updated_at, tsh.updated_by, tsh.deleted_at, tsh.deleted_by,
       				mst.title, mpm.name
				FROM t_spending_history tsh 
				JOIN m_spending_type mst ON tsh.spending_type_id = mst.id
				LEFT JOIN m_payment_methods mpm ON tsh.payment_method_id = mpm.id
				WHERE tsh.profile_id = $1 AND tsh.time_spending_history BETWEEN $2 AND $3 AND tsh.deleted_at IS NULL `

	if req.ID != "" {
		query += `AND tsh.id ` + req.Operation + ` $4 `
	}
	query += `ORDER BY tsh.id ` + req.Order + ` LIMIT 5`

	db, err := s.GetDB()
	if err != nil {
		return nil, err
	}
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
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

	var spendingHistories []repository.SpendingHistoryJoinTable
	var spendingHistory repository.SpendingHistoryJoinTable

	for rows.Next() {
		if err = rows.Scan(
			&spendingHistory.ID,
			&spendingHistory.ProfileID,
			&spendingHistory.SpendingTypeID,
			&spendingHistory.PaymentMethodID,
			&spendingHistory.PaymentName,
			&spendingHistory.BeforeBalance,
			&spendingHistory.SpendingAmount,
			&spendingHistory.AfterBalance,
			&spendingHistory.Description,
			&spendingHistory.TimeSpendingHistory,
			&spendingHistory.ShowTimeSpendingHistory,
			&spendingHistory.CreatedAt,
			&spendingHistory.CreatedBy,
			&spendingHistory.UpdatedAt,
			&spendingHistory.UpdatedBy,
			&spendingHistory.DeletedAt,
			&spendingHistory.DeletedBy,
			&spendingHistory.SpendingTypeTitle,
			&spendingHistory.PaymentMethodName,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowsScan, err)
			return nil, err
		}

		spendingHistories = append(spendingHistories, spendingHistory)
	}

	return &spendingHistories, nil
}

func (s *SpendingHistoryRepositoryImpl) GetAllAmountByTimeAndProfileID(ctx context.Context, req *repository.GetTotalSpendingHistoryByPeriode) (int, error) {
	query := `SELECT COALESCE(SUM(CASE WHEN time_spending_history BETWEEN $2 AND $3 AND deleted_at IS NULL THEN spending_amount ELSE 0 END), 0)
				FROM t_spending_history
				WHERE profile_id = $1 AND time_spending_history BETWEEN $2 AND $3 AND deleted_at IS NULL`

	db, err := s.GetDB()
	if err != nil {
		return 0, err
	}
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		return 0, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var res int
	if err = stmt.QueryRowContext(ctx, req.ProfileID, req.StartTime, req.EndTime).Scan(&res); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return 0, err
	}

	return res, nil
}

func (s *SpendingHistoryRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*repository.SpendingHistoryJoinTable, error) {
	query := `SELECT tsh.id, tsh.profile_id, tsh.spending_type_id, tsh.payment_method_id, tsh.payment_name, tsh.before_balance, 
       				tsh.spending_amount, tsh.after_balance, tsh.description, tsh.time_spending_history, tsh.show_time_spending_history, 
       				tsh.created_at, tsh.created_by, tsh.updated_at, tsh.updated_by, tsh.deleted_at, tsh.deleted_by,
       				mst.title, mpm.name
				FROM t_spending_history tsh 
				JOIN m_spending_type mst ON tsh.spending_type_id = mst.id
				LEFT JOIN m_payment_methods mpm ON tsh.payment_method_id = mpm.id
				WHERE tsh.profile_id = $1 AND tsh.id = $2 AND tsh.deleted_at IS NULL`

	db, err := s.GetDB()
	if err != nil {
		return nil, err
	}

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		return nil, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var spendingHistory repository.SpendingHistoryJoinTable
	if err = stmt.QueryRowContext(ctx, profileID, id).Scan(
		&spendingHistory.ID,
		&spendingHistory.ProfileID,
		&spendingHistory.SpendingTypeID,
		&spendingHistory.PaymentMethodID,
		&spendingHistory.PaymentName,
		&spendingHistory.BeforeBalance,
		&spendingHistory.SpendingAmount,
		&spendingHistory.AfterBalance,
		&spendingHistory.Description,
		&spendingHistory.TimeSpendingHistory,
		&spendingHistory.ShowTimeSpendingHistory,
		&spendingHistory.CreatedAt,
		&spendingHistory.CreatedBy,
		&spendingHistory.UpdatedAt,
		&spendingHistory.UpdatedBy,
		&spendingHistory.DeletedAt,
		&spendingHistory.DeletedBy,
		&spendingHistory.SpendingTypeTitle,
		&spendingHistory.PaymentMethodName,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &spendingHistory, nil
}
