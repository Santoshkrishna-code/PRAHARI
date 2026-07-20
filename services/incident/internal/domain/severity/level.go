package severity

import (
	"fmt"
)

// Level represents the severity classification of an incident.
type Level string

const (
	LevelCritical      Level = "CRITICAL"
	LevelHigh          Level = "HIGH"
	LevelMedium        Level = "MEDIUM"
	LevelLow           Level = "LOW"
	LevelInformational Level = "INFORMATIONAL"
)

// SLAResponseHours maps each severity level to its required response time target in hours.
var SLAResponseHours = map[Level]int{
	LevelCritical:      1,
	LevelHigh:          4,
	LevelMedium:        24,
	LevelLow:           72,
	LevelInformational: 168,
}

// ValidLevels enumerates all accepted severity levels.
var ValidLevels = []Level{
	LevelCritical,
	LevelHigh,
	LevelMedium,
	LevelLow,
	LevelInformational,
}

// Validate checks whether the level is among accepted severity classifications.
func (l Level) Validate() error {
	for _, valid := range ValidLevels {
		if l == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid severity level: %s", l)
}

// GetSLAHours returns the SLA response time target in hours for this severity level.
func (l Level) GetSLAHours() int {
	if hours, exists := SLAResponseHours[l]; exists {
		return hours
	}
	return 168 // Default to informational SLA
}

// IsEscalationRequired returns true if the severity demands immediate attention.
func (l Level) IsEscalationRequired() bool {
	return l == LevelCritical || l == LevelHigh
}
