package postgres

import (
	"context"
	"database/sql"

	checklistDomain "prahari/services/maintenance/internal/domain/checklist"
)

// ChecklistStore implements checklist status tracking operations.
type ChecklistStore struct {
	db *sql.DB
}

// NewChecklistStore instantiates ChecklistStore.
func NewChecklistStore(db *sql.DB) *ChecklistStore {
	return &ChecklistStore{db: db}
}

// Create persists checklist check status.
func (s *ChecklistStore) Create(ctx context.Context, c *checklistDomain.Checklist) error {
	query := `INSERT INTO maintenance_checklists (id, maintenance_id, name, is_passed) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.MaintenanceID, c.Name, c.IsPassed)
	return err
}

// FindByMaintenanceID returns checklist check items status list.
func (s *ChecklistStore) FindByMaintenanceID(ctx context.Context, maintenanceID string) ([]*checklistDomain.Checklist, error) {
	query := `SELECT id, maintenance_id, name, is_passed FROM maintenance_checklists WHERE maintenance_id = $1`
	rows, err := s.db.QueryContext(ctx, query, maintenanceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*checklistDomain.Checklist
	for rows.Next() {
		c := &checklistDomain.Checklist{}
		err = rows.Scan(&c.ID, &c.MaintenanceID, &c.Name, &c.IsPassed)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
