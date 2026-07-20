package postgres

import (
	"context"
	"database/sql"

	trainingDomain "prahari/services/contractor/internal/domain/training"
)

// TrainingStore implements training courses persistence operations.
type TrainingStore struct {
	db *sql.DB
}

// NewTrainingStore instantiates TrainingStore.
func NewTrainingStore(db *sql.DB) *TrainingStore {
	return &TrainingStore{db: db}
}

// Create persists training course record.
func (s *TrainingStore) Create(ctx context.Context, t *trainingDomain.Training) error {
	query := `INSERT INTO contractor_trainings (id, worker_id, course_name, completed_at, expiry_date)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, t.ID, t.WorkerID, t.CourseName, t.CompletedAt, t.ExpiryDate)
	return err
}

// FindByWorkerID returns worker safety courses credentials list.
func (s *TrainingStore) FindByWorkerID(ctx context.Context, workerID string) ([]*trainingDomain.Training, error) {
	query := `SELECT id, worker_id, course_name, completed_at, expiry_date FROM contractor_trainings WHERE worker_id = $1`
	rows, err := s.db.QueryContext(ctx, query, workerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*trainingDomain.Training
	for rows.Next() {
		t := &trainingDomain.Training{}
		err = rows.Scan(&t.ID, &t.WorkerID, &t.CourseName, &t.CompletedAt, &t.ExpiryDate)
		if err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	return list, nil
}
