package headers

import (
	"strings"
)

// CSPBuilder constructs standard Content-Security-Policy strings in a modular fashion.
type CSPBuilder struct {
	directives map[string][]string
}

// NewCSPBuilder instantiates a blank builder.
func NewCSPBuilder() *CSPBuilder {
	return &CSPBuilder{directives: make(map[string][]string)}
}

// Add appends sources to a target directive key.
func (b *CSPBuilder) Add(directive string, sources ...string) *CSPBuilder {
	b.directives[directive] = append(b.directives[directive], sources...)
	return b
}

// Build compiles the policy directives map into a single header string.
func (b *CSPBuilder) Build() string {
	var parts []string
	for dir, src := range b.directives {
		if len(src) > 0 {
			parts = append(parts, dir+" "+strings.Join(src, " "))
		}
	}
	return strings.Join(parts, "; ")
}
