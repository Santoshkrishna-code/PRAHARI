package blacklist

import "time"

// Entry registers blacklisted individuals barred from plant access.
type Entry struct {
	ID        string    `json:"id"`
	IDType    string    `json:"id_type"`
	IDNumber  string    `json:"id_number"`
	Reason    string    `json:"reason"`
	BlockedBy string    `json:"blocked_by"`
	BlockedAt time.Time `json:"blocked_at"`
}
