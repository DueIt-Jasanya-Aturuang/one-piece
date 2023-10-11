package incomeType_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/repository"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

func (i *IncomeTypeRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*repository.IncomeType, error) {
	query := `SELECT id, profile_id, name, description, icon, income_type, fixed_income, periode, amount, created_at, 
       				created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_income_type WHERE id = $1 AND profile_id = $2 AND deleted_at IS NULL`

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
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var incomeType repository.IncomeType
	if err = stmt.QueryRowContext(ctx, id, profileID).Scan(
		&incomeType.ID,
		&incomeType.ProfileID,
		&incomeType.Name,
		&incomeType.Description,
		&incomeType.Icon,
		&incomeType.IncomeType,
		&incomeType.FixedIncome,
		&incomeType.Periode,
		&incomeType.Amount,
		&incomeType.CreatedAt,
		&incomeType.CreatedBy,
		&incomeType.UpdatedAt,
		&incomeType.UpdatedBy,
		&incomeType.DeletedAt,
		&incomeType.DeletedBy,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		}
		return nil, err
	}

	return &incomeType, nil
}

func (i *IncomeTypeRepositoryImpl) GetAllByProfileID(ctx context.Context, req *repository.GetAllIncomeTypeWithISD) (*[]repository.IncomeType, error) {
	query := `SELECT id, profile_id, name, description, icon, income_type, fixed_income, periode, amount, created_at, 
       				created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_income_type WHERE profile_id = $1 AND deleted_at IS NULL `
	if req.ID != "" {
		query += `AND id ` + req.Operation + ` $2 `
	}
	query += `ORDER BY id ` + req.Order + ` LIMIT 5`

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

	var incomeTypes []repository.IncomeType
	var incomeType repository.IncomeType

	for rows.Next() {
		if err = rows.Scan(
			&incomeType.ID,
			&incomeType.ProfileID,
			&incomeType.Name,
			&incomeType.Description,
			&incomeType.Icon,
			&incomeType.IncomeType,
			&incomeType.FixedIncome,
			&incomeType.Periode,
			&incomeType.Amount,
			&incomeType.CreatedAt,
			&incomeType.CreatedBy,
			&incomeType.UpdatedAt,
			&incomeType.UpdatedBy,
			&incomeType.DeletedAt,
			&incomeType.DeletedBy,
		); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
			}
			return nil, err
		}

		incomeTypes = append(incomeTypes, incomeType)
	}

	return &incomeTypes, nil
}
