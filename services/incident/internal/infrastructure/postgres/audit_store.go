package postgres

import (
	"context"
	"database/sql"
	"fmt"

	auditDomain "prahari/services/incident/internal/domain/audit"
)

// AuditStore implements the append-only audit log persistence adapter.
type AuditStore struct {
	db *sql.DB
}

// NewAuditStore constructs an AuditStore.
func NewAuditStore(db *sql.DB) *AuditStore {
	return &AuditStore{db: db}
}

// Log persists an immutable audit entry.
func (s *AuditStore) Log(ctx context.Context, entry *auditDomain.Entry) error {
	query := `INSERT INTO incident_audit (id, entity_type, entity_id, action, actor_id, old_value, new_value, ip_address, user_agent, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := s.db.ExecContext(ctx, query,
		entry.ID, entry.EntityType, entry.EntityID, entry.Action, entry.ActorID,
		entry.OldValue, entry.NewValue, entry.IPAddress, entry.UserAgent, entry.Timestamp)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert audit entry: %w", err)
	}
	return nil
}

// FindByEntityID retrieves all audit entries for a given entity.
func (s *AuditStore) FindByEntityID(ctx context.Context, entityType, entityID string) ([]*auditDomain.Entry, error) {
	query := `SELECT id, entity_type, entity_id, action, actor_id, old_value, new_value, ip_address, user_agent, timestamp
		FROM incident_audit WHERE entity_type = $1 AND entity_id = $2 ORDER BY timestamp DESC`
	rows, err := s.db.QueryContext(ctx, query, entityType, entityID)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to list audit entries: %w", err)
	}
	defer rows.Close()

	var entries []*auditDomain.Entry
	for rows.Next() {
		e := &auditDomain.Entry{}
		if err := rows.Scan(&e.ID, &e.EntityType, &e.EntityID, &e.Action, &e.ActorID,
			&e.OldValue, &e.NewValue, &e.IPAddress, &e.UserAgent, &e.Timestamp); err != nil {
			return nil, fmt.Errorf("postgres: failed to scan audit entry: %w", err)
		}
		entries = append(entries, e)
	}
	return entries, nil
}
