package search

import "time"

// Criteria defines multi-dimensional search parameters for meetings.
type Criteria struct {
	PlantID      string     `json:"plant_id,omitempty"`
	MeetingType  string     `json:"meeting_type,omitempty"`
	Status       string     `json:"status,omitempty"`
	OrganizerID  string     `json:"organizer_id,omitempty"`
	DateAfter    *time.Time `json:"date_after,omitempty"`
	DateBefore   *time.Time `json:"date_before,omitempty"`
	Query        string     `json:"query,omitempty"`
	Limit        int        `json:"limit,omitempty"`
	Offset       int        `json:"offset,omitempty"`
}
