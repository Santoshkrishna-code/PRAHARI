package postgres

import (
	"context"
	"database/sql"

	matrixDomain "prahari/services/risk/internal/domain/riskmatrix"
)

// MatrixStore implements configurable 5x5 dynamic scoring matrices database.
type MatrixStore struct {
	db *sql.DB
}

// NewMatrixStore instantiates MatrixStore.
func NewMatrixStore(db *sql.DB) *MatrixStore {
	return &MatrixStore{db: db}
}

// SaveConfig registers scoring rule.
func (s *MatrixStore) SaveConfig(ctx context.Context, rm *matrixDomain.RiskMatrix) error {
	query := `INSERT INTO risk_matrix (id, likelihood, consequence) VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE SET likelihood = $2, consequence = $3`
	_, err := s.db.ExecContext(ctx, query, rm.ID, rm.Likelihood, rm.Consequence)
	return err
}
