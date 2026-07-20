package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	observationDomain "prahari/services/observation/internal/domain/observation"
	searchDomain "prahari/services/observation/internal/domain/search"
)

// ObservationStore implements the core repository port.
type ObservationStore struct {
	db *sql.DB
}

// NewObservationStore instantiates a ObservationStore.
func NewObservationStore(db *sql.DB) *ObservationStore {
	return &ObservationStore{db: db}
}

// Create inserts an observation record.
func (s *ObservationStore) Create(ctx context.Context, o *observationDomain.Observation) error {
	query := `INSERT INTO observations (
		id, observation_number, asset_id, contractor_id, observation_type,
		status_code, department_id, title, description, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := s.db.ExecContext(ctx, query,
		o.ID, o.ObservationNumber, o.AssetID, o.ContractorID, o.ObservationType,
		o.StatusCode, o.DepartmentID, o.Title, o.Description, o.CreatedAt, o.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert observation: %w", err)
	}
	return nil
}

// FindByID retrieves an observation by ID.
func (s *ObservationStore) FindByID(ctx context.Context, id string) (*observationDomain.Observation, error) {
	query := `SELECT id, observation_number, asset_id, contractor_id, observation_type,
		status_code, department_id, title, description, created_at, updated_at
		FROM observations WHERE id = $1 AND is_deleted = false`

	o := &observationDomain.Observation{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&o.ID, &o.ObservationNumber, &o.AssetID, &o.ContractorID, &o.ObservationType,
		&o.StatusCode, &o.DepartmentID, &o.Title, &o.Description, &o.CreatedAt, &o.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("observation record not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query observation: %w", err)
	}
	return o, nil
}

// FindByNumber retrieves an observation by number.
func (s *ObservationStore) FindByNumber(ctx context.Context, number string) (*observationDomain.Observation, error) {
	query := `SELECT id, observation_number, asset_id, contractor_id, observation_type,
		status_code, department_id, title, description, created_at, updated_at
		FROM observations WHERE observation_number = $1 AND is_deleted = false`

	o := &observationDomain.Observation{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&o.ID, &o.ObservationNumber, &o.AssetID, &o.ContractorID, &o.ObservationType,
		&o.StatusCode, &o.DepartmentID, &o.Title, &o.Description, &o.CreatedAt, &o.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("observation not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query observation by number: %w", err)
	}
	return o, nil
}

// Update persists modifications.
func (s *ObservationStore) Update(ctx context.Context, o *observationDomain.Observation) error {
	query := `UPDATE observations SET
		observation_type = $2, status_code = $3,
		title = $4, description = $5, updated_at = $6 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		o.ID, o.ObservationType, o.StatusCode, o.Title, o.Description, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update observation: %w", err)
	}
	return nil
}

// Delete marks an observation record deleted.
func (s *ObservationStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE observations SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete observation: %w", err)
	}
	return nil
}

// List returns pages.
func (s *ObservationStore) List(ctx context.Context, offset, limit int) ([]*observationDomain.Observation, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM observations WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, observation_number, asset_id, contractor_id, observation_type,
		status_code, department_id, title, description, created_at, updated_at
		FROM observations WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*observationDomain.Observation
	for rows.Next() {
		o := &observationDomain.Observation{}
		err = rows.Scan(
			&o.ID, &o.ObservationNumber, &o.AssetID, &o.ContractorID, &o.ObservationType,
			&o.StatusCode, &o.DepartmentID, &o.Title, &o.Description, &o.CreatedAt, &o.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, o)
	}
	return list, total, nil
}

// Search queries indexes.
func (s *ObservationStore) Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*observationDomain.Observation, int, error) {
	return s.List(ctx, criteria.Offset(), criteria.PageSize)
}
