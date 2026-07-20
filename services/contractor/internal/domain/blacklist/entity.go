package blacklist

import (
	"errors"
	"time"
)

// Blacklist tracks blocked contractor worker profiles or companies.
type Blacklist struct {
	ID          string    `json:"id" db:"id"`
	EntityID    string    `json:"entity_id" db:"entity_id"` // WorkerID or ContractorID
	EntityType  string    `json:"entity_type" db:"entity_type"` // WORKER, COMPANY
	Reason      string    `json:"reason" db:"reason"`
	BlacklistedAt time.Time `json:"blacklisted_at" db:"blacklisted_at"`
	ActorID     string    `json:"actor_id" db:"actor_id"`
}

// Validate checks domain invariants.
func (b *Blacklist) Validate() error {
	if b.EntityID == "" {
		return errors.New("entity ID reference is required")
	}
	if b.Reason == "" {
		return errors.New("blacklist reason details are required")
	}
	return nil
}
