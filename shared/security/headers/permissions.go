package headers

import (
	"strings"
)

// PermissionsPolicyBuilder constructs modern HTTP Permissions-Policy headers.
type PermissionsPolicyBuilder struct {
	policies map[string][]string
}

// NewPermissionsPolicyBuilder instantiates a blank builder.
func NewPermissionsPolicyBuilder() *PermissionsPolicyBuilder {
	return &PermissionsPolicyBuilder{policies: make(map[string][]string)}
}

// Add binds allowances list to feature keys (e.g. camera, geolocation).
func (b *PermissionsPolicyBuilder) Add(feature string, allows ...string) *PermissionsPolicyBuilder {
	b.policies[feature] = append(b.policies[feature], allows...)
	return b
}

// Build compiles maps into the format: feature1=(allow1 allow2), feature2=(allow1)
func (b *PermissionsPolicyBuilder) Build() string {
	var parts []string
	for feat, allow := range b.policies {
		if len(allow) > 0 {
			parts = append(parts, feat+"=("+strings.Join(allow, " ")+")")
		}
	}
	return strings.Join(parts, ", ")
}
