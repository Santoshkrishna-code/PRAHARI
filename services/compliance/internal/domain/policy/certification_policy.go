package policy

import (
	"errors"
	"time"

	"prahari/services/compliance/internal/domain/certification"
)

// ValidateCertificationExpiration checks license parameters.
func ValidateCertificationExpiration(c *certification.Certification) error {
	if c.ValidUntil.Before(time.Now()) {
		return errors.New("certification validity timestamp must be in future ranges")
	}
	return nil
}
