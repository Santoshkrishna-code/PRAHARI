package scorecard

import (
	"errors"
	"time"
)

// Scorecard aggregates overall ESG scores.
type Scorecard struct {
	ID             string    `json:"id" db:"id"`
	BusinessUnitID string    `json:"business_unit_id" db:"business_unit_id"`
	Period         string    `json:"period" db:"period"` // "2026"
	EnvScore       float64   `json:"env_score" db:"env_score"` // 0-100 scale
	SocialScore    float64   `json:"social_score" db:"social_score"`
	GovScore       float64   `json:"gov_score" db:"gov_score"`
	OverallScore   float64   `json:"overall_score" db:"overall_score"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks targets.
func (s *Scorecard) Validate() error {
	if s.BusinessUnitID == "" {
		return errors.New("business unit reference is required")
	}
	return nil
}

// EvaluateOverall averages indicators.
func (s *Scorecard) EvaluateOverall() {
	s.OverallScore = (s.EnvScore + s.SocialScore + s.GovScore) / 3.0
}
