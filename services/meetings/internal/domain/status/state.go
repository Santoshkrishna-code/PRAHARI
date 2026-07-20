package status

// Code represents meeting lifecycle state.
type Code string

const (
	CodePlanned            Code = "PLANNED"
	CodeScheduled          Code = "SCHEDULED"
	CodeInProgress         Code = "IN_PROGRESS"
	CodeAttendanceRecorded Code = "ATTENDANCE_RECORDED"
	CodeMinutesApproved    Code = "MINUTES_APPROVED"
	CodeActionsGenerated   Code = "ACTIONS_GENERATED"
	CodeClosed             Code = "CLOSED"
	CodeCancelled          Code = "CANCELLED"
	CodeRescheduled        Code = "RESCHEDULED"
)

func (c Code) String() string {
	return string(c)
}
