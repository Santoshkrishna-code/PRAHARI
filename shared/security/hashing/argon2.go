package hashing

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2Hasher implements security.Hasher using Argon2id.
type Argon2Hasher struct {
	memory      uint32 // in KB (e.g. 65536 for 64MB)
	iterations  uint32
	parallelism uint8
	saltSize    uint32
	keySize     uint32
}

// NewArgon2Hasher constructs the Argon2id hasher wrapper.
func NewArgon2Hasher(memory, iterations uint32, parallelism uint8) *Argon2Hasher {
	return &Argon2Hasher{
		memory:      memory,
		iterations:  iterations,
		parallelism: parallelism,
		saltSize:    16,
		keySize:     32,
	}
}

// Hash generates an Argon2id self-descriptive hash string.
func (h *Argon2Hasher) Hash(ctx context.Context, password string) (string, error) {
	salt := make([]byte, h.saltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate random salt: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		h.iterations,
		h.memory,
		h.parallelism,
		h.keySize,
	)

	// Format matching standard Argon2 formats: $argon2id$v=19$m=65536,t=3,p=4$salt$hash
	saltBase64 := base64.RawStdEncoding.EncodeToString(salt)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		h.memory,
		h.iterations,
		h.parallelism,
		saltBase64,
		hashBase64,
	)

	return encoded, nil
}

// Verify decodes parameters from the encoded hash and compares hashes in constant-time.
func (h *Argon2Hasher) Verify(ctx context.Context, password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid argon2 hash format")
	}

	if parts[1] != "argon2id" {
		return false, fmt.Errorf("unsupported hashing method: %s", parts[1])
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return false, fmt.Errorf("failed to parse version: %w", err)
	}
	if version != argon2.Version {
		return false, fmt.Errorf("incompatible argon2 version: %d", version)
	}

	var memory, iterations uint32
	var parallelism uint8
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return false, fmt.Errorf("failed to parse execution parameters: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("failed to decode hash bytes: %w", err)
	}

	actualHash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		uint32(len(expectedHash)),
	)

	// Use subtle.ConstantTimeCompare to mitigate timing side-channel attacks
	if subtle.ConstantTimeCompare(actualHash, expectedHash) == 1 {
		return true, nil
	}

	return false, nil
}
