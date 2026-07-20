package postgres

import (
	"context"
	"database/sql"
	"fmt"

	catDomain "prahari/services/hazard/internal/domain/category"
)

// CategoryStore implements category definitions operations.
type CategoryStore struct {
	db *sql.DB
}

// NewCategoryStore instantiates CategoryStore.
func NewCategoryStore(db *sql.DB) *CategoryStore {
	return &CategoryStore{db: db}
}

// Create persists category.
func (s *CategoryStore) Create(ctx context.Context, c *catDomain.Category) error {
	query := `INSERT INTO hazard_categories (id, name, description) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.Name, c.Description)
	return err
}

// FindByID returns category.
func (s *CategoryStore) FindByID(ctx context.Context, id string) (*catDomain.Category, error) {
	query := `SELECT id, name, description FROM hazard_categories WHERE id = $1`
	c := &catDomain.Category{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category not found: %s", id)
		}
		return nil, err
	}
	return c, nil
}
