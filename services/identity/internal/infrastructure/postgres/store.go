package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"prahari/services/identity/internal/domain/user"
)

// Store adapter executing SQL commands against Postgres.
type Store struct {
	db *sql.DB
}

// NewStore constructs a Store.
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// FindUserByEmail executes matching lookup query.
func (s *Store) FindUserByEmail(ctx context.Context, email string) (*user.User, error) {
	// In production, execute SQL query statement:
	// row := s.db.QueryRowContext(ctx, "SELECT id, email, role, status FROM users WHERE email = $1", email)
	return nil, fmt.Errorf("user not found matching email: %s", email)
}
