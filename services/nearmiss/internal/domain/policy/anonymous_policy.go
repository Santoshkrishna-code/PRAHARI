package policy

import (
	"errors"

	"prahari/services/nearmiss/internal/domain/nearmiss"
)

// ValidateReporterPrivacy preserves user id anonymity checks.
func ValidateReporterPrivacy(nm *nearmiss.NearMiss, isAnonymous bool, userID string) error {
	if isAnonymous && userID != "" {
		return errors.New("anonymous near miss reports must not expose user ID references")
	}
	return nil
}
