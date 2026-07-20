package postgres

import (
	"context"
	"database/sql"

	coachingDomain "prahari/services/observation/internal/domain/coaching"
)

// CoachingStore implements coaching sessions storage.
type CoachingStore struct {
	db *sql.DB
}

// NewCoachingStore instantiates CoachingStore.
func NewCoachingStore(db *sql.DB) *CoachingStore {
	return &CoachingStore{db: db}
}

// Create persists session details.
func (s *CoachingStore) Create(ctx context.Context, cs *coachingDomain.CoachingSession) error {
	query := `INSERT INTO coaching_sessions (id, observation_id, coach_id, session_date, topics, feedback)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, cs.ID, cs.ObservationID, cs.CoachID, cs.SessionDate, cs.Topics, cs.Feedback)
	return err
}

// FindByObservationID returns coaching sessions list.
func (s *CoachingStore) FindByObservationID(ctx context.Context, observationID string) ([]*coachingDomain.CoachingSession, error) {
	query := `SELECT id, observation_id, coach_id, session_date, topics, feedback FROM coaching_sessions WHERE observation_id = $1`
	rows, err := s.db.QueryContext(ctx, query, observationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*coachingDomain.CoachingSession
	for rows.Next() {
		cs := &coachingDomain.CoachingSession{}
		err = rows.Scan(&cs.ID, &cs.ObservationID, &cs.CoachID, &cs.SessionDate, &cs.Topics, &cs.Feedback)
		if err != nil {
			return nil, err
		}
		list = append(list, cs)
	}
	return list, nil
}
