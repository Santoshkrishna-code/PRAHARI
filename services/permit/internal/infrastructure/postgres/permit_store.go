package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	permitDomain "prahari/services/permit/internal/domain/permit"
)

// PermitStore implements the permit repository port using SQL.
type PermitStore struct {
	db *sql.DB
}

// NewPermitStore instantiates a PermitStore.
func NewPermitStore(db *sql.DB) *PermitStore {
	return &PermitStore{db: db}
}

// Create inserts a permit record.
func (s *PermitStore) Create(ctx context.Context, p *permitDomain.Permit) error {
	query := `INSERT INTO permits (
		id, permit_number, title, description, permit_type_id, status_code, risk_level,
		applicant_id, supervisor_id, issuer_id, receiver_id, department_id, contractor_id,
		work_area_id, work_description, planned_start_at, planned_end_at, actual_start_at,
		actual_end_at, valid_until, extension_count, linked_incident_id, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)`

	_, err := s.db.ExecContext(ctx, query,
		p.ID, p.PermitNumber, p.Title, p.Description, p.PermitTypeID, p.StatusCode, p.RiskLevel,
		p.ApplicantID, p.SupervisorID, p.IssuerID, p.ReceiverID, p.DepartmentID, p.ContractorID,
		p.WorkAreaID, p.WorkDescription, p.PlannedStartAt, p.PlannedEndAt, p.ActualStartAt,
		p.ActualEndAt, p.ValidUntil, p.ExtensionCount, p.LinkedIncidentID, p.CreatedAt, p.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to create permit: %w", err)
	}
	return nil
}

// FindByID returns a permit by ID.
func (s *PermitStore) FindByID(ctx context.Context, id string) (*permitDomain.Permit, error) {
	query := `SELECT id, permit_number, title, description, permit_type_id, status_code, risk_level,
		applicant_id, supervisor_id, issuer_id, receiver_id, department_id, contractor_id,
		work_area_id, work_description, planned_start_at, planned_end_at, actual_start_at,
		actual_end_at, valid_until, extension_count, linked_incident_id, created_at, updated_at
		FROM permits WHERE id = $1 AND is_deleted = false`

	p := &permitDomain.Permit{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.PermitNumber, &p.Title, &p.Description, &p.PermitTypeID, &p.StatusCode, &p.RiskLevel,
		&p.ApplicantID, &p.SupervisorID, &p.IssuerID, &p.ReceiverID, &p.DepartmentID, &p.ContractorID,
		&p.WorkAreaID, &p.WorkDescription, &p.PlannedStartAt, &p.PlannedEndAt, &p.ActualStartAt,
		&p.ActualEndAt, &p.ValidUntil, &p.ExtensionCount, &p.LinkedIncidentID, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("permit not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to retrieve permit: %w", err)
	}
	return p, nil
}

// FindByNumber returns a permit by human-readable reference number.
func (s *PermitStore) FindByNumber(ctx context.Context, number string) (*permitDomain.Permit, error) {
	query := `SELECT id, permit_number, title, description, permit_type_id, status_code, risk_level,
		applicant_id, supervisor_id, issuer_id, receiver_id, department_id, contractor_id,
		work_area_id, work_description, planned_start_at, planned_end_at, actual_start_at,
		actual_end_at, valid_until, extension_count, linked_incident_id, created_at, updated_at
		FROM permits WHERE permit_number = $1 AND is_deleted = false`

	p := &permitDomain.Permit{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&p.ID, &p.PermitNumber, &p.Title, &p.Description, &p.PermitTypeID, &p.StatusCode, &p.RiskLevel,
		&p.ApplicantID, &p.SupervisorID, &p.IssuerID, &p.ReceiverID, &p.DepartmentID, &p.ContractorID,
		&p.WorkAreaID, &p.WorkDescription, &p.PlannedStartAt, &p.PlannedEndAt, &p.ActualStartAt,
		&p.ActualEndAt, &p.ValidUntil, &p.ExtensionCount, &p.LinkedIncidentID, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("permit not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to retrieve permit by number: %w", err)
	}
	return p, nil
}

// Update saves edits.
func (s *PermitStore) Update(ctx context.Context, p *permitDomain.Permit) error {
	query := `UPDATE permits SET
		title = $2, description = $3, status_code = $4, risk_level = $5,
		issuer_id = $6, receiver_id = $7, contractor_id = $8, work_description = $9,
		planned_start_at = $10, planned_end_at = $11, actual_start_at = $12, actual_end_at = $13,
		valid_until = $14, extension_count = $15, linked_incident_id = $16, updated_at = $17
		WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		p.ID, p.Title, p.Description, p.StatusCode, p.RiskLevel,
		p.IssuerID, p.ReceiverID, p.ContractorID, p.WorkDescription,
		p.PlannedStartAt, p.PlannedEndAt, p.ActualStartAt, p.ActualEndAt,
		p.ValidUntil, p.ExtensionCount, p.LinkedIncidentID, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update permit: %w", err)
	}
	return nil
}

// UpdateStatus changes status tag values.
func (s *PermitStore) UpdateStatus(ctx context.Context, permitID, status string) error {
	query := `UPDATE permits SET status_code = $2, updated_at = $3 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, permitID, status, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to update status: %w", err)
	}
	return nil
}

// Delete sets deletion flag.
func (s *PermitStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE permits SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete permit: %w", err)
	}
	return nil
}

// List returns a page list.
func (s *PermitStore) List(ctx context.Context, offset, limit int) ([]*permitDomain.Permit, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM permits WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, permit_number, title, description, permit_type_id, status_code, risk_level,
		applicant_id, supervisor_id, issuer_id, receiver_id, department_id, contractor_id,
		work_area_id, work_description, planned_start_at, planned_end_at, actual_start_at,
		actual_end_at, valid_until, extension_count, linked_incident_id, created_at, updated_at
		FROM permits WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var permits []*permitDomain.Permit
	for rows.Next() {
		p := &permitDomain.Permit{}
		err = rows.Scan(
			&p.ID, &p.PermitNumber, &p.Title, &p.Description, &p.PermitTypeID, &p.StatusCode, &p.RiskLevel,
			&p.ApplicantID, &p.SupervisorID, &p.IssuerID, &p.ReceiverID, &p.DepartmentID, &p.ContractorID,
			&p.WorkAreaID, &p.WorkDescription, &p.PlannedStartAt, &p.PlannedEndAt, &p.ActualStartAt,
			&p.ActualEndAt, &p.ValidUntil, &p.ExtensionCount, &p.LinkedIncidentID, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		permits = append(permits, p)
	}
	return permits, total, nil
}
