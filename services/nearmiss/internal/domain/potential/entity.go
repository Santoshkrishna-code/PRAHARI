package potential

import (
	"errors"
)

// Potential risk potential assessments classifications.
type Potential struct {
	ID              string `json:"id" db:"id"`
	NearMissID      string `json:"near_miss_id" db:"near_miss_id"`
	SeverityScore   int    `json:"severity_score" db:"severity_score"`
	LikelihoodScore int    `json:"likelihood_score" db:"likelihood_score"`
}

// Validate checks domain invariants.
func (p *Potential) Validate() error {
	if p.NearMissID == "" {
		return errors.New("near miss ID reference is required")
	}
	return nil
}

// RiskScore evaluates potential rating levels.
func (p *Potential) RiskScore() int {
	return p.SeverityScore * p.LikelihoodScore
}
