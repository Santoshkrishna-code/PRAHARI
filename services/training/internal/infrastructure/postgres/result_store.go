package postgres

import (
	"context"
	"database/sql"

	resultDomain "prahari/services/training/internal/domain/result"
)

// ResultStore implements written test results.
type ResultStore struct {
	db *sql.DB
}

// NewResultStore instantiates ResultStore.
func NewResultStore(db *sql.DB) *ResultStore {
	return &ResultStore{db: db}
}

// Create persists pass/fail results.
func (s *ResultStore) Create(ctx context.Context, r *resultDomain.Result) error {
	query := `INSERT INTO results (id, examination_id, trainee_id, score, is_passed)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.ExaminationID, r.TraineeID, r.Score, r.IsPassed)
	return err
}
