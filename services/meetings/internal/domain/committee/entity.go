package committee

import "time"

// Committee represents a safety committee — a standing body with defined members,
// charter, and meeting schedule.
type Committee struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	Name        string    `json:"name"`
	Charter     string    `json:"charter"`
	ChairID     string    `json:"chair_id"`
	SecretaryID string    `json:"secretary_id"`
	MemberIDs   []string  `json:"member_ids"`
	MeetFreq    string    `json:"meet_freq"` // WEEKLY, BIWEEKLY, MONTHLY, QUARTERLY
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
}
