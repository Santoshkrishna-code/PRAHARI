package postgres

import (
	"context"
	"database/sql"
	"fmt"

	typeDomain "prahari/services/observation/internal/domain/observationtype"
)

// TypeStore implements classifications database logs.
type TypeStore struct {
	db *sql.DB
}

// NewTypeStore instantiates TypeStore.
func NewTypeStore(db *sql.DB) *TypeStore {
	return &TypeStore{db: db}
}

// Create persists classification type.
func (s *TypeStore) Create(ctx context.Context, t *typeDomain.ObservationType) error {
	query := `INSERT INTO observation_types (id, code, name, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, t.ID, t.Code, t.Name, t.Description)
	return err
}

// FindByID returns observation type.
func (s *TypeStore) FindByID(ctx context.Context, id string) (*typeDomain.ObservationType, error) {
	query := `SELECT id, code, name, description FROM observation_types WHERE id = $1`
	t := &typeDomain.ObservationType{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.Code, &t.Name, &t.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("observation type not found: %s", id)
		}
		return nil, err
	}
	return t, nil
}
