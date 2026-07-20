package postgres

import (
	"context"
	"database/sql"
	"fmt"

	categoryDomain "prahari/services/asset/internal/domain/assetcategory"
)

// CategoryStore implements taxonomies.
type CategoryStore struct {
	db *sql.DB
}

// NewCategoryStore instantiates a CategoryStore.
func NewCategoryStore(db *sql.DB) *CategoryStore {
	return &CategoryStore{db: db}
}

// Create inserts category.
func (s *CategoryStore) Create(ctx context.Context, ac *categoryDomain.AssetCategory) error {
	query := `INSERT INTO asset_categories (id, name, description, is_active) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, ac.ID, ac.Name, ac.Description, ac.IsActive)
	return err
}

// FindByID returns category.
func (s *CategoryStore) FindByID(ctx context.Context, id string) (*categoryDomain.AssetCategory, error) {
	query := `SELECT id, name, description, is_active FROM asset_categories WHERE id = $1`
	ac := &categoryDomain.AssetCategory{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&ac.ID, &ac.Name, &ac.Description, &ac.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category not found: %s", id)
		}
		return nil, err
	}
	return ac, nil
}
