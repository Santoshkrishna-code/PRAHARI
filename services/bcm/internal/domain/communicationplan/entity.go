package communicationplan

import "time"

// Plan defines crisis communication protocols for executive, media, regulatory, and workforce stakeholders.
type Plan struct {
	ID          string    `json:"id"`
	PlanID      string    `json:"plan_id"`
	Spokesperson string    `json:"spokesperson_id"`
	Channels    string    `json:"channels"`
	Templates   string    `json:"templates"`
	CreatedAt   time.Time `json:"created_at"`
}
