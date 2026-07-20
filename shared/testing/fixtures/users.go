package fixtures

import (
	prahariJWT "prahari/shared/security/jwt"
)

// NewAdminClaims returns standard Admin credentials.
func NewAdminClaims() *prahariJWT.Claims {
	return &prahariJWT.Claims{
		UserID: "usr-admin-99",
		Role:   "Admin",
	}
}

// NewWorkerClaims returns standard Regular credentials.
func NewWorkerClaims() *prahariJWT.Claims {
	return &prahariJWT.Claims{
		UserID: "usr-worker-11",
		Role:   "Worker",
	}
}
