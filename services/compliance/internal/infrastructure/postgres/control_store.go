package postgres

import (
	"context"
	"database/sql"

	controlDomain "prahari/services/compliance/internal/domain/control"
)

// ControlStore implements compliance check procedures database.
type ControlStore struct {
	db *sql.DB
}

// NewControlStore instantiates ControlStore.
func NewControlStore(db *sql.DB) *ControlStore {
	return &ControlStore{db: db}
}

// Create persists control procedure details.
func (s *ControlStore) Create(ctx context.Context, c *controlDomain.Control) error {
	query := `INSERT INTO controls (id, compliance_id, description) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.ComplianceID, c.Description)
	return err
}

// FindByComplianceID returns check procedures list.
func (s *ControlStore) FindByComplianceID(ctx context.Context, complianceID string) ([]*controlDomain.Control, error) {
	query := `SELECT id, compliance_id, description FROM controls WHERE compliance_id = $1`
	rows, err := s.db.QueryContext(ctx, query, complianceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*controlDomain.Control
	for rows.Next() {
		c := &controlDomain.Control{}
		err = rows.Scan(&c.ID, &c.ComplianceID, &c.Description)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
