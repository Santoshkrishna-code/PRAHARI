package postgres

import (
	"context"
	"database/sql"

	barrierDomain "prahari/services/risk/internal/domain/barrier"
)

// BarrierStore implements bow-tie analysis check barrier blocks.
type BarrierStore struct {
	db *sql.DB
}

// NewBarrierStore instantiates BarrierStore.
func NewBarrierStore(db *sql.DB) *BarrierStore {
	return &BarrierStore{db: db}
}

// Create persists barrier blocks.
func (s *BarrierStore) Create(ctx context.Context, b *barrierDomain.Barrier) error {
	query := `INSERT INTO risk_barriers (id, risk_id, barrier_type, description, is_assured)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, b.ID, b.RiskID, string(b.BarrierType), b.Description, b.IsAssured)
	return err
}

// FindByRiskID returns barriers checklist.
func (s *BarrierStore) FindByRiskID(ctx context.Context, riskID string) ([]*barrierDomain.Barrier, error) {
	query := `SELECT id, risk_id, barrier_type, description, is_assured FROM risk_barriers WHERE risk_id = $1`
	rows, err := s.db.QueryContext(ctx, query, riskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*barrierDomain.Barrier
	for rows.Next() {
		b := &barrierDomain.Barrier{}
		var bType string
		err = rows.Scan(&b.ID, &b.RiskID, &bType, &b.Description, &b.IsAssured)
		if err != nil {
			return nil, err
		}
		b.BarrierType = barrierDomain.BarrierType(bType)
		list = append(list, b)
	}
	return list, nil
}
