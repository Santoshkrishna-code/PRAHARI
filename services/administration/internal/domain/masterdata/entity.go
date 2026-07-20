package masterdata

import "time"

// Record represents lookup catalog indexes (e.g. hazard categories, action priorities).
type Record struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"` // HAZARD_TYPE, PRIORITY, EVENT_TYPE
	Code      string    `json:"code"`
	Val       string    `json:"val"`
	Active    bool      `json:"active"`
	UpdatedAt time.Time `json:"updated_at"`
}
