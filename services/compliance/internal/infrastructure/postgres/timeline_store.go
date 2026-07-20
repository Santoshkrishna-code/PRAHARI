package postgres

import (
	"context"
	"database/sql"

	timelineDomain "prahari/services/compliance/internal/domain/timeline"
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
	query := `INSERT INTO timeline (id, compliance_id, event_type, actor_id, description, metadata, occurred_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query,
		event.ID, event.ComplianceID, event.EventType, event.ActorID, event.Description, event.Metadata, event.OccurredAt,
	)
	return err
}

// FindByComplianceID returns milestone logs.
func (s *TimelineStore) FindByComplianceID(ctx context.Context, complianceID string) ([]*timelineDomain.Event, error) {
	query := `SELECT id, compliance_id, event_type, actor_id, description, metadata, occurred_at FROM timeline WHERE compliance_id = $1 ORDER BY occurred_at ASC`
	rows, err := s.db.QueryContext(ctx, query, complianceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*timelineDomain.Event
	for rows.Next() {
		e := &timelineDomain.Event{}
		err = rows.Scan(&e.ID, &e.ComplianceID, &e.EventType, &e.ActorID, &e.Description, &e.Metadata, &e.OccurredAt)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
