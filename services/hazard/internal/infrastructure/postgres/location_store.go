package postgres

import (
	"context"
	"database/sql"
	"fmt"

	locDomain "prahari/services/hazard/internal/domain/location"
)

// LocationStore implements location coordinates mappings persistence.
type LocationStore struct {
	db *sql.DB
}

// NewLocationStore instantiates LocationStore.
func NewLocationStore(db *sql.DB) *LocationStore {
	return &LocationStore{db: db}
}

// Create persists location mappings.
func (s *LocationStore) Create(ctx context.Context, l *locDomain.Location) error {
	query := `INSERT INTO hazard_locations (id, name, facility_id) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, l.ID, l.Name, l.FacilityID)
	return err
}

// FindByID returns location mapping details.
func (s *LocationStore) FindByID(ctx context.Context, id string) (*locDomain.Location, error) {
	query := `SELECT id, name, facility_id FROM hazard_locations WHERE id = $1`
	l := &locDomain.Location{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&l.ID, &l.Name, &l.FacilityID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("location not found: %s", id)
		}
		return nil, err
	}
	return l, nil
}
