package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims represents the standard and custom JWT attributes mapped across PRAHARI microservices.
type Claims struct {
	Email          string   `json:"email"`
	Role           string   `json:"role"`
	TenantID       string   `json:"tenant_id,omitempty"`
	OrganizationID string   `json:"org_id,omitempty"`
	Permissions    []string `json:"permissions,omitempty"`
	jwt.RegisteredClaims
}
