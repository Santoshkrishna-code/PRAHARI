package csrf

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateToken creates a cryptographically secure 32-byte CSRF token represented as a hex string.
func GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate secure random bytes for CSRF token: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
