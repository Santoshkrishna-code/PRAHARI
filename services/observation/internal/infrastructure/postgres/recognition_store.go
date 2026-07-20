package postgres

import (
	"context"
	"database/sql"

	recognitionDomain "prahari/services/observation/internal/domain/recognition"
)

// RecognitionStore implements rewards storage.
type RecognitionStore struct {
	db *sql.DB
}

// NewRecognitionStore instantiates RecognitionStore.
func NewRecognitionStore(db *sql.DB) *RecognitionStore {
	return &RecognitionStore{db: db}
}

// Create persists positive recognition.
func (s *RecognitionStore) Create(ctx context.Context, r *recognitionDomain.Recognition) error {
	query := `INSERT INTO recognitions (id, observation_id, recognized_person_id, granted_by_id, reason, granted_at)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.ObservationID, r.RecognizedPersonID, r.GrantedByID, r.Reason, r.GrantedAt)
	return err
}

// FindByObservationID returns recognitions list.
func (s *RecognitionStore) FindByObservationID(ctx context.Context, observationID string) ([]*recognitionDomain.Recognition, error) {
	query := `SELECT id, observation_id, recognized_person_id, granted_by_id, reason, granted_at FROM recognitions WHERE observation_id = $1`
	rows, err := s.db.QueryContext(ctx, query, observationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*recognitionDomain.Recognition
	for rows.Next() {
		r := &recognitionDomain.Recognition{}
		err = rows.Scan(&r.ID, &r.ObservationID, &r.RecognizedPersonID, &r.GrantedByID, &r.Reason, &r.GrantedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}
	return list, nil
}
