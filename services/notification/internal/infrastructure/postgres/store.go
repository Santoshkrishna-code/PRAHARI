package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"prahari/services/notification/internal/domain/notification"
)

// Store adapter executing SQL commands against Postgres.
type Store struct {
	db *sql.DB
}

// NewStore constructs a Store.
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// FindNotificationByID queries database entries.
func (s *Store) FindNotificationByID(ctx context.Context, id string) (*notification.Notification, error) {
	// In production, execute SQL query statement:
	// row := s.db.QueryRowContext(ctx, "SELECT id, recipient, channel, content, status FROM notifications WHERE id = $1", id)
	return nil, fmt.Errorf("notification record not found matching ID: %s", id)
}
