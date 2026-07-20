package service_account

import (
	"errors"
)

// ServiceAccount maps machine-to-machine client connections profiles.
type ServiceAccount struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	ClientID  string   `json:"client_id"`
	Scopes    []string `json:"scopes"`
	Suspended bool     `json:"suspended"`
}

// Validate checks model parameters.
func (s *ServiceAccount) Validate() error {
	if s.ID == "" {
		return errors.New("service account ID is required")
	}
	if s.ClientID == "" {
		return errors.New("client ID is required")
	}
	return nil
}
