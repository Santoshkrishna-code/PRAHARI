package postgres

import (
	"context"
	"database/sql"

	investigationDomain "prahari/services/nearmiss/internal/domain/investigation"
)

// InvestigationStore implements investigations metrics storage.
type InvestigationStore struct {
	db *sql.DB
}

// NewInvestigationStore instantiates InvestigationStore.
func NewInvestigationStore(db *sql.DB) *InvestigationStore {
	return &InvestigationStore{db: db}
}

// Create persists investigation details.
func (s *InvestigationStore) Create(ctx context.Context, i *investigationDomain.Investigation) error {
	query := `INSERT INTO near_miss_investigations (id, near_miss_id, lead_investigator_id, investigation_date, findings, methodology)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, i.ID, i.NearMissID, i.LeadInvestigatorID, i.InvestigationDate, i.Findings, i.Methodology)
	return err
}

// FindByNearMissID returns investigations list.
func (s *InvestigationStore) FindByNearMissID(ctx context.Context, nearmissID string) ([]*investigationDomain.Investigation, error) {
	query := `SELECT id, near_miss_id, lead_investigator_id, investigation_date, findings, methodology FROM near_miss_investigations WHERE near_miss_id = $1`
	rows, err := s.db.QueryContext(ctx, query, nearmissID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*investigationDomain.Investigation
	for rows.Next() {
		i := &investigationDomain.Investigation{}
		err = rows.Scan(&i.ID, &i.NearMissID, &i.LeadInvestigatorID, &i.InvestigationDate, &i.Findings, &i.Methodology)
		if err != nil {
			return nil, err
		}
		list = append(list, i)
	}
	return list, nil
}
