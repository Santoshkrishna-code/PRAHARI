package postgres

import (
	"context"
	"database/sql"
	"fmt"

	typeDomain "prahari/services/hazard/internal/domain/hazardtype"
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
func (s *TypeStore) Create(ctx context.Context, t *typeDomain.HazardType) error {
	query := `INSERT INTO hazard_types (id, code, name, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, t.ID, t.Code, t.Name, t.Description)
	return err
}

// FindByID returns hazard type.
func (s *TypeStore) FindByID(ctx context.Context, id string) (*typeDomain.HazardType, error) {
	query := `SELECT id, code, name, description FROM hazard_types WHERE id = $1`
	t := &typeDomain.HazardType{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.Code, &t.Name, &t.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("hazard type not found: %s", id)
		}
		return nil, err
	}
	return t, nil
}
