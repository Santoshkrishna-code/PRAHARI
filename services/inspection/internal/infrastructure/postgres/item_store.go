package postgres

import (
	"context"
	"database/sql"

	itemDomain "prahari/services/inspection/internal/domain/checklistitem"
)

// ItemStore implements individual item checks queries.
type ItemStore struct {
	db *sql.DB
}

// NewItemStore instantiates ItemStore.
func NewItemStore(db *sql.DB) *ItemStore {
	return &ItemStore{db: db}
}

// Create inserts item checks.
func (s *ItemStore) Create(ctx context.Context, ci *itemDomain.ChecklistItem) error {
	query := `INSERT INTO inspection_items (id, checklist_id, question, description, category_name, response_type, response_value, is_passed, comments)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.ExecContext(ctx, query, ci.ID, ci.ChecklistID, ci.Question, ci.Description, ci.CategoryName, ci.ResponseType, ci.ResponseValue, ci.IsPassed, ci.Comments)
	return err
}

// FindByChecklistID retrieves questions responses.
func (s *ItemStore) FindByChecklistID(ctx context.Context, checklistID string) ([]*itemDomain.ChecklistItem, error) {
	query := `SELECT id, checklist_id, question, description, category_name, response_type, response_value, is_passed, comments
		FROM inspection_items WHERE checklist_id = $1`
	rows, err := s.db.QueryContext(ctx, query, checklistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*itemDomain.ChecklistItem
	for rows.Next() {
		ci := &itemDomain.ChecklistItem{}
		err = rows.Scan(&ci.ID, &ci.ChecklistID, &ci.Question, &ci.Description, &ci.CategoryName, &ci.ResponseType, &ci.ResponseValue, &ci.IsPassed, &ci.Comments)
		if err != nil {
			return nil, err
		}
		items = append(items, ci)
	}
	return items, nil
}

// Update saves responses.
func (s *ItemStore) Update(ctx context.Context, ci *itemDomain.ChecklistItem) error {
	query := `UPDATE inspection_items SET response_value = $2, is_passed = $3, comments = $4 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, ci.ID, ci.ResponseValue, ci.IsPassed, ci.Comments)
	return err
}
