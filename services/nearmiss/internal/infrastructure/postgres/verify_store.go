package postgres

import (
	"context"
	"database/sql"

	verifyDomain "prahari/services/nearmiss/internal/domain/verification"
)

// VerifyStore implements verification check logging.
type VerifyStore struct {
	db *sql.DB
}

// NewVerifyStore instantiates VerifyStore.
func NewVerifyStore(db *sql.DB) *VerifyStore {
	return &VerifyStore{db: db}
}

// Create persists clearance checks.
func (s *VerifyStore) Create(ctx context.Context, v *verifyDomain.Verification) error {
	query := `INSERT INTO near_miss_verifications (id, near_miss_id, verifier_id, verified_date, is_passed, comments)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, v.ID, v.NearMissID, v.VerifierID, v.VerifiedDate, v.IsPassed, v.Comments)
	return err
}

// FindByNearMissID returns verification logs.
func (s *VerifyStore) FindByNearMissID(ctx context.Context, nearmissID string) ([]*verifyDomain.Verification, error) {
	query := `SELECT id, near_miss_id, verifier_id, verified_date, is_passed, comments FROM near_miss_verifications WHERE near_miss_id = $1`
	rows, err := s.db.QueryContext(ctx, query, nearmissID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*verifyDomain.Verification
	for rows.Next() {
		v := &verifyDomain.Verification{}
		err = rows.Scan(&v.ID, &v.NearMissID, &v.VerifierID, &v.VerifiedDate, &v.IsPassed, &v.Comments)
		if err != nil {
			return nil, err
		}
		list = append(list, v)
	}
	return list, nil
}
