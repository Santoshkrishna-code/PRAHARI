package jwt

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
)

// JWK represents a single JSON Web Key block.
type JWK struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// JWKS holds list of JWK public keys.
type JWKS struct {
	Keys []JWK `json:"keys"`
}

// KeyResolver downloads and caches JWKS keys to validate JWT signatures.
type KeyResolver struct {
	jwksURL string
	keys    map[string]*rsa.PublicKey
	mu      sync.RWMutex
	client  *http.Client
	ttl     time.Duration
	expiry  time.Time
}

// NewKeyResolver constructs a new KeyResolver configuration.
func NewKeyResolver(jwksURL string, ttl time.Duration) *KeyResolver {
	return &KeyResolver{
		jwksURL: jwksURL,
		keys:    make(map[string]*rsa.PublicKey),
		client:  &http.Client{Timeout: 5 * time.Second},
		ttl:     ttl,
	}
}

// GetPublicKey retrieves the RSA public key matching the key ID (kid).
func (r *KeyResolver) GetPublicKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	r.mu.RLock()
	key, exists := r.keys[kid]
	isExpired := time.Now().After(r.expiry)
	r.mu.RUnlock()

	if exists && !isExpired {
		return key, nil
	}

	// Cache miss or expired keys: reload JWKS
	if err := r.refreshKeys(ctx); err != nil {
		// Fall back to expired cached key if network is temporarily down
		if exists {
			return key, nil
		}
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()
	key, exists = r.keys[kid]
	if !exists {
		return nil, fmt.Errorf("key %s not found in JWKS from %s", kid, r.jwksURL)
	}

	return key, nil
}

func (r *KeyResolver) refreshKeys(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.jwksURL, nil)
	if err != nil {
		return err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch JWKS keys: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("JWKS endpoint returned status code: %d", resp.StatusCode)
	}

	var jwks JWKS
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return fmt.Errorf("failed to decode JWKS: %w", err)
	}

	newKeys := make(map[string]*rsa.PublicKey)
	for _, key := range jwks.Keys {
		if key.Kty != "RSA" {
			continue
		}

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

		newKeys[key.Kid] = pubKey
	}

	r.keys = newKeys
	r.expiry = time.Now().Add(r.ttl)
	return nil
}
