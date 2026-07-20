package safetymeeting

import "time"

// SafetyMeeting represents a formal periodic safety meeting — weekly, monthly, or quarterly
// safety reviews conducted by department heads and safety officers.
type SafetyMeeting struct {
	ID           string    `json:"id"`
	MeetingID    string    `json:"meeting_id"`
	Frequency    string    `json:"frequency"` // WEEKLY, MONTHLY, QUARTERLY
	Department   string    `json:"department"`
	ReviewPeriod string    `json:"review_period"`
	ChairID      string    `json:"chair_id"`
	SecretaryID  string    `json:"secretary_id"`
	CreatedAt    time.Time `json:"created_at"`
}
