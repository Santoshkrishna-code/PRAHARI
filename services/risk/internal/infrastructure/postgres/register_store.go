package postgres

import (
	"context"
	"database/sql"
	"fmt"

	riskDomain "prahari/services/risk/internal/domain/risk"
)

// RegisterStore implements process safety register catalog.
type RegisterStore struct {
	db *sql.DB
}

// NewRegisterStore instantiates RegisterStore.
func NewRegisterStore(db *sql.DB) *RegisterStore {
	return &RegisterStore{db: db}
}

// Create persists catalog register code.
func (s *RegisterStore) Create(ctx context.Context, r *riskDomain.Risk) error {
	query := `INSERT INTO risk_register (id, risk_number, title, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.RiskNumber, r.Title, r.Description)
	return err
}

// FindByID returns register details.
func (s *RegisterStore) FindByID(ctx context.Context, id string) (*riskDomain.Risk, error) {
	query := `SELECT id, risk_number, title, description FROM risk_register WHERE id = $1`
	r := &riskDomain.Risk{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&r.ID, &r.RiskNumber, &r.Title, &r.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("register details not found: %s", id)
		}
		return nil, err
	}
	return r, nil
}
