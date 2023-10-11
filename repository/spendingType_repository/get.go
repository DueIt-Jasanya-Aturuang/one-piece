package spendingType_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (s *SpendingTypeRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*repository.SpendingType, error) {
	query := `SELECT id, profile_id, title, maximum_limit, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_spending_type WHERE id = $1 AND profile_id = $2 AND deleted_at IS NULL`

	db, err := s.GetDB()
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
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var spendingType repository.SpendingType
	if err = stmt.QueryRowContext(ctx, id, profileID).Scan(
		&spendingType.ID,
		&spendingType.ProfileID,
		&spendingType.Title,
		&spendingType.MaximumLimit,
		&spendingType.Icon,
		&spendingType.CreatedAt,
		&spendingType.CreatedBy,
		&spendingType.UpdatedAt,
		&spendingType.UpdatedBy,
		&spendingType.DeletedAt,
		&spendingType.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &spendingType, nil
}

func (s *SpendingTypeRepositoryImpl) GetAllByTimeAndProfileID(ctx context.Context, req *repository.GetAllSpendingTypeByFilterWithISD) (*[]repository.SpendingTypeJoinTable, error) {
	query := `SELECT mst.id, mst.profile_id, mst.title, mst.maximum_limit, mst.icon, mst.created_at, mst.created_by, 
       				mst.updated_at, mst.updated_by, mst.deleted_at, mst.deleted_by, 
       				COALESCE(SUM(CASE WHEN tsh.time_spending_history BETWEEN $2 AND $3 AND tsh.deleted_at IS NULL THEN tsh.spending_amount ELSE 0 END), 0)
				FROM m_spending_type mst
				LEFT JOIN t_spending_history tsh ON mst.id = tsh.spending_type_id
				WHERE mst.profile_id = $1 AND mst.deleted_at IS NULL `

	if req.ID != "" {
		query += `AND mst.id ` + req.Operation + ` $4 `
	}

	query += `GROUP BY mst.id ORDER BY mst.id ` + req.Order + ` LIMIT 5`

	db, err := s.GetDB()
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

	var spendingTypes []repository.SpendingTypeJoinTable
	var spendingType repository.SpendingTypeJoinTable
	for rows.Next() {
		if err = rows.Scan(
			&spendingType.ID,
			&spendingType.ProfileID,
			&spendingType.Title,
			&spendingType.MaximumLimit,
			&spendingType.Icon,
			&spendingType.CreatedAt,
			&spendingType.CreatedBy,
			&spendingType.UpdatedAt,
			&spendingType.UpdatedBy,
			&spendingType.DeletedAt,
			&spendingType.DeletedBy,
			&spendingType.Used,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
			return nil, err
		}

		spendingTypes = append(spendingTypes, spendingType)
	}

	return &spendingTypes, nil
}

func (s *SpendingTypeRepositoryImpl) GetAllByProfileID(ctx context.Context, req *repository.GetAllSpendingTypeWithISD) (*[]repository.SpendingType, error) {
	query := `SELECT mst.id, mst.profile_id, mst.title, mst.maximum_limit, mst.icon, mst.created_at, mst.created_by, 
       				mst.updated_at, mst.updated_by, mst.deleted_at, mst.deleted_by
				FROM m_spending_type mst
				WHERE mst.profile_id = $1 AND mst.deleted_at IS NULL `
	if req.ID != "" {
		query += `AND id ` + req.Operation + ` $2 `
	}
	query += `ORDER BY id ` + req.Order + ` LIMIT 5`

	db, err := s.GetDB()
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
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var rows *sql.Rows
	if req.ID != "" {
		rows, err = stmt.QueryContext(ctx, req.ProfileID, req.ID)
	} else {
		rows, err = stmt.QueryContext(ctx, req.ProfileID)
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

	var spendingTypes []repository.SpendingType
	var spendingType repository.SpendingType
	for rows.Next() {
		if err = rows.Scan(
			&spendingType.ID,
			&spendingType.ProfileID,
			&spendingType.Title,
			&spendingType.MaximumLimit,
			&spendingType.Icon,
			&spendingType.CreatedAt,
			&spendingType.CreatedBy,
			&spendingType.UpdatedAt,
			&spendingType.UpdatedBy,
			&spendingType.DeletedAt,
			&spendingType.DeletedBy,
		); err != nil {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
			return nil, err
		}

		spendingTypes = append(spendingTypes, spendingType)
	}

	return &spendingTypes, nil
}
