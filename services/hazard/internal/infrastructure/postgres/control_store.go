package postgres

import (
	"context"
	"database/sql"

	controlDomain "prahari/services/hazard/internal/domain/control"
)

// ControlStore implements control measures hierarchy checks.
type ControlStore struct {
	db *sql.DB
}

// NewControlStore instantiates ControlStore.
func NewControlStore(db *sql.DB) *ControlStore {
	return &ControlStore{db: db}
}

// Create persists control details.
func (s *ControlStore) Create(ctx context.Context, c *controlDomain.Control) error {
	query := `INSERT INTO hazard_controls (id, hazard_id, control_type, description, is_active) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.HazardID, string(c.ControlType), c.Description, c.IsActive)
	return err
}

// FindByHazardID returns controls list.
func (s *ControlStore) FindByHazardID(ctx context.Context, hazardID string) ([]*controlDomain.Control, error) {
	query := `SELECT id, hazard_id, control_type, description, is_active FROM hazard_controls WHERE hazard_id = $1`
	rows, err := s.db.QueryContext(ctx, query, hazardID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*controlDomain.Control
	for rows.Next() {
		c := &controlDomain.Control{}
		err = rows.Scan(&c.ID, &c.HazardID, &c.ControlType, &c.Description, &c.IsActive)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
