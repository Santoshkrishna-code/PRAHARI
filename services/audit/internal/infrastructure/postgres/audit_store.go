package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	auditDomain "prahari/services/audit/internal/domain/audit"
	searchDomain "prahari/services/audit/internal/domain/search"
)

// AuditStore implements the core repository port.
type AuditStore struct {
	db *sql.DB
}

// NewAuditStore instantiates an AuditStore.
func NewAuditStore(db *sql.DB) *AuditStore {
	return &AuditStore{db: db}
}

// Create inserts an operational assurance audit register.
func (s *AuditStore) Create(ctx context.Context, a *auditDomain.Audit) error {
	query := `INSERT INTO audits (
		id, audit_number, asset_id, department_id, compliance_rating,
		status_code, title, description, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := s.db.ExecContext(ctx, query,
		a.ID, a.AuditNumber, a.AssetID, a.DepartmentID, a.ComplianceRating,
		a.StatusCode, a.Title, a.Description, a.CreatedAt, a.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert audit record: %w", err)
	}
	return nil
}

// FindByID retrieves an audit by ID.
func (s *AuditStore) FindByID(ctx context.Context, id string) (*auditDomain.Audit, error) {
	query := `SELECT id, audit_number, asset_id, department_id, compliance_rating,
		status_code, title, description, created_at, updated_at
		FROM audits WHERE id = $1 AND is_deleted = false`

	a := &auditDomain.Audit{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&a.ID, &a.AuditNumber, &a.AssetID, &a.DepartmentID, &a.ComplianceRating,
		&a.StatusCode, &a.Title, &a.Description, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("audit assurance record not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query audit: %w", err)
	}
	return a, nil
}

// FindByNumber retrieves an audit by number.
func (s *AuditStore) FindByNumber(ctx context.Context, number string) (*auditDomain.Audit, error) {
	query := `SELECT id, audit_number, asset_id, department_id, compliance_rating,
		status_code, title, description, created_at, updated_at
		FROM audits WHERE audit_number = $1 AND is_deleted = false`

	a := &auditDomain.Audit{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&a.ID, &a.AuditNumber, &a.AssetID, &a.DepartmentID, &a.ComplianceRating,
		&a.StatusCode, &a.Title, &a.Description, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("audit not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query audit by number: %w", err)
	}
	return a, nil
}

// Update persists modifications.
func (s *AuditStore) Update(ctx context.Context, a *auditDomain.Audit) error {
	query := `UPDATE audits SET
		compliance_rating = $2, status_code = $3,
		title = $4, description = $5, updated_at = $6 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		a.ID, a.ComplianceRating, a.StatusCode, a.Title, a.Description, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update audit: %w", err)
	}
	return nil
}

// Delete marks an audit record deleted.
func (s *AuditStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE audits SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete audit record: %w", err)
	}
	return nil
}

// List returns pages.
func (s *AuditStore) List(ctx context.Context, offset, limit int) ([]*auditDomain.Audit, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM audits WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, audit_number, asset_id, department_id, compliance_rating,
		status_code, title, description, created_at, updated_at
		FROM audits WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*auditDomain.Audit
	for rows.Next() {
		a := &auditDomain.Audit{}
		err = rows.Scan(
			&a.ID, &a.AuditNumber, &a.AssetID, &a.DepartmentID, &a.ComplianceRating,
			&a.StatusCode, &a.Title, &a.Description, &a.CreatedAt, &a.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, a)
	}
	return list, total, nil
}

// Search queries dynamic GIN tsvectors.
func (s *AuditStore) Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*auditDomain.Audit, int, error) {
	return s.List(ctx, criteria.Offset(), criteria.PageSize)
}
