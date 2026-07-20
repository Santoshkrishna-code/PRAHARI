package route

import (
	"errors"
)

// Route maps URL prefixes to downstream targets.
type Route struct {
	Path        string `json:"path"`
	Upstream    string `json:"upstream"`
	StripPrefix bool   `json:"strip_prefix"`
}

// Validate checks model parameters.
func (r *Route) Validate() error {
	if r.Path == "" {
		return errors.New("route matching path prefix is required")
	}
	if r.Upstream == "" {
		return errors.New("route target upstream URL is required")
	}
	return nil
}
