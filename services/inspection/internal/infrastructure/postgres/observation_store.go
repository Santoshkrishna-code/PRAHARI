package postgres

import (
	"context"
	"database/sql"

	observationDomain "prahari/services/inspection/internal/domain/observation"
)

// ObservationStore implements walkthrough walk notes.
type ObservationStore struct {
	db *sql.DB
}

// NewObservationStore instantiates ObservationStore.
func NewObservationStore(db *sql.DB) *ObservationStore {
	return &ObservationStore{db: db}
}

// Create inserts notes.
func (s *ObservationStore) Create(ctx context.Context, o *observationDomain.Observation) error {
	query := `INSERT INTO inspection_observations (id, inspection_id, description, observer_id, observed_at)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, o.ID, o.InspectionID, o.Description, o.ObserverID, o.ObservedAt)
	return err
}

// FindByInspectionID returns observations on an inspection.
func (s *ObservationStore) FindByInspectionID(ctx context.Context, inspectionID string) ([]*observationDomain.Observation, error) {
	query := `SELECT id, inspection_id, description, observer_id, observed_at FROM inspection_observations WHERE inspection_id = $1`
	rows, err := s.db.QueryContext(ctx, query, inspectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var observations []*observationDomain.Observation
	for rows.Next() {
		o := &observationDomain.Observation{}
		err = rows.Scan(&o.ID, &o.InspectionID, &o.Description, &o.ObserverID, &o.ObservedAt)
		if err != nil {
			return nil, err
		}
		observations = append(observations, o)
	}
	return observations, nil
}
