package curriculum

import (
	"errors"
)

// Curriculum defines frameworks layouts.
type Curriculum struct {
	ID          string `json:"id" db:"id"`
	CourseID    string `json:"course_id" db:"course_id"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (c *Curriculum) Validate() error {
	if c.CourseID == "" {
		return errors.New("course ID reference is required")
	}
	return nil
}
