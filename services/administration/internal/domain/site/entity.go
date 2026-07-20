package site

import "time"

// Site represents a geographical region or campus containing plants.
type Site struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id"`
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	CreatedAt      time.Time `json:"created_at"`
}
