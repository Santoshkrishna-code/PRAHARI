package api_key

import (
	"errors"
	"time"
)

// APIKey represents scoped token keys.
type APIKey struct {
	ID        string    `json:"id"`
	Hash      string    `json:"-"`
	OwnerID   string    `json:"owner_id"`
	Scopes    []string  `json:"scopes"`
	ExpiresAt time.Time `json:"expires_at"`
}

// IsExpired checks validity.
func (k *APIKey) IsExpired() bool {
	return time.Now().After(k.ExpiresAt)
}

// Validate checks parameters.
func (k *APIKey) Validate() error {
	if k.ID == "" {
		return errors.New("API key ID is required")
	}
	if k.OwnerID == "" {
		return errors.New("owner ID is required")
	}
	return nil
}
