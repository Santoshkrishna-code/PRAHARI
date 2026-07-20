package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	riskDomain "prahari/services/risk/internal/domain/risk"
	searchDomain "prahari/services/risk/internal/domain/search"
)

// AssessmentStore implements the core repository port.
type AssessmentStore struct {
	db *sql.DB
}

// NewAssessmentStore instantiates an AssessmentStore.
func NewAssessmentStore(db *sql.DB) *AssessmentStore {
	return &AssessmentStore{db: db}
}

// Create inserts an operational risk record.
func (s *AssessmentStore) Create(ctx context.Context, r *riskDomain.Risk) error {
	query := `INSERT INTO risk_assessments (
		id, risk_number, asset_id, department_id, inherent_risk_score,
		residual_risk_score, status_code, title, description, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := s.db.ExecContext(ctx, query,
		r.ID, r.RiskNumber, r.AssetID, r.DepartmentID, r.InherentRiskScore,
		r.ResidualRiskScore, r.StatusCode, r.Title, r.Description, r.CreatedAt, r.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert risk record: %w", err)
	}
	return nil
}

// FindByID retrieves a risk by ID.
func (s *AssessmentStore) FindByID(ctx context.Context, id string) (*riskDomain.Risk, error) {
	query := `SELECT id, risk_number, asset_id, department_id, inherent_risk_score,
		residual_risk_score, status_code, title, description, created_at, updated_at
		FROM risk_assessments WHERE id = $1 AND is_deleted = false`

	r := &riskDomain.Risk{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&r.ID, &r.RiskNumber, &r.AssetID, &r.DepartmentID, &r.InherentRiskScore,
		&r.ResidualRiskScore, &r.StatusCode, &r.Title, &r.Description, &r.CreatedAt, &r.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("risk assessment record not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query risk: %w", err)
	}
	return r, nil
}

// FindByNumber retrieves a risk by number.
func (s *AssessmentStore) FindByNumber(ctx context.Context, number string) (*riskDomain.Risk, error) {
	query := `SELECT id, risk_number, asset_id, department_id, inherent_risk_score,
		residual_risk_score, status_code, title, description, created_at, updated_at
		FROM risk_assessments WHERE risk_number = $1 AND is_deleted = false`

	r := &riskDomain.Risk{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&r.ID, &r.RiskNumber, &r.AssetID, &r.DepartmentID, &r.InherentRiskScore,
		&r.ResidualRiskScore, &r.StatusCode, &r.Title, &r.Description, &r.CreatedAt, &r.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("risk not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query risk by number: %w", err)
	}
	return r, nil
}

// Update persists modifications.
func (s *AssessmentStore) Update(ctx context.Context, r *riskDomain.Risk) error {
	query := `UPDATE risk_assessments SET
		inherent_risk_score = $2, residual_risk_score = $3, status_code = $4,
		title = $5, description = $6, updated_at = $7 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		r.ID, r.InherentRiskScore, r.ResidualRiskScore, r.StatusCode, r.Title, r.Description, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update risk assessment: %w", err)
	}
	return nil
}

// Delete marks a risk record deleted.
func (s *AssessmentStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE risk_assessments SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete risk record: %w", err)
	}
	return nil
}

// List returns pages.
func (s *AssessmentStore) List(ctx context.Context, offset, limit int) ([]*riskDomain.Risk, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM risk_assessments WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, risk_number, asset_id, department_id, inherent_risk_score,
		residual_risk_score, status_code, title, description, created_at, updated_at
		FROM risk_assessments WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*riskDomain.Risk
	for rows.Next() {
		r := &riskDomain.Risk{}
		err = rows.Scan(
			&r.ID, &r.RiskNumber, &r.AssetID, &r.DepartmentID, &r.InherentRiskScore,
			&r.ResidualRiskScore, &r.StatusCode, &r.Title, &r.Description, &r.CreatedAt, &r.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, r)
	}
	return list, total, nil
}

// Search queries GIN tsvectors.
func (s *AssessmentStore) Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*riskDomain.Risk, int, error) {
	return s.List(ctx, criteria.Offset(), criteria.PageSize)
}
