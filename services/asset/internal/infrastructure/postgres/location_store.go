package postgres

import (
	"context"
	"database/sql"
	"fmt"

	locationDomain "prahari/services/asset/internal/domain/location"
)

// LocationStore implements plant location hierarchy queries.
type LocationStore struct {
	db *sql.DB
}

// NewLocationStore instantiates LocationStore.
func NewLocationStore(db *sql.DB) *LocationStore {
	return &LocationStore{db: db}
}

// Create persists nodes.
func (s *LocationStore) Create(ctx context.Context, l *locationDomain.Location) error {
	query := `INSERT INTO asset_locations (id, parent_id, name, description, is_active)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, l.ID, l.ParentID, l.Name, l.Description, l.IsActive)
	return err
}

// FindByID returns node.
func (s *LocationStore) FindByID(ctx context.Context, id string) (*locationDomain.Location, error) {
	query := `SELECT id, parent_id, name, description, is_active FROM asset_locations WHERE id = $1`
	l := &locationDomain.Location{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&l.ID, &l.ParentID, &l.Name, &l.Description, &l.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("location not found: %s", id)
		}
		return nil, err
	}
	return l, nil
}
