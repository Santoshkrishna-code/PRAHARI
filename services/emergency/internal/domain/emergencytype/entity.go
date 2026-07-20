package emergencytype

import "time"

// Type defines standardized emergency classification metadata.
type Type struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"` // E.g., FIRE_ALKYLATION, TOXIC_H2S
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Severity    string    `json:"severity"`
	CreatedAt   time.Time `json:"created_at"`
}
