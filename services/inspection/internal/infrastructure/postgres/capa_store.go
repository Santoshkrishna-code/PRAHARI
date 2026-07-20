package postgres

import (
	"context"
	"database/sql"
	"time"

	capaDomain "prahari/services/inspection/internal/domain/capa"
)

// CAPAStore implements action tracking tables operations.
type CAPAStore struct {
	db *sql.DB
}

// NewCAPAStore instantiates CAPAStore.
func NewCAPAStore(db *sql.DB) *CAPAStore {
	return &CAPAStore{db: db}
}

// Create persists action schedules.
func (s *CAPAStore) Create(ctx context.Context, c *capaDomain.CAPA) error {
	query := `INSERT INTO inspection_actions (id, inspection_id, finding_id, action_type, description, assignee_id, due_date, completed_at, status, verified_by, verified_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := s.db.ExecContext(ctx, query,
		c.ID, c.InspectionID, c.FindingID, c.ActionType, c.Description, c.AssigneeID, c.DueDate, c.CompletedAt, c.Status, c.VerifiedBy, c.VerifiedAt, c.CreatedAt,
	)
	return err
}

// FindByID returns an action details.
func (s *CAPAStore) FindByID(ctx context.Context, id string) (*capaDomain.CAPA, error) {
	query := `SELECT id, inspection_id, finding_id, action_type, description, assignee_id, due_date, completed_at, status, verified_by, verified_at, created_at
		FROM inspection_actions WHERE id = $1`
	c := &capaDomain.CAPA{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.InspectionID, &c.FindingID, &c.ActionType, &c.Description, &c.AssigneeID, &c.DueDate, &c.CompletedAt, &c.Status, &c.VerifiedBy, &c.VerifiedAt, &c.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// FindByInspectionID returns actions on an inspection.
func (s *CAPAStore) FindByInspectionID(ctx context.Context, inspectionID string) ([]*capaDomain.CAPA, error) {
	query := `SELECT id, inspection_id, finding_id, action_type, description, assignee_id, due_date, completed_at, status, verified_by, verified_at, created_at
		FROM inspection_actions WHERE inspection_id = $1`
	rows, err := s.db.QueryContext(ctx, query, inspectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actions []*capaDomain.CAPA
	for rows.Next() {
		c := &capaDomain.CAPA{}
		err = rows.Scan(&c.ID, &c.InspectionID, &c.FindingID, &c.ActionType, &c.Description, &c.AssigneeID, &c.DueDate, &c.CompletedAt, &c.Status, &c.VerifiedBy, &c.VerifiedAt, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		actions = append(actions, c)
	}
	return actions, nil
}

// Update registers edits.
func (s *CAPAStore) Update(ctx context.Context, c *capaDomain.CAPA) error {
	query := `UPDATE inspection_actions SET completed_at = $2, status = $3, verified_by = $4, verified_at = $5 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.CompletedAt, c.Status, c.VerifiedBy, c.VerifiedAt)
	return err
}
