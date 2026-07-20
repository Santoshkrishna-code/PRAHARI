package postgres

import (
	"context"
	"database/sql"
	"fmt"

	checklistDomain "prahari/services/audit/internal/domain/checklist"
)

// ChecklistStore implements checklist frameworks frameworks database.
type ChecklistStore struct {
	db *sql.DB
}

// NewChecklistStore instantiates ChecklistStore.
func NewChecklistStore(db *sql.DB) *ChecklistStore {
	return &ChecklistStore{db: db}
}

// Create persists checklist name.
func (s *ChecklistStore) Create(ctx context.Context, c *checklistDomain.Checklist) error {
	query := `INSERT INTO audit_checklists (id, name, description) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.Name, c.Description)
	return err
}

// FindByID returns checklist details.
func (s *ChecklistStore) FindByID(ctx context.Context, id string) (*checklistDomain.Checklist, error) {
	query := `SELECT id, name, description FROM audit_checklists WHERE id = $1`
	c := &checklistDomain.Checklist{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("checklist not found: %s", id)
		}
		return nil, err
	}
	return c, nil
}
