package postgres

import (
	"context"
	"database/sql"

	controlDomain "prahari/services/risk/internal/domain/control"
)

// ControlStore implements mitigation controls database.
type ControlStore struct {
	db *sql.DB
}

// NewControlStore instantiates ControlStore.
func NewControlStore(db *sql.DB) *ControlStore {
	return &ControlStore{db: db}
}

// Create persists control mitigation detail.
func (s *ControlStore) Create(ctx context.Context, c *controlDomain.Control) error {
	query := `INSERT INTO risk_controls (id, risk_id, control_type, description, effect_value)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.RiskID, string(c.ControlType), c.Description, c.EffectValue)
	return err
}

// FindByRiskID returns controls checklist.
func (s *ControlStore) FindByRiskID(ctx context.Context, riskID string) ([]*controlDomain.Control, error) {
	query := `SELECT id, risk_id, control_type, description, effect_value FROM risk_controls WHERE risk_id = $1`
	rows, err := s.db.QueryContext(ctx, query, riskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*controlDomain.Control
	for rows.Next() {
		c := &controlDomain.Control{}
		var ctrlType string
		err = rows.Scan(&c.ID, &c.RiskID, &ctrlType, &c.Description, &c.EffectValue)
		if err != nil {
			return nil, err
		}
		c.ControlType = controlDomain.ControlType(ctrlType)
		list = append(list, c)
	}
	return list, nil
}
