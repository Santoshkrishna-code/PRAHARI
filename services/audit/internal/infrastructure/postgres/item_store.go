package postgres

import (
	"context"
	"database/sql"

	itemDomain "prahari/services/audit/internal/domain/checklistitem"
)

// ItemStore implements specific compliance checklist questions.
type ItemStore struct {
	db *sql.DB
}

// NewItemStore instantiates ItemStore.
func NewItemStore(db *sql.DB) *ItemStore {
	return &ItemStore{db: db}
}

// Create persists question item.
func (s *ItemStore) Create(ctx context.Context, ci *itemDomain.ChecklistItem) error {
	query := `INSERT INTO audit_checklist_items (id, checklist_id, question) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, ci.ID, ci.ChecklistID, ci.Question)
	return err
}

// FindByChecklistID returns questions lists.
func (s *ItemStore) FindByChecklistID(ctx context.Context, checklistID string) ([]*itemDomain.ChecklistItem, error) {
	query := `SELECT id, checklist_id, question FROM audit_checklist_items WHERE checklist_id = $1`
	rows, err := s.db.QueryContext(ctx, query, checklistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*itemDomain.ChecklistItem
	for rows.Next() {
		ci := &itemDomain.ChecklistItem{}
		err = rows.Scan(&ci.ID, &ci.ChecklistID, &ci.Question)
		if err != nil {
			return nil, err
		}
		list = append(list, ci)
	}
	return list, nil
}
