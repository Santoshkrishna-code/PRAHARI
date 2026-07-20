package zone

import "time"

// Area represents restricted hazard polygon bounds mapped inside camera fields.
type Area struct {
	ID          string    `json:"id"`
	CameraID    string    `json:"camera_id"`
	Name        string    `json:"name"` // E.g., High Voltage Restricted Zone
	Coordinates string    `json:"coordinates"` // JSON coordinates array E.g., [[x, y], ...]
	UpdatedAt   time.Time `json:"updated_at"`
}
