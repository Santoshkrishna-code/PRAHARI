package postgres

import (
	"context"
	"database/sql"
	"fmt"

	classDomain "prahari/services/nearmiss/internal/domain/classification"
)

// ClassStore implements classifications database logs.
type ClassStore struct {
	db *sql.DB
}

// NewClassStore instantiates ClassStore.
func NewClassStore(db *sql.DB) *ClassStore {
	return &ClassStore{db: db}
}

// Create persists classification type.
func (s *ClassStore) Create(ctx context.Context, c *classDomain.Classification) error {
	query := `INSERT INTO near_miss_classifications (id, code, name, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.Code, c.Name, c.Description)
	return err
}

// FindByID returns classification.
func (s *ClassStore) FindByID(ctx context.Context, id string) (*classDomain.Classification, error) {
	query := `SELECT id, code, name, description FROM near_miss_classifications WHERE id = $1`
	c := &classDomain.Classification{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Code, &c.Name, &c.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("classification not found: %s", id)
		}
		return nil, err
	}
	return c, nil
}
