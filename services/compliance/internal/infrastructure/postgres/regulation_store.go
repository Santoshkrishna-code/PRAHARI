package postgres

import (
	"context"
	"database/sql"
	"fmt"

	regDomain "prahari/services/compliance/internal/domain/regulation"
)

// RegulationStore implements regulatory requirements codes catalog.
type RegulationStore struct {
	db *sql.DB
}

// NewRegulationStore instantiates RegulationStore.
func NewRegulationStore(db *sql.DB) *RegulationStore {
	return &RegulationStore{db: db}
}

// Create persists regulation code.
func (s *RegulationStore) Create(ctx context.Context, r *regDomain.Regulation) error {
	query := `INSERT INTO regulations (id, code, name, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.Code, r.Name, r.Description)
	return err
}

// FindByID returns regulation details.
func (s *RegulationStore) FindByID(ctx context.Context, id string) (*regDomain.Regulation, error) {
	query := `SELECT id, code, name, description FROM regulations WHERE id = $1`
	r := &regDomain.Regulation{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&r.ID, &r.Code, &r.Name, &r.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("regulation details not found: %s", id)
		}
		return nil, err
	}
	return r, nil
}
