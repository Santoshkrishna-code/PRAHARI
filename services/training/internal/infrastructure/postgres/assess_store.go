package postgres

import (
	"context"
	"database/sql"

	assessDomain "prahari/services/training/internal/domain/assessment"
)

// AssessStore implements assessments evaluations database.
type AssessStore struct {
	db *sql.DB
}

// NewAssessStore instantiates AssessStore.
func NewAssessStore(db *sql.DB) *AssessStore {
	return &AssessStore{db: db}
}

// Create persists evaluations.
func (s *AssessStore) Create(ctx context.Context, a *assessDomain.Assessment) error {
	query := `INSERT INTO assessments (id, training_id, trainee_id, score, is_passed)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.TrainingID, a.TraineeID, a.Score, a.IsPassed)
	return err
}
