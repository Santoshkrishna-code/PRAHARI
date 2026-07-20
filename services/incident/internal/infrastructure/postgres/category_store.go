package postgres

import (
	"context"
	"database/sql"
	"fmt"

	categoryDomain "prahari/services/incident/internal/domain/category"
)

// CategoryStore implements the category persistence adapter against PostgreSQL.
type CategoryStore struct {
	db *sql.DB
}

// NewCategoryStore constructs a CategoryStore.
func NewCategoryStore(db *sql.DB) *CategoryStore {
	return &CategoryStore{db: db}
}

// Create persists a new category.
func (s *CategoryStore) Create(ctx context.Context, cat *categoryDomain.Category) error {
	query := `INSERT INTO incident_categories (id, name, description, parent_id, sort_order, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, cat.ID, cat.Name, cat.Description, cat.ParentID, cat.SortOrder, cat.IsActive)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert category: %w", err)
	}
	return nil
}

// FindByID retrieves a category by its unique identifier.
func (s *CategoryStore) FindByID(ctx context.Context, id string) (*categoryDomain.Category, error) {
	query := `SELECT id, name, description, parent_id, sort_order, is_active FROM incident_categories WHERE id = $1`
	cat := &categoryDomain.Category{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&cat.ID, &cat.Name, &cat.Description, &cat.ParentID, &cat.SortOrder, &cat.IsActive)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to query category: %w", err)
	}
	return cat, nil
}

// ListActive retrieves all active categories ordered by sort order.
func (s *CategoryStore) ListActive(ctx context.Context) ([]*categoryDomain.Category, error) {
	query := `SELECT id, name, description, parent_id, sort_order, is_active FROM incident_categories WHERE is_active = true ORDER BY sort_order`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to list categories: %w", err)
	}
	defer rows.Close()

	var categories []*categoryDomain.Category
	for rows.Next() {
		cat := &categoryDomain.Category{}
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.Description, &cat.ParentID, &cat.SortOrder, &cat.IsActive); err != nil {
			return nil, fmt.Errorf("postgres: failed to scan category: %w", err)
		}
		categories = append(categories, cat)
	}
	return categories, nil
}
