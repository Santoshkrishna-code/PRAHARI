package utils

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"prahari/services/auth-service/internal/domain"
)

type JWK struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

// TokenVerifier validates Cognito tokens against cached JWKs.
type TokenVerifier struct {
	jwksURL string
	issuer  string
	keys    map[string]*rsa.PublicKey
	mu      sync.RWMutex
	client  *http.Client
}

func NewTokenVerifier(jwksURL, issuer string) *TokenVerifier {
	return &TokenVerifier{
		jwksURL: jwksURL,
		issuer:  issuer,
		keys:    make(map[string]*rsa.PublicKey),
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

// Verify decodes and validates a Cognito access token, returning its claims.
func (v *TokenVerifier) Verify(ctx context.Context, tokenStr string) (*domain.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 1. Verify signing algorithm matches RSA
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("missing kid header key")
		}
		
		// 2. Fetch RSA public key matching kid
		return v.getPublicKey(ctx, kid)
	})
	
	if err != nil {
		if err == jwt.ErrTokenExpired {
			return nil, domain.ErrTokenExpired
		}
		return nil, fmt.Errorf("%w: %s", domain.ErrInvalidToken, err.Error())
	}
	
	claims, ok := token.Claims.(*domain.JWTClaims)
	if !ok || !token.Valid {
		return nil, domain.ErrInvalidToken
	}
	
	// 3. Verify issuer matches Cognito pool
	if claims.Issuer != v.issuer {
		return nil, fmt.Errorf("%w: invalid issuer claim", domain.ErrInvalidToken)
	}
	
	return claims, nil
}

func (v *TokenVerifier) getPublicKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	v.mu.RLock()
	key, exists := v.keys[kid]
	v.mu.RUnlock()
	if exists {
		return key, nil
	}
	
	// Cache miss: fetch JWKS from Cognito
	if err := v.refreshKeys(ctx); err != nil {
		return nil, err
	}
	
	v.mu.RLock()
	key, exists = v.keys[kid]
	v.mu.RUnlock()
	if !exists {
		return nil, fmt.Errorf("key %s not found in Cognito JWKS set", kid)
	}
	
	return key, nil
}

func (v *TokenVerifier) refreshKeys(ctx context.Context) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, v.jwksURL, nil)
	if err != nil {
		return err
	}
	
	resp, err := v.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch JWKS keys: %w", err)
	}
	defer resp.Body.Close()
	
	var jwks JWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return fmt.Errorf("failed to decode JWKS: %w", err)
	}
	
	for _, key := range jwks.Keys {
		if key.Kty != "RSA" {
			continue
		}
		
		// Decode modulus and exponent components to RSA Public Key
		nb, err := base64.RawURLEncoding.DecodeString(key.N)
		if err != nil {
			continue
		}
		eb, err := base64.RawURLEncoding.DecodeString(key.E)
		if err != nil {
			continue
		}
		
		var val int
		for _, b := range eb {
			val = (val << 8) + int(b)
		}
		
		pubKey := &rsa.PublicKey{
			N: new(big.Int).SetBytes(nb),
			E: val,
		}
		
		v.keys[key.Kid] = pubKey
	}
	
	return nil
}
