package training

import (
	"errors"
	"time"
)

// Training represents the central aggregate root of the Training & Competency domain.
type Training struct {
	ID             string    `json:"id" db:"id"`
	TrainingNumber string    `json:"training_number" db:"training_number"`
	CourseID       string    `json:"course_id" db:"course_id"`
	DepartmentID   string    `json:"department_id" db:"department_id"`
	StatusCode     string    `json:"status_code" db:"status_code"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted      bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants.
func (t *Training) Validate() error {
	if t.Title == "" {
		return errors.New("training title is required")
	}
	if len(t.Title) > 200 {
		return errors.New("training title must not exceed 200 characters")
	}
	if t.CourseID == "" {
		return errors.New("course ID reference is required")
	}
	if t.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
