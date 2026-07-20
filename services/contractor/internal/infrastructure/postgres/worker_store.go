package postgres

import (
	"context"
	"database/sql"
	"fmt"

	workerDomain "prahari/services/contractor/internal/domain/worker"
)

// WorkerStore implements contractor employees persistence operations.
type WorkerStore struct {
	db *sql.DB
}

// NewWorkerStore instantiates WorkerStore.
func NewWorkerStore(db *sql.DB) *WorkerStore {
	return &WorkerStore{db: db}
}

// Create persists a worker.
func (s *WorkerStore) Create(ctx context.Context, w *workerDomain.Worker) error {
	query := `INSERT INTO contractor_workers (id, contractor_id, first_name, last_name, passport_id, onboarding_status)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, w.ID, w.ContractorID, w.FirstName, w.LastName, w.PassportID, w.OnboardingStatus)
	return err
}

// FindByID returns a worker.
func (s *WorkerStore) FindByID(ctx context.Context, id string) (*workerDomain.Worker, error) {
	query := `SELECT id, contractor_id, first_name, last_name, passport_id, onboarding_status FROM contractor_workers WHERE id = $1`
	w := &workerDomain.Worker{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&w.ID, &w.ContractorID, &w.FirstName, &w.LastName, &w.PassportID, &w.OnboardingStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("worker not found: %s", id)
		}
		return nil, err
	}
	return w, nil
}

// FindByContractorID returns worker crew members list.
func (s *WorkerStore) FindByContractorID(ctx context.Context, contractorID string) ([]*workerDomain.Worker, error) {
	query := `SELECT id, contractor_id, first_name, last_name, passport_id, onboarding_status FROM contractor_workers WHERE contractor_id = $1`
	rows, err := s.db.QueryContext(ctx, query, contractorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*workerDomain.Worker
	for rows.Next() {
		w := &workerDomain.Worker{}
		err = rows.Scan(&w.ID, &w.ContractorID, &w.FirstName, &w.LastName, &w.PassportID, &w.OnboardingStatus)
		if err != nil {
			return nil, err
		}
		list = append(list, w)
	}
	return list, nil
}
