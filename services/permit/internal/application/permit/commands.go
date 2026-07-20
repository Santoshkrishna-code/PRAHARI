package permit

import (
	"time"
)

// CreatePermitCommand carries data required to construct a new permit request.
type CreatePermitCommand struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	PermitTypeID    string    `json:"permit_type_id"`
	ApplicantID     string    `json:"applicant_id"`
	SupervisorID    string    `json:"supervisor_id"`
	DepartmentID    string    `json:"department_id"`
	ContractorID    string    `json:"contractor_id,omitempty"`
	WorkAreaID      string    `json:"work_area_id"`
	WorkDescription string    `json:"work_description"`
	PlannedStartAt  time.Time `json:"planned_start_at"`
	PlannedEndAt    time.Time `json:"planned_end_at"`
}

// UpdatePermitCommand carries fields that can be modified on a draft/pre-approval permit.
type UpdatePermitCommand struct {
	Title           string    `json:"title,omitempty"`
	Description     string    `json:"description,omitempty"`
	WorkDescription string    `json:"work_description,omitempty"`
	PlannedStartAt  time.Time `json:"planned_start_at,omitempty"`
	PlannedEndAt    time.Time `json:"planned_end_at,omitempty"`
}

// TransitionStatusCommand carries inputs for a manual or workflow status step change.
type TransitionStatusCommand struct {
	PermitID   string `json:"permit_id"`
	TargetCode string `json:"target_code"`
	ActorID    string `json:"actor_id"`
	Reason     string `json:"reason,omitempty"`
}
