package postgres

import (
	"context"
	"database/sql"

	insuranceDomain "prahari/services/contractor/internal/domain/insurance"
)

// InsuranceStore implements company insurance liability checks.
type InsuranceStore struct {
	db *sql.DB
}

// NewInsuranceStore instantiates InsuranceStore.
func NewInsuranceStore(db *sql.DB) *InsuranceStore {
	return &InsuranceStore{db: db}
}

// Create persists company insurance details.
func (s *InsuranceStore) Create(ctx context.Context, i *insuranceDomain.Insurance) error {
	query := `INSERT INTO contractor_insurance (id, contractor_id, policy_number, expiry_date, limit_amount)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, i.ID, i.ContractorID, i.PolicyNumber, i.ExpiryDate, i.LimitAmount)
	return err
}

// FindByContractorID returns company insurance lists.
func (s *InsuranceStore) FindByContractorID(ctx context.Context, contractorID string) ([]*insuranceDomain.Insurance, error) {
	query := `SELECT id, contractor_id, policy_number, expiry_date, limit_amount FROM contractor_insurance WHERE contractor_id = $1`
	rows, err := s.db.QueryContext(ctx, query, contractorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*insuranceDomain.Insurance
	for rows.Next() {
		i := &insuranceDomain.Insurance{}
		err = rows.Scan(&i.ID, &i.ContractorID, &i.PolicyNumber, &i.ExpiryDate, &i.LimitAmount)
		if err != nil {
			return nil, err
		}
		list = append(list, i)
	}
	return list, nil
}
