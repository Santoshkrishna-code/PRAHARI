package postgres

import (
	"context"
	"database/sql"

	residualDomain "prahari/services/risk/internal/domain/residualrisk"
)

// ResidualStore implements post-mitigation scores database.
type ResidualStore struct {
	db *sql.DB
}

// NewResidualStore instantiates ResidualStore.
func NewResidualStore(db *sql.DB) *ResidualStore {
	return &ResidualStore{db: db}
}

// Create persists residual score values.
func (s *ResidualStore) Create(ctx context.Context, rr *residualDomain.ResidualRisk) error {
	query := `INSERT INTO risk_residual (id, risk_id, residual_likelihood, residual_consequence)
		VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, rr.ID, rr.RiskID, rr.ResidualLikelihood, rr.ResidualConsequence)
	return err
}

// FindByRiskID returns residual logs.
func (s *ResidualStore) FindByRiskID(ctx context.Context, riskID string) ([]*residualDomain.ResidualRisk, error) {
	query := `SELECT id, risk_id, residual_likelihood, residual_consequence FROM risk_residual WHERE risk_id = $1`
	rows, err := s.db.QueryContext(ctx, query, riskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*residualDomain.ResidualRisk
	for rows.Next() {
		rr := &residualDomain.ResidualRisk{}
		err = rows.Scan(&rr.ID, &rr.RiskID, &rr.ResidualLikelihood, &rr.ResidualConsequence)
		if err != nil {
			return nil, err
		}
		list = append(list, rr)
	}
	return list, nil
}
