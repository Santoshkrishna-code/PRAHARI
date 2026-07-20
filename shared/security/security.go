package security

import (
	"context"
)

// Manager aggregates encryption, decryption, and hashing adapters.
type Manager struct {
	cfg    Config
	cipher Cipher
	hasher Hasher
}

// NewManager constructs a unified security controller.
func NewManager(cfg Config, cipher Cipher, hasher Hasher) *Manager {
	return &Manager{
		cfg:    cfg,
		cipher: cipher,
		hasher: hasher,
	}
}

// Encrypt delegates encryption calls.
func (m *Manager) Encrypt(ctx context.Context, plaintext []byte) ([]byte, error) {
	return m.cipher.Encrypt(ctx, plaintext)
}

// Decrypt delegates decryption calls.
func (m *Manager) Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	return m.cipher.Decrypt(ctx, ciphertext)
}

// HashPassword delegates password hashing algorithms.
func (m *Manager) HashPassword(ctx context.Context, password string) (string, error) {
	return m.hasher.Hash(ctx, password)
}

// VerifyPassword compares password hashes.
func (m *Manager) VerifyPassword(ctx context.Context, password, hash string) (bool, error) {
	return m.hasher.Verify(ctx, password, hash)
}
