package user

import (
	"errors"
	"strings"
	"time"
)

// User represents the central profile aggregate in the IAM domain.
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	MFASecret    string    `json:"-"`
	MFAEnabled   bool      `json:"mfa_enabled"`
	Role         string    `json:"role"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

// Validate checks user model parameters.
func (u *User) Validate() error {
	if u.ID == "" {
		return errors.New("user ID is required")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email address format")
	}
	if u.Role == "" {
		return errors.New("user role assignment is required")
	}
	return nil
}
