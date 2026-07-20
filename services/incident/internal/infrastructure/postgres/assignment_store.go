package postgres

import (
	"context"
	"database/sql"
	"fmt"

	assignmentDomain "prahari/services/incident/internal/domain/assignment"
)

// AssignmentStore implements the assignment persistence adapter against PostgreSQL.
type AssignmentStore struct {
	db *sql.DB
}

// NewAssignmentStore constructs an AssignmentStore.
func NewAssignmentStore(db *sql.DB) *AssignmentStore {
	return &AssignmentStore{db: db}
}

// Create persists a new assignment.
func (s *AssignmentStore) Create(ctx context.Context, a *assignmentDomain.Assignment) error {
	query := `INSERT INTO incident_assignments (id, incident_id, assignee_id, assigner_id, role, assigned_at, is_active, note)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.IncidentID, a.AssigneeID, a.AssignerID, a.Role, a.AssignedAt, a.IsActive, a.Note)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert assignment: %w", err)
	}
	return nil
}

// FindByID retrieves an assignment by its unique identifier.
func (s *AssignmentStore) FindByID(ctx context.Context, id string) (*assignmentDomain.Assignment, error) {
	query := `SELECT id, incident_id, assignee_id, assigner_id, role, assigned_at, accepted_at, is_active, note
		FROM incident_assignments WHERE id = $1`
	a := &assignmentDomain.Assignment{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&a.ID, &a.IncidentID, &a.AssigneeID, &a.AssignerID, &a.Role, &a.AssignedAt, &a.AcceptedAt, &a.IsActive, &a.Note)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to query assignment: %w", err)
	}
	return a, nil
}

// FindByIncidentID retrieves all assignments for an incident.
func (s *AssignmentStore) FindByIncidentID(ctx context.Context, incidentID string) ([]*assignmentDomain.Assignment, error) {
	query := `SELECT id, incident_id, assignee_id, assigner_id, role, assigned_at, accepted_at, is_active, note
		FROM incident_assignments WHERE incident_id = $1 ORDER BY assigned_at DESC`
	rows, err := s.db.QueryContext(ctx, query, incidentID)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to list assignments: %w", err)
	}
	defer rows.Close()

	var assignments []*assignmentDomain.Assignment
	for rows.Next() {
		a := &assignmentDomain.Assignment{}
		if err := rows.Scan(&a.ID, &a.IncidentID, &a.AssigneeID, &a.AssignerID, &a.Role, &a.AssignedAt, &a.AcceptedAt, &a.IsActive, &a.Note); err != nil {
			return nil, fmt.Errorf("postgres: failed to scan assignment: %w", err)
		}
		assignments = append(assignments, a)
	}
	return assignments, nil
}

// Update persists modifications to an existing assignment.
func (s *AssignmentStore) Update(ctx context.Context, a *assignmentDomain.Assignment) error {
	query := `UPDATE incident_assignments SET accepted_at = $2, is_active = $3, note = $4 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.AcceptedAt, a.IsActive, a.Note)
	if err != nil {
		return fmt.Errorf("postgres: failed to update assignment: %w", err)
	}
	return nil
}
