package postgres

import (
	"context"
	"database/sql"

	medicalDomain "prahari/services/contractor/internal/domain/medical"
)

// MedicalStore implements medical clearances logging.
type MedicalStore struct {
	db *sql.DB
}

// NewMedicalStore instantiates MedicalStore.
func NewMedicalStore(db *sql.DB) *MedicalStore {
	return &MedicalStore{db: db}
}

// Create persists physical clearance checks.
func (s *MedicalStore) Create(ctx context.Context, m *medicalDomain.Medical) error {
	query := `INSERT INTO contractor_medicals (id, worker_id, evaluated_at, expiry_date, is_fit)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, m.ID, m.WorkerID, m.EvaluatedAt, m.ExpiryDate, m.IsFit)
	return err
}

// FindByWorkerID returns physical fitness logs.
func (s *MedicalStore) FindByWorkerID(ctx context.Context, workerID string) ([]*medicalDomain.Medical, error) {
	query := `SELECT id, worker_id, evaluated_at, expiry_date, is_fit FROM contractor_medicals WHERE worker_id = $1`
	rows, err := s.db.QueryContext(ctx, query, workerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*medicalDomain.Medical
	for rows.Next() {
		m := &medicalDomain.Medical{}
		err = rows.Scan(&m.ID, &m.WorkerID, &m.EvaluatedAt, &m.ExpiryDate, &m.IsFit)
		if err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, nil
}
