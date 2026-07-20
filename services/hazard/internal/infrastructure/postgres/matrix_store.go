package postgres

import (
	"context"
	"database/sql"
	"fmt"

	matrixDomain "prahari/services/hazard/internal/domain/riskmatrix"
)

// MatrixStore implements 5x5 matrix settings.
type MatrixStore struct {
	db *sql.DB
}

// NewMatrixStore instantiates MatrixStore.
func NewMatrixStore(db *sql.DB) *MatrixStore {
	return &MatrixStore{db: db}
}

// Create persists matrix logic values.
func (s *MatrixStore) Create(ctx context.Context, rm *matrixDomain.RiskMatrix) error {
	query := `INSERT INTO hazard_risk_matrix (id, likelihood, consequence) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, rm.ID, rm.Likelihood, rm.Consequence)
	return err
}

// FindByID returns matrix.
func (s *MatrixStore) FindByID(ctx context.Context, id string) (*matrixDomain.RiskMatrix, error) {
	query := `SELECT id, likelihood, consequence FROM hazard_risk_matrix WHERE id = $1`
	rm := &matrixDomain.RiskMatrix{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&rm.ID, &rm.Likelihood, &rm.Consequence)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("matrix not found: %s", id)
		}
		return nil, err
	}
	return rm, nil
}
