package severity

import (
	"errors"
)

// Severity potential rating scales.
type Severity struct {
	ID          string `json:"id" db:"id"`
	Level       string `json:"level" db:"level"` // Low, Medium, High, Serious
	Score       int    `json:"score" db:"score"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (s *Severity) Validate() error {
	if s.Level == "" {
		return errors.New("severity level classification tag is required")
	}
	return nil
}
