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

type IncomeTypeRepositoryImpl struct {
	domain.UnitOfWorkRepository
}

func NewIncomeTypeRepositoryImpl(uow domain.UnitOfWorkRepository) domain.IncomeTypeRepository {
	return &IncomeTypeRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}

func (i *IncomeTypeRepositoryImpl) Create(ctx context.Context, income *domain.IncomeType) error {
	query := `INSERT INTO m_income_type (id, profile_id, name, description, icon, income_type, fixed_income, periode,
                             amount, created_at, created_by, updated_at) 
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
		income.Name,
		income.Description,
		income.Icon,
		income.IncomeType,
		income.FixedIncome,
		income.Periode,
		income.Amount,
		income.CreatedAt,
		income.CreatedBy,
		income.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (i *IncomeTypeRepositoryImpl) Update(ctx context.Context, income *domain.IncomeType) error {
	query := `UPDATE m_income_type SET name = $1, description = $2, icon = $3, income_type = $4, fixed_income = $5,
                         periode = $6, amount = $7, updated_at = $8, updated_by = $9 
                    WHERE id = $10 AND profile_id = $11 AND deleted_at IS NULL`

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
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	if _, err = stmt.ExecContext(
		ctx,
		income.Name,
		income.Description,
		income.Icon,
		income.IncomeType,
		income.FixedIncome,
		income.Periode,
		income.Amount,
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

func (i *IncomeTypeRepositoryImpl) Delete(ctx context.Context, id string, profileID string) error {
	query := `UPDATE m_income_type SET deleted_by = $1, deleted_at = $2 WHERE id = $3 AND profile_id = $4`

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
		profileID,
		time.Now().Unix(),
		id,
		profileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (i *IncomeTypeRepositoryImpl) CheckByNameAndProfileID(ctx context.Context, profileID string, name string) (bool, error) {
	query := `SELECT EXISTS(SELECT id FROM m_income_type WHERE profile_id = $1 AND name = $2 AND deleted_at IS NULL);`

	conn, err := i.GetConn()
	if err != nil {
		return false, err
	}

	stmt, err := conn.PrepareContext(ctx, query)
	if err != nil {
		log.Warn().Msgf(util.LogErrPrepareContext, err)
		return false, err
	}
	defer func() {
		if errClose := stmt.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var exist bool
	if err = stmt.QueryRowContext(ctx, profileID, name).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	return exist, nil
}

func (i *IncomeTypeRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.IncomeType, error) {
	query := `SELECT id, profile_id, name, description, icon, income_type, fixed_income, periode, amount, created_at, 
       				created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_income_type WHERE id = $1 AND profile_id = $2 AND deleted_at IS NULL`

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
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	var incomeType domain.IncomeType
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

func (i *IncomeTypeRepositoryImpl) GetAllByProfileID(ctx context.Context, profileID string) (*[]domain.IncomeType, error) {
	query := `SELECT id, profile_id, name, description, icon, income_type, fixed_income, periode, amount, created_at, 
       				created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_income_type WHERE profile_id = $1 AND deleted_at IS NULL`

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
			log.Warn().Msgf(util.LogErrPrepareContextClose, err)
		}
	}()

	rows, err := stmt.QueryContext(ctx, profileID)
	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
		return nil, err
	}
	defer func() {
		if errClose := rows.Close(); errClose != nil {
			log.Warn().Msgf(util.LogErrQueryRowsClose, err)
		}
	}()

	var incomeTypes []domain.IncomeType
	var incomeType domain.IncomeType

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
