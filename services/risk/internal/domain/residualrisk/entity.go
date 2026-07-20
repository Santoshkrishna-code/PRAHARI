package residualrisk

import (
	"errors"
)

// ResidualRisk maps post-mitigation calculations.
type ResidualRisk struct {
	ID                 string `json:"id" db:"id"`
	RiskID             string `json:"risk_id" db:"risk_id"`
	ResidualLikelihood int    `json:"residual_likelihood" db:"residual_likelihood"`
	ResidualConsequence int   `json:"residual_consequence" db:"residual_consequence"`
}

// Validate checks domain invariants.
func (rr *ResidualRisk) Validate() error {
	if rr.RiskID == "" {
		return errors.New("risk ID is required")
	}
	if rr.ResidualLikelihood < 1 || rr.ResidualLikelihood > 5 {
		return errors.New("residual likelihood must fall inside a 1-5 range")
	}
	return nil
}

// Score returns post-mitigation score.
func (rr *ResidualRisk) Score() int {
	return rr.ResidualLikelihood * rr.ResidualConsequence
}
