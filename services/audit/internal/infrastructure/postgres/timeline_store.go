package postgres

import (
	"context"
	"database/sql"

	timelineDomain "prahari/services/audit/internal/domain/timeline"
)

// TimelineStore implements milestone logging.
type TimelineStore struct {
	db *sql.DB
}

// NewTimelineStore instantiates TimelineStore.
func NewTimelineStore(db *sql.DB) *TimelineStore {
	return &TimelineStore{db: db}
}

// Record persists timeline event.
func (s *TimelineStore) Record(ctx context.Context, event *timelineDomain.Event) error {
	query := `INSERT INTO timeline (id, audit_id, event_type, actor_id, description, metadata, occurred_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query,
		event.ID, event.AuditID, event.EventType, event.ActorID, event.Description, event.Metadata, event.OccurredAt,
	)
	return err
}

// FindByAuditID returns milestone logs.
func (s *TimelineStore) FindByAuditID(ctx context.Context, auditID string) ([]*timelineDomain.Event, error) {
	query := `SELECT id, audit_id, event_type, actor_id, description, metadata, occurred_at FROM timeline WHERE audit_id = $1 ORDER BY occurred_at ASC`
	rows, err := s.db.QueryContext(ctx, query, auditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*timelineDomain.Event
	for rows.Next() {
		e := &timelineDomain.Event{}
		err = rows.Scan(&e.ID, &e.AuditID, &e.EventType, &e.ActorID, &e.Description, &e.Metadata, &e.OccurredAt)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
