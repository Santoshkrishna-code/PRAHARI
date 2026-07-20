package postgres

import (
	"context"
	"database/sql"

	trailDomain "prahari/services/audit/internal/domain/audittrail"
)

// TrailStore implements snapshot audit trail logs database.
type TrailStore struct {
	db *sql.DB
}

// NewTrailStore instantiates TrailStore.
func NewTrailStore(db *sql.DB) *TrailStore {
	return &TrailStore{db: db}
}

// Log registers mutation.
func (s *TrailStore) Log(ctx context.Context, entry *trailDomain.Entry) error {
	query := `INSERT INTO audit_trail (id, entity_type, entity_id, action, actor_id, old_value, new_value, ip_address, user_agent, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := s.db.ExecContext(ctx, query,
		entry.ID, entry.EntityType, entry.EntityID, entry.Action, entry.ActorID,
		entry.OldValue, entry.NewValue, entry.IPAddress, entry.UserAgent, entry.Timestamp,
	)
	return err
}
