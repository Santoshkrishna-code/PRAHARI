package riskassessment

import (
	"encoding/json"
	"errors"
	"time"

	"prahari/services/permit/internal/domain/permit"
)

// RiskAssessment represents a safety score evaluation.
type RiskAssessment struct {
	ID               string          `json:"id" db:"id"`
	PermitID         string          `json:"permit_id" db:"permit_id"`
	AssessorID       string          `json:"assessor_id" db:"assessor_id"`
	LikelihoodScore  int             `json:"likelihood_score" db:"likelihood_score"` // 1 to 5
	ConsequenceScore int             `json:"consequence_score" db:"consequence_score"` // 1 to 5
	RiskScore        int             `json:"risk_score" db:"risk_score"` // L * C
	RiskLevel        permit.RiskLevel `json:"risk_level" db:"risk_level"`
	ControlMeasures  json.RawMessage `json:"control_measures" db:"control_measures"` // JSON array of strings
	ResidualRisk     permit.RiskLevel `json:"residual_risk" db:"residual_risk"`
	PPERequired      json.RawMessage `json:"ppe_required" db:"ppe_required"` // JSON array of strings
	AssessedAt       time.Time       `json:"assessed_at" db:"assessed_at"`
}

// Validate checks domain invariants for RiskAssessment.
func (ra *RiskAssessment) Validate() error {
	if ra.PermitID == "" {
		return errors.New("permit ID is required for risk assessment")
	}
	if ra.AssessorID == "" {
		return errors.New("assessor ID is required")
	}
	if ra.LikelihoodScore < 1 || ra.LikelihoodScore > 5 {
		return errors.New("likelihood score must be between 1 and 5")
	}
	if ra.ConsequenceScore < 1 || ra.ConsequenceScore > 5 {
		return errors.New("consequence score must be between 1 and 5")
	}
	return nil
}

// CalculateRiskScore computes the basic hazard severity level.
func (ra *RiskAssessment) CalculateRiskScore() {
	ra.RiskScore = ra.LikelihoodScore * ra.ConsequenceScore
	ra.RiskLevel = ra.DeriveRiskLevel(ra.RiskScore)
}

// DeriveRiskLevel returns a RiskLevel string corresponding to a score value.
func (ra *RiskAssessment) DeriveRiskLevel(score int) permit.RiskLevel {
	switch {
	case score >= 16:
		return permit.RiskLevelCritical
	case score >= 10:
		return permit.RiskLevelHigh
	case score >= 5:
		return permit.RiskLevelMedium
	default:
		return permit.RiskLevelLow
	}
}
