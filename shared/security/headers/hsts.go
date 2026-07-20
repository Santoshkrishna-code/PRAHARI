package headers

import (
	"fmt"
)

// HSTSConfig holds variables to compile HSTS headers.
type HSTSConfig struct {
	MaxAge            int  `json:"max_age"`
	IncludeSubDomains bool `json:"include_subdomains"`
	Preload           bool `json:"preload"`
}

// Build compiles config into a standard HSTS header string value.
func (c HSTSConfig) Build() string {
	val := fmt.Sprintf("max-age=%d", c.MaxAge)
	if c.IncludeSubDomains {
		val += "; includeSubDomains"
	}
	if c.Preload {
		val += "; preload"
	}
	return val
}
