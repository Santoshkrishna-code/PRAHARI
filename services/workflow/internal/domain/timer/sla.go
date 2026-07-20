package timer

import (
	"errors"
)

// SLAPolicy represents deadline limits and escalation routes.
type SLAPolicy struct {
	TaskID      string `json:"task_id"`
	DurationMin int    `json:"duration_min"`
	EscalateTo  string `json:"escalate_to"`
}

// Validate checks fields.
func (p *SLAPolicy) Validate() error {
	if p.TaskID == "" {
		return errors.New("SLA policy task ID is required")
	}
	if p.DurationMin <= 0 {
		return errors.New("SLA timeout duration must be positive")
	}
	return nil
}
