package random

import (
	"crypto/rand"
	"fmt"
)

// Bytes generates a cryptographically secure random slice of bytes of the specified size.
func Bytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("failed to generate secure random bytes: %w", err)
	}
	return b, nil
}
