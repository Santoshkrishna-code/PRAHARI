package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	hazardDomain "prahari/services/hazard/internal/domain/hazard"
)

// HazardStore implements the core hazard repository port.
type HazardStore struct {
	db *sql.DB
}

// NewHazardStore instantiates a HazardStore.
func NewHazardStore(db *sql.DB) *HazardStore {
	return &HazardStore{db: db}
}

// Create inserts a hazard record.
func (s *HazardStore) Create(ctx context.Context, h *hazardDomain.Hazard) error {
	query := `INSERT INTO hazards (
		id, hazard_number, asset_id, contractor_id, hazard_type, initial_risk_score,
		residual_risk_score, status_code, department_id, title, description, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := s.db.ExecContext(ctx, query,
		h.ID, h.HazardNumber, h.AssetID, h.ContractorID, h.HazardType, h.InitialRiskScore,
		h.ResidualRiskScore, h.StatusCode, h.DepartmentID, h.Title, h.Description, h.CreatedAt, h.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert hazard: %w", err)
	}
	return nil
}

// FindByID retrieves a hazard by ID.
func (s *HazardStore) FindByID(ctx context.Context, id string) (*hazardDomain.Hazard, error) {
	query := `SELECT id, hazard_number, asset_id, contractor_id, hazard_type, initial_risk_score,
		residual_risk_score, status_code, department_id, title, description, created_at, updated_at
		FROM hazards WHERE id = $1 AND is_deleted = false`

	h := &hazardDomain.Hazard{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&h.ID, &h.HazardNumber, &h.AssetID, &h.ContractorID, &h.HazardType, &h.InitialRiskScore,
		&h.ResidualRiskScore, &h.StatusCode, &h.DepartmentID, &h.Title, &h.Description, &h.CreatedAt, &h.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("hazard record not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query hazard: %w", err)
	}
	return h, nil
}

// FindByNumber retrieves a hazard by number.
func (s *HazardStore) FindByNumber(ctx context.Context, number string) (*hazardDomain.Hazard, error) {
	query := `SELECT id, hazard_number, asset_id, contractor_id, hazard_type, initial_risk_score,
		residual_risk_score, status_code, department_id, title, description, created_at, updated_at
		FROM hazards WHERE hazard_number = $1 AND is_deleted = false`

	h := &hazardDomain.Hazard{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&h.ID, &h.HazardNumber, &h.AssetID, &h.ContractorID, &h.HazardType, &h.InitialRiskScore,
		&h.ResidualRiskScore, &h.StatusCode, &h.DepartmentID, &h.Title, &h.Description, &h.CreatedAt, &h.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("hazard not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query hazard by number: %w", err)
	}
	return h, nil
}

// Update persists modifications.
func (s *HazardStore) Update(ctx context.Context, h *hazardDomain.Hazard) error {
	query := `UPDATE hazards SET
		initial_risk_score = $2, residual_risk_score = $3, status_code = $4,
		title = $5, description = $6, updated_at = $7 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		h.ID, h.InitialRiskScore, h.ResidualRiskScore, h.StatusCode, h.Title, h.Description, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update hazard: %w", err)
	}
	return nil
}

// Delete marks a hazard record deleted.
func (s *HazardStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE hazards SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete hazard: %w", err)
	}
	return nil
}

// List returns pages.
func (s *HazardStore) List(ctx context.Context, offset, limit int) ([]*hazardDomain.Hazard, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM hazards WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, hazard_number, asset_id, contractor_id, hazard_type, initial_risk_score,
		residual_risk_score, status_code, department_id, title, description, created_at, updated_at
		FROM hazards WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*hazardDomain.Hazard
	for rows.Next() {
		h := &hazardDomain.Hazard{}
		err = rows.Scan(
			&h.ID, &h.HazardNumber, &h.AssetID, &h.ContractorID, &h.HazardType, &h.InitialRiskScore,
			&h.ResidualRiskScore, &h.StatusCode, &h.DepartmentID, &h.Title, &h.Description, &h.CreatedAt, &h.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, h)
	}
	return list, total, nil
}
