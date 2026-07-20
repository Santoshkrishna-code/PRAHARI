package preference

import (
	"errors"
)

// Preference maps user-level channel exclusions.
type Preference struct {
	UserID  string `json:"user_id"`
	Channel string `json:"channel"`
	Enabled bool   `json:"enabled"`
}

// Validate checks fields.
func (p *Preference) Validate() error {
	if p.UserID == "" {
		return errors.New("user ID is required")
	}
	if p.Channel == "" {
		return errors.New("channel type designation is required")
	}
	return nil
}
