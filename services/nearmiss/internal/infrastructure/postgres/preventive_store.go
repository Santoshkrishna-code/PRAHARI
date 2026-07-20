package postgres

import (
	"context"
	"database/sql"

	preventiveDomain "prahari/services/nearmiss/internal/domain/preventiveaction"
)

// PreventiveStore implements preventive actions storage.
type PreventiveStore struct {
	db *sql.DB
}

// NewPreventiveStore instantiates PreventiveStore.
func NewPreventiveStore(db *sql.DB) *PreventiveStore {
	return &PreventiveStore{db: db}
}

// Create persists preventive actions.
func (s *PreventiveStore) Create(ctx context.Context, pa *preventiveDomain.PreventiveAction) error {
	query := `INSERT INTO near_miss_preventive_actions (id, near_miss_id, description, target_date, responsible_party_id, is_implemented)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, pa.ID, pa.NearMissID, pa.Description, pa.TargetDate, pa.ResponsiblePartyID, pa.IsImplemented)
	return err
}

// FindByNearMissID returns preventive actions list.
func (s *PreventiveStore) FindByNearMissID(ctx context.Context, nearmissID string) ([]*preventiveDomain.PreventiveAction, error) {
	query := `SELECT id, near_miss_id, description, target_date, responsible_party_id, is_implemented FROM near_miss_preventive_actions WHERE near_miss_id = $1`
	rows, err := s.db.QueryContext(ctx, query, nearmissID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*preventiveDomain.PreventiveAction
	for rows.Next() {
		pa := &preventiveDomain.PreventiveAction{}
		err = rows.Scan(&pa.ID, &pa.NearMissID, &pa.Description, &pa.TargetDate, &pa.ResponsiblePartyID, &pa.IsImplemented)
		if err != nil {
			return nil, err
		}
		list = append(list, pa)
	}
	return list, nil
}
