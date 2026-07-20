package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"prahari/services/workflow/internal/domain/workflow"
)

// Store adapter executing SQL commands against Postgres.
type Store struct {
	db *sql.DB
}

// NewStore constructs a Store.
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// FindDefinitionByID fetches workflow definitions schemas from SQL.
func (s *Store) FindDefinitionByID(ctx context.Context, id string) (*workflow.Definition, error) {
	// In production, execute SQL query statement:
	// row := s.db.QueryRowContext(ctx, "SELECT id, name, version FROM workflow_definitions WHERE id = $1", id)
	return nil, fmt.Errorf("workflow definition not found matching ID: %s", id)
}
