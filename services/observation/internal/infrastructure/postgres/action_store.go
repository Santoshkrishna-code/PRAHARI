package postgres

import (
	"context"
	"database/sql"

	actionDomain "prahari/services/observation/internal/domain/action"
)

// ActionStore implements corrective actions storage.
type ActionStore struct {
	db *sql.DB
}

// NewActionStore instantiates ActionStore.
func NewActionStore(db *sql.DB) *ActionStore {
	return &ActionStore{db: db}
}

// Create persists actions.
func (s *ActionStore) Create(ctx context.Context, a *actionDomain.Action) error {
	query := `INSERT INTO observation_actions (id, observation_id, description, target_date, responsible_party_id, is_implemented)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.ObservationID, a.Description, a.TargetDate, a.ResponsiblePartyID, a.IsImplemented)
	return err
}

// FindByObservationID returns action items.
func (s *ActionStore) FindByObservationID(ctx context.Context, observationID string) ([]*actionDomain.Action, error) {
	query := `SELECT id, observation_id, description, target_date, responsible_party_id, is_implemented FROM observation_actions WHERE observation_id = $1`
	rows, err := s.db.QueryContext(ctx, query, observationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*actionDomain.Action
	for rows.Next() {
		a := &actionDomain.Action{}
		err = rows.Scan(&a.ID, &a.ObservationID, &a.Description, &a.TargetDate, &a.ResponsiblePartyID, &a.IsImplemented)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}
