package incidentcommand

import "time"

// Structure represents an Incident Command System (ICS) structure activated for an emergency.
type Structure struct {
	ID           string    `json:"id"`
	EmergencyID  string    `json:"emergency_id"`
	CommanderID  string    `json:"commander_id"`
	CommandPost  string    `json:"command_post_location"`
	EstablishedAt time.Time `json:"established_at"`
}
