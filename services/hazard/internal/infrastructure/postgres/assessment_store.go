package postgres

import (
	"context"
	"database/sql"
	"fmt"

	matrixDomain "prahari/services/hazard/internal/domain/riskmatrix"
)

// AssessmentStore implements assessments parameters storage.
type AssessmentStore struct {
	db *sql.DB
}

// NewAssessmentStore instantiates AssessmentStore.
func NewAssessmentStore(db *sql.DB) *AssessmentStore {
	return &AssessmentStore{db: db}
}

// Create persists assessment metrics.
func (s *AssessmentStore) Create(ctx context.Context, rm *matrixDomain.RiskMatrix) error {
	query := `INSERT INTO hazard_assessments (id, likelihood, consequence) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, rm.ID, rm.Likelihood, rm.Consequence)
	return err
}

// FindByID returns assessment details.
func (s *AssessmentStore) FindByID(ctx context.Context, id string) (*matrixDomain.RiskMatrix, error) {
	query := `SELECT id, likelihood, consequence FROM hazard_assessments WHERE id = $1`
	rm := &matrixDomain.RiskMatrix{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&rm.ID, &rm.Likelihood, &rm.Consequence)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("assessment not found: %s", id)
		}
		return nil, err
	}
	return rm, nil
}
