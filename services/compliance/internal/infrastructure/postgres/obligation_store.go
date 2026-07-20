package postgres

import (
	"context"
	"database/sql"

	obDomain "prahari/services/compliance/internal/domain/obligation"
)

// ObligationStore implements statutory checklist obligations database.
type ObligationStore struct {
	db *sql.DB
}

// NewObligationStore instantiates ObligationStore.
func NewObligationStore(db *sql.DB) *ObligationStore {
	return &ObligationStore{db: db}
}

// Create persists obligation item.
func (s *ObligationStore) Create(ctx context.Context, o *obDomain.Obligation) error {
	query := `INSERT INTO obligations (id, compliance_id, regulation_id, standard_id, due_date, expiration_date)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, o.ID, o.ComplianceID, o.RegulationID, o.StandardID, o.DueDate, o.ExpirationDate)
	return err
}

// FindByComplianceID returns obligations checklist.
func (s *ObligationStore) FindByComplianceID(ctx context.Context, complianceID string) ([]*obDomain.Obligation, error) {
	query := `SELECT id, compliance_id, regulation_id, standard_id, due_date, expiration_date FROM obligations WHERE compliance_id = $1`
	rows, err := s.db.QueryContext(ctx, query, complianceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*obDomain.Obligation
	for rows.Next() {
		o := &obDomain.Obligation{}
		err = rows.Scan(&o.ID, &o.ComplianceID, &o.RegulationID, &o.StandardID, &o.DueDate, &o.ExpirationDate)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, nil
}
