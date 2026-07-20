package businessunit

import "time"

// BusinessUnit represents a logical division or sector of the organization.
type BusinessUnit struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	LeaderID       string    `json:"leader_id"`
	CreatedAt      time.Time `json:"created_at"`
}
