package hashing

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHasher implements security.Hasher using bcrypt.
type BcryptHasher struct {
	cost int
}

// NewBcryptHasher constructs the Bcrypt hasher wrapper.
func NewBcryptHasher(cost int) *BcryptHasher {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		cost = bcrypt.DefaultCost
	}
	return &BcryptHasher{cost: cost}
}

// Hash generates a bcrypt password hash.
func (h *BcryptHasher) Hash(ctx context.Context, password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", fmt.Errorf("failed to generate bcrypt hash: %w", err)
	}
	return string(hashBytes), nil
}

// Verify checks the password against a bcrypt hash.
func (h *BcryptHasher) Verify(ctx context.Context, password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
