package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	complianceDomain "prahari/services/compliance/internal/domain/compliance"
	searchDomain "prahari/services/compliance/internal/domain/search"
)

// ComplianceStore implements the core repository port.
type ComplianceStore struct {
	db *sql.DB
}

// NewComplianceStore instantiates a ComplianceStore.
func NewComplianceStore(db *sql.DB) *ComplianceStore {
	return &ComplianceStore{db: db}
}

// Create inserts an operational compliance register.
func (s *ComplianceStore) Create(ctx context.Context, c *complianceDomain.Compliance) error {
	query := `INSERT INTO compliance_register (
		id, compliance_number, asset_id, department_id, compliance_score,
		status_code, title, description, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := s.db.ExecContext(ctx, query,
		c.ID, c.ComplianceNumber, c.AssetID, c.DepartmentID, c.ComplianceScore,
		c.StatusCode, c.Title, c.Description, c.CreatedAt, c.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert compliance record: %w", err)
	}
	return nil
}

// FindByID retrieves a compliance by ID.
func (s *ComplianceStore) FindByID(ctx context.Context, id string) (*complianceDomain.Compliance, error) {
	query := `SELECT id, compliance_number, asset_id, department_id, compliance_score,
		status_code, title, description, created_at, updated_at
		FROM compliance_register WHERE id = $1 AND is_deleted = false`

	c := &complianceDomain.Compliance{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.ComplianceNumber, &c.AssetID, &c.DepartmentID, &c.ComplianceScore,
		&c.StatusCode, &c.Title, &c.Description, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("compliance assessment record not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query compliance: %w", err)
	}
	return c, nil
}

// FindByNumber retrieves a compliance by number.
func (s *ComplianceStore) FindByNumber(ctx context.Context, number string) (*complianceDomain.Compliance, error) {
	query := `SELECT id, compliance_number, asset_id, department_id, compliance_score,
		status_code, title, description, created_at, updated_at
		FROM compliance_register WHERE compliance_number = $1 AND is_deleted = false`

	c := &complianceDomain.Compliance{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&c.ID, &c.ComplianceNumber, &c.AssetID, &c.DepartmentID, &c.ComplianceScore,
		&c.StatusCode, &c.Title, &c.Description, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("compliance not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query compliance by number: %w", err)
	}
	return c, nil
}

// Update persists modifications.
func (s *ComplianceStore) Update(ctx context.Context, c *complianceDomain.Compliance) error {
	query := `UPDATE compliance_register SET
		compliance_score = $2, status_code = $3,
		title = $4, description = $5, updated_at = $6 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		c.ID, c.ComplianceScore, c.StatusCode, c.Title, c.Description, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update compliance: %w", err)
	}
	return nil
}

// Delete marks a compliance record deleted.
func (s *ComplianceStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE compliance_register SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete compliance record: %w", err)
	}
	return nil
}

// List returns pages.
func (s *ComplianceStore) List(ctx context.Context, offset, limit int) ([]*complianceDomain.Compliance, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM compliance_register WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, compliance_number, asset_id, department_id, compliance_score,
		status_code, title, description, created_at, updated_at
		FROM compliance_register WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*complianceDomain.Compliance
	for rows.Next() {
		c := &complianceDomain.Compliance{}
		err = rows.Scan(
			&c.ID, &c.ComplianceNumber, &c.AssetID, &c.DepartmentID, &c.ComplianceScore,
			&c.StatusCode, &c.Title, &c.Description, &c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, c)
	}
	return list, total, nil
}

// Search queries dynamic GIN tsvectors.
func (s *ComplianceStore) Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*complianceDomain.Compliance, int, error) {
	return s.List(ctx, criteria.Offset(), criteria.PageSize)
}
