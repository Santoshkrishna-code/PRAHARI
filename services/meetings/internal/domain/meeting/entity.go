package meeting

import "time"

// Meeting is the root aggregate representing any type of safety meeting.
type Meeting struct {
	ID           string     `json:"id"`
	PlantID      string     `json:"plant_id"`
	MeetingType  string     `json:"meeting_type"` // TOOLBOX_TALK, SHIFT_BRIEFING, PREJOB_BRIEFING, SAFETY_COMMITTEE, GENERAL
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Location     string     `json:"location"`
	ScheduledAt  time.Time  `json:"scheduled_at"`
	StartedAt    *time.Time `json:"started_at,omitempty"`
	EndedAt      *time.Time `json:"ended_at,omitempty"`
	OrganizerID  string     `json:"organizer_id"`
	FacilitatorID string   `json:"facilitator_id"`
	ShiftID      string     `json:"shift_id,omitempty"`
	PermitID     string     `json:"permit_id,omitempty"`
	Status       string     `json:"status"`
	DurationMin  int        `json:"duration_min"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
