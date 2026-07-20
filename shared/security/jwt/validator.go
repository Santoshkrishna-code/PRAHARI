package jwt

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// Validator orchestrates token signature checks and audience/issuer verifications.
type Validator struct {
	resolver *KeyResolver
	issuer   string
	audience string
}

// NewValidator constructs a token Validator instance.
func NewValidator(resolver *KeyResolver, issuer, audience string) *Validator {
	return &Validator{
		resolver: resolver,
		issuer:   issuer,
		audience: audience,
	}
}

// Validate parses, decrypts, and cryptographically asserts token parameters.
func (v *Validator) Validate(ctx context.Context, tokenStr string) (*Claims, error) {
	// Parse with claims dynamically resolving keys via kid header
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing kid header")
		}

		return v.resolver.GetPublicKey(ctx, kid)
	})

	if err != nil {
		return nil, fmt.Errorf("jwt validation failed: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("jwt token is invalid")
	}

	// 1. Verify Issuer
	if v.issuer != "" && claims.Issuer != v.issuer {
		return nil, fmt.Errorf("issuer mismatch: expected %s, got %s", v.issuer, claims.Issuer)
	}

	// 2. Verify Audience
	if v.audience != "" {
		audList, err := claims.GetAudience()
		if err != nil {
			return nil, fmt.Errorf("failed to parse audience list: %w", err)
		}
		found := false
		for _, aud := range audList {
			if aud == v.audience {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("audience mismatch: expected %s", v.audience)
		}
	}

	return claims, nil
}
