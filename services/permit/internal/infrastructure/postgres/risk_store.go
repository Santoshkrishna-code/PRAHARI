package postgres

import (
	"context"
	"database/sql"

	riskDomain "prahari/services/permit/internal/domain/riskassessment"
)

// RiskStore implements risk score card persistency.
type RiskStore struct {
	db *sql.DB
}

// NewRiskStore instantiates RiskStore.
func NewRiskStore(db *sql.DB) *RiskStore {
	return &RiskStore{db: db}
}

// Create saves card values.
func (s *RiskStore) Create(ctx context.Context, ra *riskDomain.RiskAssessment) error {
	query := `INSERT INTO permit_risk_assessments (id, permit_id, assessor_id, likelihood_score, consequence_score, risk_score, risk_level, control_measures, residual_risk, ppe_required, assessed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := s.db.ExecContext(ctx, query,
		ra.ID, ra.PermitID, ra.AssessorID, ra.LikelihoodScore, ra.ConsequenceScore, ra.RiskScore, ra.RiskLevel, ra.ControlMeasures, ra.ResidualRisk, ra.PPERequired, ra.AssessedAt,
	)
	return err
}

// FindByPermitID returns evaluations.
func (s *RiskStore) FindByPermitID(ctx context.Context, permitID string) (*riskDomain.RiskAssessment, error) {
	query := `SELECT id, permit_id, assessor_id, likelihood_score, consequence_score, risk_score, risk_level, control_measures, residual_risk, ppe_required, assessed_at
		FROM permit_risk_assessments WHERE permit_id = $1`
	ra := &riskDomain.RiskAssessment{}
	err := s.db.QueryRowContext(ctx, query, permitID).Scan(
		&ra.ID, &ra.PermitID, &ra.AssessorID, &ra.LikelihoodScore, &ra.ConsequenceScore, &ra.RiskScore, &ra.RiskLevel, &ra.ControlMeasures, &ra.ResidualRisk, &ra.PPERequired, &ra.AssessedAt,
	)
	if err != nil {
		return nil, err
	}
	return ra, nil
}
