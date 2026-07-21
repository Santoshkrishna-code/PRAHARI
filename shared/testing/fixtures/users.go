package fixtures

import (
	"github.com/golang-jwt/jwt/v5"
	prahariJWT "prahari/shared/security/jwt"
)

// NewAdminClaims returns standard Admin credentials.
func NewAdminClaims() *prahariJWT.Claims {
	return &prahariJWT.Claims{
		Role: "Admin",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: "usr-admin-99",
		},
	}
}

// NewWorkerClaims returns standard Regular credentials.
func NewWorkerClaims() *prahariJWT.Claims {
	return &prahariJWT.Claims{
		Role: "Worker",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: "usr-worker-11",
		},
	}
}
