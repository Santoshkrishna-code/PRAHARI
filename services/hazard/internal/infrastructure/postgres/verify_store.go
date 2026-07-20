package postgres

import (
	"context"
	"database/sql"

	verifyDomain "prahari/services/hazard/internal/domain/verification"
)

// VerifyStore implements verification check logging.
type VerifyStore struct {
	db *sql.DB
}

// NewVerifyStore instantiates VerifyStore.
func NewVerifyStore(db *sql.DB) *VerifyStore {
	return &VerifyStore{db: db}
}

// Create persists physical clearance checks.
func (s *VerifyStore) Create(ctx context.Context, v *verifyDomain.Verification) error {
	query := `INSERT INTO hazard_verifications (id, hazard_id, verifier_id, verified_date, residual_risk_score, comments)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, v.ID, v.HazardID, v.VerifierID, v.VerifiedDate, v.ResidualRiskScore, v.Comments)
	return err
}

// FindByHazardID returns verification logs.
func (s *VerifyStore) FindByHazardID(ctx context.Context, hazardID string) ([]*verifyDomain.Verification, error) {
	query := `SELECT id, hazard_id, verifier_id, verified_date, residual_risk_score, comments FROM hazard_verifications WHERE hazard_id = $1`
	rows, err := s.db.QueryContext(ctx, query, hazardID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*verifyDomain.Verification
	for rows.Next() {
		v := &verifyDomain.Verification{}
		err = rows.Scan(&v.ID, &v.HazardID, &v.VerifierID, &v.VerifiedDate, &v.ResidualRiskScore, &v.Comments)
		if err != nil {
			return nil, err
		}
		list = append(list, v)
	}
	return list, nil
}
