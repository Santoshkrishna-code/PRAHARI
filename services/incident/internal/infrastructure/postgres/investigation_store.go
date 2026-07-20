package postgres

import (
	"context"
	"database/sql"
	"fmt"

	investigationDomain "prahari/services/incident/internal/domain/investigation"
	rootcauseDomain "prahari/services/incident/internal/domain/rootcause"
)

// InvestigationStore implements the investigation persistence adapter against PostgreSQL.
type InvestigationStore struct {
	db *sql.DB
}

// NewInvestigationStore constructs an InvestigationStore.
func NewInvestigationStore(db *sql.DB) *InvestigationStore {
	return &InvestigationStore{db: db}
}

// Create persists a new investigation.
func (s *InvestigationStore) Create(ctx context.Context, inv *investigationDomain.Investigation) error {
	query := `INSERT INTO incident_investigations (id, incident_id, investigator_id, methodology, findings, recommendations, status, started_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := s.db.ExecContext(ctx, query, inv.ID, inv.IncidentID, inv.InvestigatorID, inv.Methodology,
		inv.Findings, inv.Recommendations, inv.Status, inv.StartedAt, inv.CreatedAt, inv.UpdatedAt)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert investigation: %w", err)
	}
	return nil
}

// FindByID retrieves an investigation by its unique identifier.
func (s *InvestigationStore) FindByID(ctx context.Context, id string) (*investigationDomain.Investigation, error) {
	query := `SELECT id, incident_id, investigator_id, methodology, findings, recommendations, status, started_at, completed_at, created_at, updated_at
		FROM incident_investigations WHERE id = $1`
	inv := &investigationDomain.Investigation{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&inv.ID, &inv.IncidentID, &inv.InvestigatorID, &inv.Methodology, &inv.Findings,
		&inv.Recommendations, &inv.Status, &inv.StartedAt, &inv.CompletedAt, &inv.CreatedAt, &inv.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to query investigation: %w", err)
	}
	return inv, nil
}

// FindByIncidentID retrieves the investigation for an incident.
func (s *InvestigationStore) FindByIncidentID(ctx context.Context, incidentID string) (*investigationDomain.Investigation, error) {
	query := `SELECT id, incident_id, investigator_id, methodology, findings, recommendations, status, started_at, completed_at, created_at, updated_at
		FROM incident_investigations WHERE incident_id = $1 ORDER BY created_at DESC LIMIT 1`
	inv := &investigationDomain.Investigation{}
	err := s.db.QueryRowContext(ctx, query, incidentID).Scan(
		&inv.ID, &inv.IncidentID, &inv.InvestigatorID, &inv.Methodology, &inv.Findings,
		&inv.Recommendations, &inv.Status, &inv.StartedAt, &inv.CompletedAt, &inv.CreatedAt, &inv.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no investigation found for incident: %s", incidentID)
		}
		return nil, fmt.Errorf("postgres: failed to query investigation: %w", err)
	}
	return inv, nil
}

// Update persists modifications to an investigation.
func (s *InvestigationStore) Update(ctx context.Context, inv *investigationDomain.Investigation) error {
	query := `UPDATE incident_investigations SET findings = $2, recommendations = $3, status = $4, completed_at = $5, updated_at = $6 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, inv.ID, inv.Findings, inv.Recommendations, inv.Status, inv.CompletedAt, inv.UpdatedAt)
	if err != nil {
		return fmt.Errorf("postgres: failed to update investigation: %w", err)
	}
	return nil
}

// RootCauseStore implements the root cause persistence adapter.
type RootCauseStore struct {
	db *sql.DB
}

// NewRootCauseStore constructs a RootCauseStore.
func NewRootCauseStore(db *sql.DB) *RootCauseStore {
	return &RootCauseStore{db: db}
}

// Create persists a new root cause.
func (s *RootCauseStore) Create(ctx context.Context, rc *rootcauseDomain.RootCause) error {
	query := `INSERT INTO incident_root_causes (id, investigation_id, incident_id, category, description, contributing_factors, identified_by, identified_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, rc.ID, rc.InvestigationID, rc.IncidentID, rc.Category,
		rc.Description, rc.ContributingFactors, rc.IdentifiedBy, rc.IdentifiedAt)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert root cause: %w", err)
	}
	return nil
}

// FindByInvestigationID retrieves all root causes for an investigation.
func (s *RootCauseStore) FindByInvestigationID(ctx context.Context, investigationID string) ([]*rootcauseDomain.RootCause, error) {
	query := `SELECT id, investigation_id, incident_id, category, description, contributing_factors, identified_by, identified_at
		FROM incident_root_causes WHERE investigation_id = $1`
	rows, err := s.db.QueryContext(ctx, query, investigationID)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to list root causes: %w", err)
	}
	defer rows.Close()

	var causes []*rootcauseDomain.RootCause
	for rows.Next() {
		rc := &rootcauseDomain.RootCause{}
		if err := rows.Scan(&rc.ID, &rc.InvestigationID, &rc.IncidentID, &rc.Category, &rc.Description,
			&rc.ContributingFactors, &rc.IdentifiedBy, &rc.IdentifiedAt); err != nil {
			return nil, fmt.Errorf("postgres: failed to scan root cause: %w", err)
		}
		causes = append(causes, rc)
	}
	return causes, nil
}
