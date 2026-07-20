package course

import (
	"errors"
)

// Course models training courses.
type Course struct {
	ID                string `json:"id" db:"id"`
	CourseCode        string `json:"course_code" db:"course_code"`
	Title             string `json:"title" db:"title"`
	DurationHours     int    `json:"duration_hours" db:"duration_hours"`
}

// Validate checks domain invariants.
func (c *Course) Validate() error {
	if c.CourseCode == "" {
		return errors.New("course code is required")
	}
	if c.Title == "" {
		return errors.New("course title is required")
	}
	return nil
}
