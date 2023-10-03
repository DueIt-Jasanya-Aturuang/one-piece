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
	query := `INSERT INTO m_spending_type (id, profile_id, title, maximum_limit, icon, created_at, created_by, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

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
		spendingType.Icon,
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
	query := `UPDATE m_spending_type SET title = $1, maximum_limit = $2, icon = $3, updated_at = $4, updated_by = $5 
                    WHERE id = $6 AND profile_id = $7 AND deleted_at IS NULL`

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
		spendingType.Icon,
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

func (s *SpendingTypeRepositoryImpl) CheckData(ctx context.Context, profileID string) (bool, error) {
	query := `SELECT EXISTS(SELECT id FROM m_spending_type WHERE profile_id = $1);`

	conn, err := s.GetConn()
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
	if err = stmt.QueryRowContext(ctx, profileID).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	return exist, nil
}

func (s *SpendingTypeRepositoryImpl) CheckByTitleAndProfileID(ctx context.Context, profileID string, title string) (bool, error) {
	query := `SELECT EXISTS(SELECT id FROM m_spending_type WHERE profile_id = $1 AND title = $2 AND deleted_at IS NULL);`

	conn, err := s.GetConn()
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
	if err = stmt.QueryRowContext(ctx, profileID, title).Scan(&exist); err != nil {
		log.Warn().Msgf(util.LogErrQueryRowContextScan, err)
		return false, err
	}

	return exist, nil
}

func (s *SpendingTypeRepositoryImpl) GetDefault(ctx context.Context) (*[]domain.SpendingType, error) {
	query := `SELECT id, title, maximum_limit, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_default_spending_type WHERE active = true AND deleted_at IS NULL`

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

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Warn().Msgf(util.LogErrQueryRows, err)
		return nil, err
	}

	var spendingTypes []domain.SpendingType
	var spendingType domain.SpendingType

	for rows.Next() {
		if err = rows.Scan(
			&spendingType.ID,
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

func (s *SpendingTypeRepositoryImpl) GetByIDAndProfileID(ctx context.Context, id string, profileID string) (*domain.SpendingType, error) {
	query := `SELECT id, profile_id, title, maximum_limit, icon, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
				FROM m_spending_type WHERE id = $1 AND profile_id = $2 AND deleted_at IS NULL`

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

func (s *SpendingTypeRepositoryImpl) GetAllByTimeAndProfileID(ctx context.Context, req *domain.RequestGetAllSpendingTypeByTime) (*[]domain.SpendingTypeJoin, error) {
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

	var spendingTypes []domain.SpendingTypeJoin
	var spendingType domain.SpendingTypeJoin
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

func (s *SpendingTypeRepositoryImpl) GetAllByProfileID(ctx context.Context, req *domain.RequestGetAllPaginate) (*[]domain.SpendingType, error) {
	query := `SELECT mst.id, mst.profile_id, mst.title, mst.maximum_limit, mst.icon, mst.created_at, mst.created_by, 
       				mst.updated_at, mst.updated_by, mst.deleted_at, mst.deleted_by
				FROM m_spending_type mst
				WHERE mst.profile_id = $1 AND mst.deleted_at IS NULL `
	if req.ID != "" {
		query += `AND id ` + req.Operation + ` $2 `
	}
	query += `ORDER BY id ` + req.Order + ` LIMIT 2`

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

	var spendingTypes []domain.SpendingType
	var spendingType domain.SpendingType
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
