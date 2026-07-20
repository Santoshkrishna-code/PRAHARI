package performance

import (
	"errors"
)

// Performance calculates contractor compliance score ratings.
type Performance struct {
	ID              string  `json:"id" db:"id"`
	ContractorID    string  `json:"contractor_id" db:"contractor_id"`
	ComplianceScore float64 `json:"compliance_score" db:"compliance_score"` // 0.0 - 100.0
	SafetyViolations int    `json:"safety_violations" db:"safety_violations"`
}

// Validate checks domain invariants.
func (p *Performance) Validate() error {
	if p.ContractorID == "" {
		return errors.New("contractor ID reference is required")
	}
	return nil
}

// RecordViolation updates scores.
func (p *Performance) RecordViolation() {
	p.SafetyViolations++
	p.ComplianceScore -= 10.0
	if p.ComplianceScore < 0.0 {
		p.ComplianceScore = 0.0
	}
}
