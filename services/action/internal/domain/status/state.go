package status

// Code represents LOTO lifecycle state.
type Code string

const (
	CodeCreated             Code = "CREATED"
	CodeAssigned            Code = "ASSIGNED"
	CodeInProgress          Code = "IN_PROGRESS"
	CodeEvidenceSubmitted   Code = "EVIDENCE_SUBMITTED"
	CodeEffectivenessReview Code = "EFFECTIVENESS_REVIEW"
	CodeClosed              Code = "CLOSED"
	CodeCancelled           Code = "CANCELLED"
	CodeOverdue             Code = "OVERDUE"
	CodeRejected            Code = "REJECTED"
)

func (c Code) String() string {
	return string(c)
}
