package riskmatrix

import (
	"errors"
)

// RiskLevel maps classifications.
type RiskLevel string

const (
	LevelLow      RiskLevel = "LOW"
	LevelMedium   RiskLevel = "MEDIUM"
	LevelHigh     RiskLevel = "HIGH"
	LevelCritical RiskLevel = "CRITICAL"
)

// RiskMatrix holds configurable 5x5 scoring rules.
type RiskMatrix struct {
	ID          string `json:"id" db:"id"`
	Likelihood  int    `json:"likelihood" db:"likelihood"`   // 1 to 5
	Consequence int    `json:"consequence" db:"consequence"` // 1 to 5
}

// Validate checks domain invariants.
func (rm *RiskMatrix) Validate() error {
	if rm.Likelihood < 1 || rm.Likelihood > 5 {
		return errors.New("likelihood must fall inside a 1-5 range")
	}
	if rm.Consequence < 1 || rm.Consequence > 5 {
		return errors.New("consequence must fall inside a 1-5 range")
	}
	return nil
}

// CalculateRiskScore multiplies likelihood and consequence.
func (rm *RiskMatrix) CalculateRiskScore() int {
	return rm.Likelihood * rm.Consequence
}

// GetRiskLevel classifies scores to LOW, MEDIUM, HIGH, CRITICAL.
func (rm *RiskMatrix) GetRiskLevel() RiskLevel {
	score := rm.CalculateRiskScore()
	switch {
	case score <= 4:
		return LevelLow
	case score <= 9:
		return LevelMedium
	case score <= 16:
		return LevelHigh
	default:
		return LevelCritical
	}
}
