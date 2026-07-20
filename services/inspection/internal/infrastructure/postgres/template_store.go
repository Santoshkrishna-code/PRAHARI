package postgres

import (
	"context"
	"database/sql"
	"fmt"

	templateDomain "prahari/services/inspection/internal/domain/checklisttemplate"
)

// TemplateStore implements reusable checklist templates queries.
type TemplateStore struct {
	db *sql.DB
}

// NewTemplateStore instantiates a TemplateStore.
func NewTemplateStore(db *sql.DB) *TemplateStore {
	return &TemplateStore{db: db}
}

// Create inserts template.
func (s *TemplateStore) Create(ctx context.Context, ct *templateDomain.ChecklistTemplate) error {
	query := `INSERT INTO inspection_templates (id, name, description, categories, items, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, ct.ID, ct.Name, ct.Description, ct.Categories, ct.Items, ct.IsActive)
	return err
}

// FindByID returns a checklist template.
func (s *TemplateStore) FindByID(ctx context.Context, id string) (*templateDomain.ChecklistTemplate, error) {
	query := `SELECT id, name, description, categories, items, is_active FROM inspection_templates WHERE id = $1`
	ct := &templateDomain.ChecklistTemplate{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&ct.ID, &ct.Name, &ct.Description, &ct.Categories, &ct.Items, &ct.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("checklist template not found: %s", id)
		}
		return nil, err
	}
	return ct, nil
}

// ListActive retrieves templates marked active.
func (s *TemplateStore) ListActive(ctx context.Context) ([]*templateDomain.ChecklistTemplate, error) {
	query := `SELECT id, name, description, categories, items, is_active FROM inspection_templates WHERE is_active = true`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []*templateDomain.ChecklistTemplate
	for rows.Next() {
		ct := &templateDomain.ChecklistTemplate{}
		if err := rows.Scan(&ct.ID, &ct.Name, &ct.Description, &ct.Categories, &ct.Items, &ct.IsActive); err != nil {
			return nil, err
		}
		templates = append(templates, ct)
	}
	return templates, nil
}
