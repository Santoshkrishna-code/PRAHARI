package policy

import (
	"errors"

	"prahari/services/observation/internal/domain/observation"
)

// CheckBBSCoachingRequirements checks if unsafe behavior triggers coaching session plan targets.
func CheckBBSCoachingRequirements(o *observation.Observation) error {
	if o.ObservationType == "UNSAFE_BEHAVIOR" && o.StatusCode == "RECORDED" {
		return errors.New("unsafe behaviors require safety coaching dialogues to enforce EHS culture")
	}
	return nil
}
