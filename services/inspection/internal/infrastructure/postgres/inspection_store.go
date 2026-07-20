package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	inspectionDomain "prahari/services/inspection/internal/domain/inspection"
)

// InspectionStore implements the core inspection repository port.
type InspectionStore struct {
	db *sql.DB
}

// NewInspectionStore instantiates an InspectionStore.
func NewInspectionStore(db *sql.DB) *InspectionStore {
	return &InspectionStore{db: db}
}

// Create inserts an inspection record.
func (s *InspectionStore) Create(ctx context.Context, i *inspectionDomain.Inspection) error {
	query := `INSERT INTO inspections (
		id, inspection_number, title, description, inspection_type, status_code,
		schedule_id, inspector_id, department_id, asset_id, linked_permit_id,
		linked_incident_id, compliance_score, started_at, completed_at, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`

	_, err := s.db.ExecContext(ctx, query,
		i.ID, i.InspectionNumber, i.Title, i.Description, i.InspectionType, i.StatusCode,
		i.ScheduleID, i.InspectorID, i.DepartmentID, i.AssetID, i.LinkedPermitID,
		i.LinkedIncidentID, i.ComplianceScore, i.StartedAt, i.CompletedAt, i.CreatedAt, i.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert inspection: %w", err)
	}
	return nil
}

// FindByID retrieves an inspection by ID.
func (s *InspectionStore) FindByID(ctx context.Context, id string) (*inspectionDomain.Inspection, error) {
	query := `SELECT id, inspection_number, title, description, inspection_type, status_code,
		schedule_id, inspector_id, department_id, asset_id, linked_permit_id,
		linked_incident_id, compliance_score, started_at, completed_at, created_at, updated_at
		FROM inspections WHERE id = $1 AND is_deleted = false`

	i := &inspectionDomain.Inspection{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&i.ID, &i.InspectionNumber, &i.Title, &i.Description, &i.InspectionType, &i.StatusCode,
		&i.ScheduleID, &i.InspectorID, &i.DepartmentID, &i.AssetID, &i.LinkedPermitID,
		&i.LinkedIncidentID, &i.ComplianceScore, &i.StartedAt, &i.CompletedAt, &i.CreatedAt, &i.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("inspection not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query inspection: %w", err)
	}
	return i, nil
}

// FindByNumber retrieves an inspection by walkthrough code number.
func (s *InspectionStore) FindByNumber(ctx context.Context, number string) (*inspectionDomain.Inspection, error) {
	query := `SELECT id, inspection_number, title, description, inspection_type, status_code,
		schedule_id, inspector_id, department_id, asset_id, linked_permit_id,
		linked_incident_id, compliance_score, started_at, completed_at, created_at, updated_at
		FROM inspections WHERE inspection_number = $1 AND is_deleted = false`

	i := &inspectionDomain.Inspection{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&i.ID, &i.InspectionNumber, &i.Title, &i.Description, &i.InspectionType, &i.StatusCode,
		&i.ScheduleID, &i.InspectorID, &i.DepartmentID, &i.AssetID, &i.LinkedPermitID,
		&i.LinkedIncidentID, &i.ComplianceScore, &i.StartedAt, &i.CompletedAt, &i.CreatedAt, &i.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("inspection not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query inspection by number: %w", err)
	}
	return i, nil
}

// Update persists edits.
func (s *InspectionStore) Update(ctx context.Context, i *inspectionDomain.Inspection) error {
	query := `UPDATE inspections SET
		title = $2, description = $3, status_code = $4, compliance_score = $5,
		started_at = $6, completed_at = $7, linked_incident_id = $8, updated_at = $9
		WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		i.ID, i.Title, i.Description, i.StatusCode, i.ComplianceScore,
		i.StartedAt, i.CompletedAt, i.LinkedIncidentID, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update inspection: %w", err)
	}
	return nil
}

// Delete soft deletes an inspection.
func (s *InspectionStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE inspections SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete inspection: %w", err)
	}
	return nil
}

// List returns a page list of inspections.
func (s *InspectionStore) List(ctx context.Context, offset, limit int) ([]*inspectionDomain.Inspection, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM inspections WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, inspection_number, title, description, inspection_type, status_code,
		schedule_id, inspector_id, department_id, asset_id, linked_permit_id,
		linked_incident_id, compliance_score, started_at, completed_at, created_at, updated_at
		FROM inspections WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var inspections []*inspectionDomain.Inspection
	for rows.Next() {
		i := &inspectionDomain.Inspection{}
		err = rows.Scan(
			&i.ID, &i.InspectionNumber, &i.Title, &i.Description, &i.InspectionType, &i.StatusCode,
			&i.ScheduleID, &i.InspectorID, &i.DepartmentID, &i.AssetID, &i.LinkedPermitID,
			&i.LinkedIncidentID, &i.ComplianceScore, &i.StartedAt, &i.CompletedAt, &i.CreatedAt, &i.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		inspections = append(inspections, i)
	}
	return inspections, total, nil
}
