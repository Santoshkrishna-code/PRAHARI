package postgres

import (
	"context"
	"database/sql"

	followupDomain "prahari/services/observation/internal/domain/followup"
)

// FollowUpStore implements follow-up verifications logs storage.
type FollowUpStore struct {
	db *sql.DB
}

// NewFollowUpStore instantiates FollowUpStore.
func NewFollowUpStore(db *sql.DB) *FollowUpStore {
	return &FollowUpStore{db: db}
}

// Create persists validation checks.
func (s *FollowUpStore) Create(ctx context.Context, f *followupDomain.FollowUp) error {
	query := `INSERT INTO followups (id, observation_id, follower_id, follow_up_date, notes, is_passed)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, f.ID, f.ObservationID, f.FollowerID, f.FollowUpDate, f.Notes, f.IsPassed)
	return err
}

// FindByObservationID returns followups.
func (s *FollowUpStore) FindByObservationID(ctx context.Context, observationID string) ([]*followupDomain.FollowUp, error) {
	query := `SELECT id, observation_id, follower_id, follow_up_date, notes, is_passed FROM followups WHERE observation_id = $1`
	rows, err := s.db.QueryContext(ctx, query, observationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*followupDomain.FollowUp
	for rows.Next() {
		f := &followupDomain.FollowUp{}
		err = rows.Scan(&f.ID, &f.ObservationID, &f.FollowerID, &f.FollowUpDate, &f.Notes, &f.IsPassed)
		if err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, nil
}
