package compliance

import (
	"prahari/services/chemical/internal/domain/chemical"
)

// ValidateREACH checks if chemical complies with EU REACH annexes.
func ValidateREACH(c *chemical.Chemical) (bool, string) {
	if c == nil {
		return false, "Chemical definition not found"
	}
	if c.IsRestricted {
		return false, "Chemical is on REACH Annex XIV authorization list"
	}
	return true, "Complies with REACH regulations"
}
