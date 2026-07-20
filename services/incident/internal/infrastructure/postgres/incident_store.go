package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	incidentDomain "prahari/services/incident/internal/domain/incident"
)

// IncidentStore implements the incident repository port against PostgreSQL.
type IncidentStore struct {
	db *sql.DB
}

// NewIncidentStore constructs an IncidentStore.
func NewIncidentStore(db *sql.DB) *IncidentStore {
	return &IncidentStore{db: db}
}

// Create persists a new incident aggregate to PostgreSQL.
func (s *IncidentStore) Create(ctx context.Context, inc *incidentDomain.Incident) error {
	query := `
		INSERT INTO incidents (
			id, incident_number, title, description, type, category_id,
			severity_level, priority_level, status_code, reporter_id,
			assignee_id, department_id, location_id, location_detail,
			occurred_at, reported_at, created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)`

	_, err := s.db.ExecContext(ctx, query,
		inc.ID, inc.IncidentNumber, inc.Title, inc.Description, inc.Type,
		inc.CategoryID, inc.SeverityLevel, inc.PriorityLevel, inc.StatusCode,
		inc.ReporterID, inc.AssigneeID, inc.DepartmentID, inc.LocationID,
		inc.LocationDetail, inc.OccurredAt, inc.ReportedAt, inc.CreatedAt, inc.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert incident: %w", err)
	}
	return nil
}

// FindByID retrieves an incident by its unique identifier.
func (s *IncidentStore) FindByID(ctx context.Context, id string) (*incidentDomain.Incident, error) {
	query := `
		SELECT id, incident_number, title, description, type, category_id,
			severity_level, priority_level, status_code, reporter_id,
			assignee_id, department_id, location_id, location_detail,
			occurred_at, reported_at, resolved_at, closed_at, created_at, updated_at
		FROM incidents WHERE id = $1 AND is_deleted = false`

	inc := &incidentDomain.Incident{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&inc.ID, &inc.IncidentNumber, &inc.Title, &inc.Description, &inc.Type,
		&inc.CategoryID, &inc.SeverityLevel, &inc.PriorityLevel, &inc.StatusCode,
		&inc.ReporterID, &inc.AssigneeID, &inc.DepartmentID, &inc.LocationID,
		&inc.LocationDetail, &inc.OccurredAt, &inc.ReportedAt, &inc.ResolvedAt,
		&inc.ClosedAt, &inc.CreatedAt, &inc.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("incident not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query incident: %w", err)
	}
	return inc, nil
}

// FindByNumber retrieves an incident by its human-readable incident number.
func (s *IncidentStore) FindByNumber(ctx context.Context, number string) (*incidentDomain.Incident, error) {
	query := `
		SELECT id, incident_number, title, description, type, category_id,
			severity_level, priority_level, status_code, reporter_id,
			assignee_id, department_id, location_id, location_detail,
			occurred_at, reported_at, resolved_at, closed_at, created_at, updated_at
		FROM incidents WHERE incident_number = $1 AND is_deleted = false`

	inc := &incidentDomain.Incident{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&inc.ID, &inc.IncidentNumber, &inc.Title, &inc.Description, &inc.Type,
		&inc.CategoryID, &inc.SeverityLevel, &inc.PriorityLevel, &inc.StatusCode,
		&inc.ReporterID, &inc.AssigneeID, &inc.DepartmentID, &inc.LocationID,
		&inc.LocationDetail, &inc.OccurredAt, &inc.ReportedAt, &inc.ResolvedAt,
		&inc.ClosedAt, &inc.CreatedAt, &inc.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("incident not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query incident by number: %w", err)
	}
	return inc, nil
}

// Update persists modifications to an existing incident.
func (s *IncidentStore) Update(ctx context.Context, inc *incidentDomain.Incident) error {
	query := `
		UPDATE incidents SET
			title = $2, description = $3, category_id = $4,
			severity_level = $5, priority_level = $6, status_code = $7,
			assignee_id = $8, location_id = $9, location_detail = $10,
			resolved_at = $11, closed_at = $12, updated_at = $13
		WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		inc.ID, inc.Title, inc.Description, inc.CategoryID,
		inc.SeverityLevel, inc.PriorityLevel, inc.StatusCode,
		inc.AssigneeID, inc.LocationID, inc.LocationDetail,
		inc.ResolvedAt, inc.ClosedAt, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update incident: %w", err)
	}
	return nil
}

// Delete performs a soft delete on an incident.
func (s *IncidentStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE incidents SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete incident: %w", err)
	}
	return nil
}

// List retrieves a paginated list of incidents.
func (s *IncidentStore) List(ctx context.Context, offset, limit int) ([]*incidentDomain.Incident, int, error) {
	countQuery := `SELECT COUNT(*) FROM incidents WHERE is_deleted = false`
	var totalCount int
	if err := s.db.QueryRowContext(ctx, countQuery).Scan(&totalCount); err != nil {
		return nil, 0, fmt.Errorf("postgres: failed to count incidents: %w", err)
	}

	query := `
		SELECT id, incident_number, title, description, type, category_id,
			severity_level, priority_level, status_code, reporter_id,
			assignee_id, department_id, location_id, location_detail,
			occurred_at, reported_at, resolved_at, closed_at, created_at, updated_at
		FROM incidents WHERE is_deleted = false
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("postgres: failed to list incidents: %w", err)
	}
	defer rows.Close()

	var incidents []*incidentDomain.Incident
	for rows.Next() {
		inc := &incidentDomain.Incident{}
		if err := rows.Scan(
			&inc.ID, &inc.IncidentNumber, &inc.Title, &inc.Description, &inc.Type,
			&inc.CategoryID, &inc.SeverityLevel, &inc.PriorityLevel, &inc.StatusCode,
			&inc.ReporterID, &inc.AssigneeID, &inc.DepartmentID, &inc.LocationID,
			&inc.LocationDetail, &inc.OccurredAt, &inc.ReportedAt, &inc.ResolvedAt,
			&inc.ClosedAt, &inc.CreatedAt, &inc.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("postgres: failed to scan incident row: %w", err)
		}
		incidents = append(incidents, inc)
	}

	return incidents, totalCount, nil
}

// UpdateAssignee updates the incident's assignee ID.
func (s *IncidentStore) UpdateAssignee(ctx context.Context, incidentID, assigneeID string) error {
	query := `UPDATE incidents SET assignee_id = $2, updated_at = $3 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, incidentID, assigneeID, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to update incident assignee: %w", err)
	}
	return nil
}

// UpdateStatus updates the incident's status code.
func (s *IncidentStore) UpdateStatus(ctx context.Context, incidentID, statusCode string) error {
	query := `UPDATE incidents SET status_code = $2, updated_at = $3 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, incidentID, statusCode, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to update incident status: %w", err)
	}
	return nil
}
