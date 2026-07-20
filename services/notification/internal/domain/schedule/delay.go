package schedule

import (
	"time"
)

// Schedule maps timing requirements scheduling dispatches.
type Schedule struct {
	ID        string    `json:"id"`
	DeliverAt time.Time `json:"deliver_at"`
}
