package session

import (
	"errors"
	"time"
)

// Session represents an active authenticated device session token details.
type Session struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	DeviceFingerprint string    `json:"device_fingerprint"`
	ClientIP          string    `json:"client_ip"`
	LastActivity      time.Time `json:"last_activity"`
	ExpiresAt         time.Time `json:"expires_at"`
}

// IsExpired checks if the device session exceeded timeout limits.
func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

// Validate checks model parameters.
func (s *Session) Validate() error {
	if s.ID == "" {
		return errors.New("session ID is required")
	}
	if s.UserID == "" {
		return errors.New("associated user ID is required")
	}
	return nil
}
