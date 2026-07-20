package status

// Code represents Shift Management lifecycle state.
type Code string

const (
	CodeScheduled         Code = "SCHEDULED"
	CodeCrewAssigned      Code = "CREW_ASSIGNED"
	CodeShiftStarted      Code = "SHIFT_STARTED"
	CodeOperational       Code = "OPERATIONAL"
	CodeHandoverInitiated Code = "HANDOVER_INITIATED"
	CodeHandoverAccepted  Code = "HANDOVER_ACCEPTED"
	CodeShiftClosed       Code = "SHIFT_CLOSED"
	CodeCancelled         Code = "CANCELLED"
)

func (c Code) String() string {
	return string(c)
}
