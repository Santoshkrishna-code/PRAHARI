package events

import (
	"strings"
)

// IsCompatible asserts that the eventVersion shares the same major version index as the requiredVersion.
// This supports backward compatibility checks for schema evolution (e.g. 1.2.0 is compatible with 1.0.0, but 2.0.0 is not).
func IsCompatible(requiredVersion, eventVersion string) bool {
	if requiredVersion == "" || eventVersion == "" {
		return false
	}

	reqParts := strings.Split(requiredVersion, ".")
	evtParts := strings.Split(eventVersion, ".")

	if len(reqParts) == 0 || len(evtParts) == 0 {
		return false
	}

	// Compare major versions
	return reqParts[0] == evtParts[0]
}
