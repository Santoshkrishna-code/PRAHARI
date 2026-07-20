package postgres

import (
	"context"
	"database/sql"

	renewalDomain "prahari/services/training/internal/domain/renewal"
)

// RenewalStore implements certification renewals database.
type RenewalStore struct {
	db *sql.DB
}

// NewRenewalStore instantiates RenewalStore.
func NewRenewalStore(db *sql.DB) *RenewalStore {
	return &RenewalStore{db: db}
}

// Create persists renewal schedules.
func (s *RenewalStore) Create(ctx context.Context, r *renewalDomain.Renewal) error {
	query := `INSERT INTO renewals (id, certification_id, scheduled_date, is_completed)
		VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.CertificationID, r.ScheduledDate, r.IsCompleted)
	return err
}
