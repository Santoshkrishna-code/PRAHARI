package crewassignment

import "time"

// Assignment maps an operator or user to a crew for a specific time period.
type Assignment struct {
	ID         string    `json:"id"`
	CrewID     string    `json:"crew_id"`
	UserID     string    `json:"user_id"`
	Role       string    `json:"role"` // Panel Operator, Field Operator, Shift Lead, Safety Warden
	AssignedAt time.Time `json:"assigned_at"`
}
