package training

// CreateTrainingCommand carries course parameters.
type CreateTrainingCommand struct {
	CourseID     string `json:"course_id"`
	DepartmentID string `json:"department_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// TransitionStatusCommand carries status updates.
type TransitionStatusCommand struct {
	TrainingID string `json:"training_id"`
	TargetCode string `json:"target_code"`
	ActorID    string `json:"actor_id"`
	Reason     string `json:"reason,omitempty"`
}
