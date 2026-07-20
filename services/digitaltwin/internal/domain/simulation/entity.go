package simulation

import "time"

// Scenario represents a what-if simulation run against a digital twin.
type Scenario struct {
	ID          string    `json:"id"`
	TwinID      string    `json:"twin_id"`
	Name        string    `json:"name"` // E.g., Reactor Trip Simulation
	Status      string    `json:"status"` // PENDING, RUNNING, COMPLETED, FAILED
	Parameters  string    `json:"parameters"` // JSON parameters
	ResultData  string    `json:"result_data,omitempty"` // JSON results
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}
