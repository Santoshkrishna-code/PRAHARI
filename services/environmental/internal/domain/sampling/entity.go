package sampling

import (
	"errors"
	"time"
)

// Sampling records sample extraction actions.
type Sampling struct {
	ID             string    `json:"id" db:"id"`
	ProgramID      string    `json:"program_id" db:"program_id"`
	SampleNumber   string    `json:"sample_number" db:"sample_number"`
	SampledBy      string    `json:"sampled_by" db:"sampled_by"`
	LocationCode   string    `json:"location_code" db:"location_code"`
	SampleDate     time.Time `json:"sample_date" db:"sample_date"`
	SampledMedium  string    `json:"sampled_medium" db:"sampled_medium"` // "WATER", "SOIL", "AIR", "EFFLUENT"
	OutcomeStatus  string    `json:"outcome_status" db:"outcome_status"`  // "COLLECTED", "LAB_TRANSIT", "ANALYZED"
	Notes          string    `json:"notes" db:"notes"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks sampling records.
func (s *Sampling) Validate() error {
	if s.ProgramID == "" {
		return errors.New("monitoring program reference is required")
	}
	if s.SampleNumber == "" {
		return errors.New("sample reference code is required")
	}
	return nil
}
