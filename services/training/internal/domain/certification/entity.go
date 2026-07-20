package certification

import (
	"errors"
	"time"
)

// Certification tracks completed course credentials.
type Certification struct {
	ID         string    `json:"id" db:"id"`
	TraineeID  string    `json:"trainee_id" db:"trainee_id"`
	CourseID   string    `json:"course_id" db:"course_id"`
	Issuer     string    `json:"issuer" db:"issuer"`
	ValidUntil time.Time `json:"valid_until" db:"valid_until"`
}

// Validate checks domain invariants.
func (c *Certification) Validate() error {
	if c.TraineeID == "" {
		return errors.New("trainee ID reference is required")
	}
	if c.CourseID == "" {
		return errors.New("course ID reference is required")
	}
	return nil
}
