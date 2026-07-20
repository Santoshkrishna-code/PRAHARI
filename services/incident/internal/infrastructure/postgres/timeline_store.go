package postgres

import (
	"context"
	"database/sql"
	"fmt"

	timelineDomain "prahari/services/incident/internal/domain/timeline"
)

// TimelineStore implements the timeline event persistence adapter.
type TimelineStore struct {
	db *sql.DB
}

// NewTimelineStore constructs a TimelineStore.
func NewTimelineStore(db *sql.DB) *TimelineStore {
	return &TimelineStore{db: db}
}

// Record persists an immutable timeline event.
func (s *TimelineStore) Record(ctx context.Context, event *timelineDomain.Event) error {
	query := `INSERT INTO incident_timeline (id, incident_id, event_type, actor_id, description, metadata, occurred_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query,
		event.ID, event.IncidentID, event.EventType, event.ActorID,
		event.Description, event.Metadata, event.OccurredAt)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert timeline event: %w", err)
	}
	return nil
}

// FindByIncidentID retrieves all timeline events for an incident in chronological order.
func (s *TimelineStore) FindByIncidentID(ctx context.Context, incidentID string) ([]*timelineDomain.Event, error) {
	query := `SELECT id, incident_id, event_type, actor_id, description, metadata, occurred_at
		FROM incident_timeline WHERE incident_id = $1 ORDER BY occurred_at ASC`
	rows, err := s.db.QueryContext(ctx, query, incidentID)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to list timeline events: %w", err)
	}
	defer rows.Close()

	var events []*timelineDomain.Event
	for rows.Next() {
		e := &timelineDomain.Event{}
		if err := rows.Scan(&e.ID, &e.IncidentID, &e.EventType, &e.ActorID,
			&e.Description, &e.Metadata, &e.OccurredAt); err != nil {
			return nil, fmt.Errorf("postgres: failed to scan timeline event: %w", err)
		}
		events = append(events, e)
	}
	return events, nil
}
