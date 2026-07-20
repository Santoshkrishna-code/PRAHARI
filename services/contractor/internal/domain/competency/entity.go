package competency

import (
	"errors"
)

// Competency defines worker skills ratings.
type Competency struct {
	ID          string `json:"id" db:"id"`
	WorkerID    string `json:"worker_id" db:"worker_id"`
	SkillName   string `json:"skill_name" db:"skill_name"`
	SkillLevel  string `json:"skill_level" db:"skill_level"` // Apprentice, Journeyman, Master
}

// Validate checks domain invariants.
func (c *Competency) Validate() error {
	if c.WorkerID == "" {
		return errors.New("worker ID reference is required")
	}
	if c.SkillName == "" {
		return errors.New("competency skill name is required")
	}
	return nil
}
