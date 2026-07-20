package status

// Code represents Visitor Management lifecycle state.
type Code string

const (
	CodeScheduled            Code = "SCHEDULED"
	CodeHostApproval         Code = "HOST_APPROVAL"
	CodeSecurityVerification Code = "SECURITY_VERIFICATION"
	CodeGatePassIssued       Code = "GATE_PASS_ISSUED"
	CodeCheckedIn            Code = "CHECKED_IN"
	CodeOnSite               Code = "ON_SITE"
	CodeCheckedOut           Code = "CHECKED_OUT"
	CodeClosed               Code = "CLOSED"
	CodeRejected             Code = "REJECTED"
	CodeCancelled            Code = "CANCELLED"
	CodeBlacklisted          Code = "BLACKLISTED"
)

func (c Code) String() string {
	return string(c)
}
