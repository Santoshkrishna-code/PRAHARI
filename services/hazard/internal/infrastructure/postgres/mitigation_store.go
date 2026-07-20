package postgres

import (
	"context"
	"database/sql"

	mitigationDomain "prahari/services/hazard/internal/domain/mitigation"
)

// MitigationStore implements mitigation plan details storage.
type MitigationStore struct {
	db *sql.DB
}

// NewMitigationStore instantiates MitigationStore.
func NewMitigationStore(db *sql.DB) *MitigationStore {
	return &MitigationStore{db: db}
}

// Create persists plan parameters.
func (s *MitigationStore) Create(ctx context.Context, m *mitigationDomain.Mitigation) error {
	query := `INSERT INTO hazard_mitigations (id, hazard_id, description, target_completion_date, responsible_party_id, is_implemented)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, m.ID, m.HazardID, m.Description, m.TargetCompletionDate, m.ResponsiblePartyID, m.IsImplemented)
	return err
}

// FindByHazardID returns mitigations list.
func (s *MitigationStore) FindByHazardID(ctx context.Context, hazardID string) ([]*mitigationDomain.Mitigation, error) {
	query := `SELECT id, hazard_id, description, target_completion_date, responsible_party_id, is_implemented FROM hazard_mitigations WHERE hazard_id = $1`
	rows, err := s.db.QueryContext(ctx, query, hazardID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*mitigationDomain.Mitigation
	for rows.Next() {
		m := &mitigationDomain.Mitigation{}
		err = rows.Scan(&m.ID, &m.HazardID, &m.Description, &m.TargetCompletionDate, &m.ResponsiblePartyID, &m.IsImplemented)
		if err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, nil
}
