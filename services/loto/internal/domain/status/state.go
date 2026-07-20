package status

// Code represents LOTO lifecycle state.
type Code string

const (
	CodePlanned               Code = "PLANNED"
	CodeIsolationApproved     Code = "ISOLATION_APPROVED"
	CodeLocksApplied          Code = "LOCKS_APPLIED"
	CodeTagsApplied           Code = "TAGS_APPLIED"
	CodeZeroEnergyVerified    Code = "ZERO_ENERGY_VERIFIED"
	CodeMaintenanceInProgress Code = "MAINTENANCE_IN_PROGRESS"
	CodeRestorationPlanned    Code = "RESTORATION_PLANNED"
	CodeLocksRemoved          Code = "LOCKS_REMOVED"
	CodeReturnedToService     Code = "RETURNED_TO_SERVICE"
	CodeClosed                Code = "CLOSED"
	CodeCancelled             Code = "CANCELLED"
	CodeFailedVerification    Code = "FAILED_VERIFICATION"
)

func (c Code) String() string {
	return string(c)
}
