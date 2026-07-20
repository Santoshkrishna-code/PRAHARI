package postgres

import (
	"context"
	"database/sql"
	"fmt"

	checklistDomain "prahari/services/inspection/internal/domain/checklist"
)

// ChecklistStore implements checklist instances storage operations.
type ChecklistStore struct {
	db *sql.DB
}

// NewChecklistStore instantiates ChecklistStore.
func NewChecklistStore(db *sql.DB) *ChecklistStore {
	return &ChecklistStore{db: db}
}

// Create persists a checklist.
func (s *ChecklistStore) Create(ctx context.Context, c *checklistDomain.Checklist) error {
	query := `INSERT INTO inspection_checklists (id, inspection_id, checklist_template_id, name)
		VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.InspectionID, c.ChecklistTemplateID, c.Name)
	return err
}

// FindByID returns a checklist.
func (s *ChecklistStore) FindByID(ctx context.Context, id string) (*checklistDomain.Checklist, error) {
	query := `SELECT id, inspection_id, checklist_template_id, name FROM inspection_checklists WHERE id = $1`
	c := &checklistDomain.Checklist{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.InspectionID, &c.ChecklistTemplateID, &c.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("checklist not found: %s", id)
		}
		return nil, err
	}
	return c, nil
}

// FindByInspectionID returns generated checklists for an inspection.
func (s *ChecklistStore) FindByInspectionID(ctx context.Context, inspectionID string) ([]*checklistDomain.Checklist, error) {
	query := `SELECT id, inspection_id, checklist_template_id, name FROM inspection_checklists WHERE inspection_id = $1`
	rows, err := s.db.QueryContext(ctx, query, inspectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var checklists []*checklistDomain.Checklist
	for rows.Next() {
		c := &checklistDomain.Checklist{}
		err = rows.Scan(&c.ID, &c.InspectionID, &c.ChecklistTemplateID, &c.Name)
		if err != nil {
			return nil, err
		}
		checklists = append(checklists, c)
	}
	return checklists, nil
}
