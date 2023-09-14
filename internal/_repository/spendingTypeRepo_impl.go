package _repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/DueIt-Jasanya-Aturuang/one-piece/domain"
	"github.com/DueIt-Jasanya-Aturuang/one-piece/util"
)

type SpendingTypeRepositoryImpl struct {
	domain.UnitOfWorkRepository
}

func NewSpendingTypeRepositoryImpl(
	uow domain.UnitOfWorkRepository,
) domain.SpendingTypeRepository {
	return &SpendingTypeRepositoryImpl{
		UnitOfWorkRepository: uow,
	}
}

func (s *SpendingTypeRepositoryImpl) Create(ctx context.Context, spendingType *domain.SpendingType) error {
	query := `INSERT INTO m_spending_type (id, profile_id, title, maximum_limit, created_at, created_by, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7)`

	tx, err := s.GetTx()
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
		spendingType.ID,
		spendingType.ProfileID,
		spendingType.Title,
		spendingType.MaximumLimit,
		spendingType.CreatedAt,
		spendingType.CreatedBy,
		spendingType.UpdatedAt,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (s *SpendingTypeRepositoryImpl) Update(ctx context.Context, spendingType *domain.SpendingType) error {
	query := `UPDATE m_spending_type SET title = $1, maximum_limit = $2, updated_at = $3, updated_by = $4 
                    WHERE id = $5 AND profile_id = $6 AND deleted_at IS NULL`

	tx, err := s.GetTx()
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
		spendingType.Title,
		spendingType.MaximumLimit,
		spendingType.UpdatedAt,
		spendingType.UpdatedBy,
		spendingType.ID,
		spendingType.ProfileID,
	); err != nil {
		log.Warn().Msgf(util.LogErrExecContext, err)
		return err
	}

	return nil
}

func (s *SpendingTypeRepositoryImpl) Delete(ctx context.Context, id string, profileID string) error {
	query := `UPDATE m_spending_type SET deleted_by = $1, deleted_at = $2 WHERE id = $3 AND profile_id = $4`

	tx, err := s.GetTx()
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

func (s *SpendingTypeRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.SpendingType, error) {
	query := `SELECT id, profile_id, title, maximum_limit, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_spending_type WHERE id = $1 AND deleted_at IS NULL`

	conn, err := s.GetConn()
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

	var spendingType domain.SpendingType
	if err = stmt.QueryRowContext(ctx, id).Scan(
		&spendingType.ID,
		&spendingType.ProfileID,
		&spendingType.Title,
		&spendingType.MaximumLimit,
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

func (s *SpendingTypeRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.SpendingType, error) {
	query := `SELECT id, profile_id, title, maximum_limit, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_spending_type WHERE id = $1 AND profile_id = $2 AND deleted_at IS NULL `

	conn, err := s.GetConn()
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

	var spendingType domain.SpendingType
	if err = stmt.QueryRowContext(ctx, id, profileID).Scan(
		&spendingType.ID,
		&spendingType.ProfileID,
		&spendingType.Title,
		&spendingType.MaximumLimit,
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

func (s *SpendingTypeRepositoryImpl) GetAllByProfileID(ctx context.Context, profileID string) (*[]domain.SpendingType, error) {
	query := `SELECT id, profile_id, title, maximum_limit, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_spending_type WHERE profile_id = $1 AND deleted_at IS NULL `

	conn, err := s.GetConn()
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

	var spendingTypes []domain.SpendingType
	for rows.Next() {
		var spendingType domain.SpendingType
		if err = rows.Scan(
			&spendingType.ID,
			&spendingType.ProfileID,
			&spendingType.Title,
			&spendingType.MaximumLimit,
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
