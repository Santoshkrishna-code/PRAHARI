package jwt_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	prahariJWT "prahari/shared/security/jwt"
)

func TestExtractor_Header(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer token-123")

	token, err := prahariJWT.ExtractTokenFromHeader(req)
	if err != nil {
		t.Fatalf("failed to extract: %v", err)
	}

	if token != "token-123" {
		t.Errorf("expected token-123, got '%s'", token)
	}
}

func TestExtractor_Cookie(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: "cookie-123"})

	token, err := prahariJWT.ExtractTokenFromCookie(req, "session")
	if err != nil {
		t.Fatalf("failed to extract: %v", err)
	}

	if token != "cookie-123" {
		t.Errorf("expected cookie-123, got '%s'", token)
	}
}

func TestValidator_ValidToken(t *testing.T) {
	// 1. Generate RSA key pair for signing and verifying
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate rsa key: %v", err)
	}

	kid := "test-key-id"

	// 2. Setup mock HTTP server returning JWKS containing the public key
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		// Encode modulus and exponent to base64
		n := base64.RawURLEncoding.EncodeToString(privKey.N.Bytes())
		eb := make([]byte, 4)
		binaryBigEndianPut(eb, uint32(privKey.E))
		// Strip leading zeros
		start := 0
		for start < len(eb) && eb[start] == 0 {
			start++
		}
		e := base64.RawURLEncoding.EncodeToString(eb[start:])

		jwks := prahariJWT.JWKS{
			Keys: []prahariJWT.JWK{
				{
					Kty: "RSA",
					Kid: kid,
					Use: "sig",
					Alg: "RS256",
					N:   n,
					E:   e,
				},
			},
		}

		_ = json.NewEncoder(w).Encode(jwks)
	}))
	defer server.Close()

	// 3. Create key resolver and validator pointing to the server
	resolver := prahariJWT.NewKeyResolver(server.URL, 5*time.Minute)
	validator := prahariJWT.NewValidator(resolver, "prahari-auth", "prahari-dashboard")

	// 4. Generate and sign a token using the private key
	claims := &prahariJWT.Claims{
		Email: "worker@prahari.internal",
		Role:  "Worker",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user-123",
			Issuer:    "prahari-auth",
			Audience:  jwt.ClaimStrings{"prahari-dashboard"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = kid

	tokenStr, err := token.SignedString(privKey)
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}

	// 5. Validate the token
	parsedClaims, err := validator.Validate(context.Background(), tokenStr)
	if err != nil {
		t.Fatalf("validation failed: %v", err)
	}

	if parsedClaims.Email != "worker@prahari.internal" {
		t.Errorf("expected email worker@prahari.internal, got %s", parsedClaims.Email)
	}
}

func binaryBigEndianPut(b []byte, v uint32) {
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}
