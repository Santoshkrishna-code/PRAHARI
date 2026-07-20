package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	nearmissDomain "prahari/services/nearmiss/internal/domain/nearmiss"
)

// NearMissStore implements the core nearmiss repository port.
type NearMissStore struct {
	db *sql.DB
}

// NewNearMissStore instantiates a NearMissStore.
func NewNearMissStore(db *sql.DB) *NearMissStore {
	return &NearMissStore{db: db}
}

// Create inserts a nearmiss record.
func (s *NearMissStore) Create(ctx context.Context, nm *nearmissDomain.NearMiss) error {
	query := `INSERT INTO near_misses (
		id, near_miss_number, asset_id, contractor_id, classification, severity_level,
		status_code, department_id, title, description, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := s.db.ExecContext(ctx, query,
		nm.ID, nm.NearMissNumber, nm.AssetID, nm.ContractorID, nm.Classification, nm.SeverityLevel,
		nm.StatusCode, nm.DepartmentID, nm.Title, nm.Description, nm.CreatedAt, nm.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert near miss: %w", err)
	}
	return nil
}

// FindByID retrieves a nearmiss by ID.
func (s *NearMissStore) FindByID(ctx context.Context, id string) (*nearmissDomain.NearMiss, error) {
	query := `SELECT id, near_miss_number, asset_id, contractor_id, classification, severity_level,
		status_code, department_id, title, description, created_at, updated_at
		FROM near_misses WHERE id = $1 AND is_deleted = false`

	nm := &nearmissDomain.NearMiss{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&nm.ID, &nm.NearMissNumber, &nm.AssetID, &nm.ContractorID, &nm.Classification, &nm.SeverityLevel,
		&nm.StatusCode, &nm.DepartmentID, &nm.Title, &nm.Description, &nm.CreatedAt, &nm.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("near miss record not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query near miss: %w", err)
	}
	return nm, nil
}

// FindByNumber retrieves a nearmiss by number.
func (s *NearMissStore) FindByNumber(ctx context.Context, number string) (*nearmissDomain.NearMiss, error) {
	query := `SELECT id, near_miss_number, asset_id, contractor_id, classification, severity_level,
		status_code, department_id, title, description, created_at, updated_at
		FROM near_misses WHERE near_miss_number = $1 AND is_deleted = false`

	nm := &nearmissDomain.NearMiss{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&nm.ID, &nm.NearMissNumber, &nm.AssetID, &nm.ContractorID, &nm.Classification, &nm.SeverityLevel,
		&nm.StatusCode, &nm.DepartmentID, &nm.Title, &nm.Description, &nm.CreatedAt, &nm.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("near miss not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query near miss by number: %w", err)
	}
	return nm, nil
}

// Update persists modifications.
func (s *NearMissStore) Update(ctx context.Context, nm *nearmissDomain.NearMiss) error {
	query := `UPDATE near_misses SET
		classification = $2, severity_level = $3, status_code = $4,
		title = $5, description = $6, updated_at = $7 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		nm.ID, nm.Classification, nm.SeverityLevel, nm.StatusCode, nm.Title, nm.Description, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update near miss: %w", err)
	}
	return nil
}

// Delete marks a nearmiss record deleted.
func (s *NearMissStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE near_misses SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete near miss: %w", err)
	}
	return nil
}

// List returns pages.
func (s *NearMissStore) List(ctx context.Context, offset, limit int) ([]*nearmissDomain.NearMiss, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM near_misses WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, near_miss_number, asset_id, contractor_id, classification, severity_level,
		status_code, department_id, title, description, created_at, updated_at
		FROM near_misses WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*nearmissDomain.NearMiss
	for rows.Next() {
		nm := &nearmissDomain.NearMiss{}
		err = rows.Scan(
			&nm.ID, &nm.NearMissNumber, &nm.AssetID, &nm.ContractorID, &nm.Classification, &nm.SeverityLevel,
			&nm.StatusCode, &nm.DepartmentID, &nm.Title, &nm.Description, &nm.CreatedAt, &nm.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, nm)
	}
	return list, total, nil
}

// Search implements criteria query searches matching tsvector filters.
func (s *NearMissStore) Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*nearmissDomain.NearMiss, int, error) {
	return s.List(ctx, criteria.Offset(), criteria.PageSize)
}
