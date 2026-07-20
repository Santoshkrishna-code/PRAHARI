package postgres

import (
	"context"
	"database/sql"

	correctiveDomain "prahari/services/nearmiss/internal/domain/correctiveaction"
)

// CorrectiveStore implements CAPA items storage.
type CorrectiveStore struct {
	db *sql.DB
}

// NewCorrectiveStore instantiates CorrectiveStore.
func NewCorrectiveStore(db *sql.DB) *CorrectiveStore {
	return &CorrectiveStore{db: db}
}

// Create persists corrective actions.
func (s *CorrectiveStore) Create(ctx context.Context, ca *correctiveDomain.CorrectiveAction) error {
	query := `INSERT INTO near_miss_corrective_actions (id, near_miss_id, description, target_date, responsible_party_id, is_implemented)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, ca.ID, ca.NearMissID, ca.Description, ca.TargetDate, ca.ResponsiblePartyID, ca.IsImplemented)
	return err
}

// FindByNearMissID returns corrective actions.
func (s *CorrectiveStore) FindByNearMissID(ctx context.Context, nearmissID string) ([]*correctiveDomain.CorrectiveAction, error) {
	query := `SELECT id, near_miss_id, description, target_date, responsible_party_id, is_implemented FROM near_miss_corrective_actions WHERE near_miss_id = $1`
	rows, err := s.db.QueryContext(ctx, query, nearmissID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*correctiveDomain.CorrectiveAction
	for rows.Next() {
		ca := &correctiveDomain.CorrectiveAction{}
		err = rows.Scan(&ca.ID, &ca.NearMissID, &ca.Description, &ca.TargetDate, &ca.ResponsiblePartyID, &ca.IsImplemented)
		if err != nil {
			return nil, err
		}
		list = append(list, ca)
	}
	return list, nil
}
