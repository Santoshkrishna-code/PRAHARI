package encryption

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// AESCipher implements security.Cipher using AES-256 GCM mode.
type AESCipher struct {
	key []byte
}

// NewAESCipher instantiates the AES-256 GCM cipher wrapper.
// Key must be exactly 32 bytes (256 bits).
func NewAESCipher(key []byte) (*AESCipher, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("aes-256 key must be exactly 32 bytes, got %d", len(key))
	}
	return &AESCipher{key: key}, nil
}

// Encrypt encrypts standard data using AES-GCM and prepends the random nonce.
func (c *AESCipher) Encrypt(ctx context.Context, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, fmt.Errorf("failed to create aes block: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create gcm: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate random nonce: %w", err)
	}

	// Encrypt appends the ciphertext to the prefix nonce
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// Decrypt decodes the prefix nonce and decrypts the remaining payload.
func (c *AESCipher) Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, fmt.Errorf("failed to create aes block: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create gcm: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("malformed ciphertext payload: smaller than nonce size")
	}

	nonce := ciphertext[:nonceSize]
	actualCiphertext := ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, actualCiphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt ciphertext: integrity check failed: %w", err)
	}

	return plaintext, nil
}
