package trainingverification

import "time"

// Record tracks workforce competency verification for modified processes or assets.
type Record struct {
	ID              string    `json:"id"`
	ChangeRequestID string    `json:"change_request_id"`
	CourseID        string    `json:"course_id,omitempty"`
	TrainedCount    int       `json:"trained_count"`
	TargetCount     int       `json:"target_count"`
	IsComplete      bool      `json:"is_complete"`
	VerifiedBy      string    `json:"verified_by"`
	VerifiedAt      time.Time `json:"verified_at"`
}
