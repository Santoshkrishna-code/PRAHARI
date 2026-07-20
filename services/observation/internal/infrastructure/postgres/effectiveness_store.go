package postgres

import (
	"context"
	"database/sql"

	effectivenessDomain "prahari/services/observation/internal/domain/effectiveness"
)

// EffectivenessStore implements coaching effectiveness ratings.
type EffectivenessStore struct {
	db *sql.DB
}

// NewEffectivenessStore instantiates EffectivenessStore.
func NewEffectivenessStore(db *sql.DB) *EffectivenessStore {
	return &EffectivenessStore{db: db}
}

// Create persists evaluation.
func (s *EffectivenessStore) Create(ctx context.Context, e *effectivenessDomain.Effectiveness) error {
	query := `INSERT INTO effectiveness (id, observation_id, evaluator_id, evaluation_date, improvement_rate, notes)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, e.ID, e.ObservationID, e.EvaluatorID, e.EvaluationDate, e.ImprovementRate, e.Notes)
	return err
}

// FindByObservationID returns evaluations list.
func (s *EffectivenessStore) FindByObservationID(ctx context.Context, observationID string) ([]*effectivenessDomain.Effectiveness, error) {
	query := `SELECT id, observation_id, evaluator_id, evaluation_date, improvement_rate, notes FROM effectiveness WHERE observation_id = $1`
	rows, err := s.db.QueryContext(ctx, query, observationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*effectivenessDomain.Effectiveness
	for rows.Next() {
		e := &effectivenessDomain.Effectiveness{}
		err = rows.Scan(&e.ID, &e.ObservationID, &e.EvaluatorID, &e.EvaluationDate, &e.ImprovementRate, &e.Notes)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}
