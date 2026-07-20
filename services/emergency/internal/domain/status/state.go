package status

// Code represents industrial emergency state.
type Code string

const (
	CodePrepared           Code = "PREPARED"
	CodeDetected           Code = "DETECTED"
	CodeDeclared           Code = "DECLARED"
	CodeResponseActivated  Code = "RESPONSE_ACTIVATED"
	CodeCommandEstablished Code = "COMMAND_ESTABLISHED"
	CodeResourceDeployment Code = "RESOURCE_DEPLOYMENT"
	CodeEvacuation         Code = "EVACUATION"
	CodeStabilized         Code = "STABILIZED"
	CodeRecovery           Code = "RECOVERY"
	CodeAfterActionReview  Code = "AFTER_ACTION_REVIEW"
	CodeClosed             Code = "CLOSED"
	CodeCancelled          Code = "CANCELLED"
)

func (c Code) String() string {
	return string(c)
}
