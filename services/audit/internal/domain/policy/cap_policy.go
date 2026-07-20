package policy

import (
	"errors"
	"time"

	"prahari/services/audit/internal/domain/correctiveaction"
)

// ValidateCorrectiveActionExpiration checks CAPA parameters.
func ValidateCorrectiveActionExpiration(ca *correctiveaction.CorrectiveAction) error {
	if ca.TargetDate.Before(time.Now()) {
		return errors.New("CAPA validity target date must be in future ranges")
	}
	return nil
}
