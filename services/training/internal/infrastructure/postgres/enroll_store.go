package postgres

import (
	"context"
	"database/sql"

	enrollDomain "prahari/services/training/internal/domain/enrollment"
)

// EnrollStore implements trainee course enrollments database.
type EnrollStore struct {
	db *sql.DB
}

// NewEnrollStore instantiates EnrollStore.
func NewEnrollStore(db *sql.DB) *EnrollStore {
	return &EnrollStore{db: db}
}

// Create persists enrollment metrics.
func (s *EnrollStore) Create(ctx context.Context, e *enrollDomain.Enrollment) error {
	query := `INSERT INTO enrollments (id, training_id, trainee_id, status) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, e.ID, e.TrainingID, e.TraineeID, e.Status)
	return err
}
