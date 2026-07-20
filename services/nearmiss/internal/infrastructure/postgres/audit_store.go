package postgres

import (
	"context"
	"database/sql"

	auditDomain "prahari/services/nearmiss/internal/domain/audit"
)

// AuditStore implements snapshots logging.
type AuditStore struct {
	db *sql.DB
}

// NewAuditStore instantiates AuditStore.
func NewAuditStore(db *sql.DB) *AuditStore {
	return &AuditStore{db: db}
}

// Log registers mutation.
func (s *AuditStore) Log(ctx context.Context, entry *auditDomain.Entry) error {
	query := `INSERT INTO near_miss_audit (id, entity_type, entity_id, action, actor_id, old_value, new_value, ip_address, user_agent, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := s.db.ExecContext(ctx, query,
		entry.ID, entry.EntityType, entry.EntityID, entry.Action, entry.ActorID,
		entry.OldValue, entry.NewValue, entry.IPAddress, entry.UserAgent, entry.Timestamp,
	)
	return err
}
