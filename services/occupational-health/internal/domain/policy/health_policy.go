package policy

import (
	"errors"

	"prahari/services/occupational-health/internal/domain/exposure"
	"prahari/services/occupational-health/internal/domain/restriction"
)

// EvaluatePermitEligibility validates if active restrictions prevent hot work or specific job tasks.
func EvaluatePermitEligibility(restrictions []restriction.MedicalRestriction, permitType string) error {
	for _, r := range restrictions {
		if r.RestrictionCode == "NO_HEIGHT" && permitType == "WORKING_AT_HEIGHT" {
			return errors.New("worker has active medical restriction preventing working at heights")
		}
		if r.RestrictionCode == "NO_HEAVY_LIFT" && permitType == "HEAVY_LIFTING" {
			return errors.New("worker has active medical restriction preventing heavy lifting operations")
		}
		if r.RestrictionCode == "NO_CHEMICAL" && permitType == "CHEMICAL_HANDLING" {
			return errors.New("worker has active medical restriction preventing chemical agent exposure")
		}
	}
	return nil
}

// EvaluateExposureLimit checks if agent level exceeds statutory limits.
func EvaluateExposureLimit(rec *exposure.ExposureRecord) bool {
	if rec.ExposureLevel > rec.LimitThreshold {
		return true
	}
	return false
}
