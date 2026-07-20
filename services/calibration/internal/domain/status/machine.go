package status

import (
	"fmt"
)

var transitionMatrix = map[Code][]Code{
	CodeRegistered:            {CodeScheduled, CodeRetired},
	CodeScheduled:             {CodeCalibrationStarted, CodeRetired},
	CodeCalibrationStarted:    {CodeMeasurementRecorded, CodeFailed, CodeRetired},
	CodeMeasurementRecorded:   {CodeToleranceVerification, CodeFailed},
	CodeToleranceVerification: {CodeCertificateGenerated, CodeOutOfTolerance, CodeFailed},
	CodeCertificateGenerated:  {CodeApproved, CodeFailed},
	CodeApproved:              {CodeActive, CodeRetired},
	CodeActive:                {CodeScheduled, CodeRetired},
	CodeFailed:                {CodeScheduled, CodeRetired},
	CodeOutOfTolerance:        {CodeScheduled, CodeRetired},
	CodeRetired:               {},
}

// ValidateTransition checks if from → to state transition is allowed in Calibration lifecycle.
func ValidateTransition(from, to Code) error {
	allowed, ok := transitionMatrix[from]
	if !ok {
		return fmt.Errorf("unknown current calibration state: %s", from)
	}
	for _, target := range allowed {
		if target == to {
			return nil
		}
	}
	return fmt.Errorf("invalid calibration state transition from %s to %s", from, to)
}
