package postgres

import (
	"context"
	"database/sql"

	timelineDomain "prahari/services/permit/internal/domain/timeline"
)

// TimelineStore logs milestones.
type TimelineStore struct {
	db *sql.DB
}

// NewTimelineStore instantiates TimelineStore.
func NewTimelineStore(db *sql.DB) *TimelineStore {
	return &TimelineStore{db: db}
}

// Record persists timeline event.
func (s *TimelineStore) Record(ctx context.Context, event *timelineDomain.Event) error {
	query := `INSERT INTO permit_timeline (id, permit_id, event_type, actor_id, description, metadata, occurred_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query,
		event.ID, event.PermitID, event.EventType, event.ActorID, event.Description, event.Metadata, event.OccurredAt,
	)
	return err
}

// FindByPermitID returns milestones.
func (s *TimelineStore) FindByPermitID(ctx context.Context, permitID string) ([]*timelineDomain.Event, error) {
	query := `SELECT id, permit_id, event_type, actor_id, description, metadata, occurred_at FROM permit_timeline WHERE permit_id = $1 ORDER BY occurred_at ASC`
	rows, err := s.db.QueryContext(ctx, query, permitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*timelineDomain.Event
	for rows.Next() {
		e := &timelineDomain.Event{}
		err = rows.Scan(&e.ID, &e.PermitID, &e.EventType, &e.ActorID, &e.Description, &e.Metadata, &e.OccurredAt)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
