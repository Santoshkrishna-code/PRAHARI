package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// ParseUnverified extracts claims without verifying the cryptographic signature.
// WARNING: Only use this for metadata parsing (e.g. locating kid header values) or inside trusted VPC segments.
func ParseUnverified(tokenStr string) (*Claims, error) {
	parser := jwt.NewParser()
	var claims Claims
	
	_, _, err := parser.ParseUnverified(tokenStr, &claims)
	if err != nil {
		return nil, fmt.Errorf("failed to parse unverified token: %w", err)
	}

	return &claims, nil
}

// ParseWithKey parses and cryptographically validates the token signature using a static key.
func ParseWithKey(tokenStr string, key interface{}) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// Enforce algorithm checks to prevent algorithm confusion attacks (e.g. signing HS256 with public keys)
		if _, ok := t.Method.(*jwt.SigningMethodRSA); ok {
			return key, nil
		}
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return key, nil
		}
		return nil, fmt.Errorf("unsupported signing method: %v", t.Header["alg"])
	})

	if err != nil {
		return nil, fmt.Errorf("token signature verification failed: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
