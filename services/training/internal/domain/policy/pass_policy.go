package policy

import (
	"errors"

	"prahari/services/training/internal/domain/result"
)

// ValidatePassingScore checks scores metrics.
func ValidatePassingScore(r *result.Result) error {
	if r.Score < 60 {
		return errors.New("passing score requirements is at least 60%")
	}
	return nil
}
