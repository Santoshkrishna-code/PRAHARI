package encryption

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// RSACipher implements security.Cipher using RSA-OAEP padding.
type RSACipher struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	label      []byte
}

// NewRSACipher instantiates the RSA cipher wrapper.
func NewRSACipher(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *RSACipher {
	return &RSACipher{
		privateKey: privateKey,
		publicKey:  publicKey,
		label:      []byte("prahari-rsa-label"),
	}
}

// Encrypt encrypts data using the public key and SHA-256 hashing.
func (c *RSACipher) Encrypt(ctx context.Context, plaintext []byte) ([]byte, error) {
	if c.publicKey == nil {
		return nil, fmt.Errorf("public key is uninitialized for encryption")
	}

	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		c.publicKey,
		plaintext,
		c.label,
	)
	if err != nil {
		return nil, fmt.Errorf("rsa encrypt failed: %w", err)
	}

	return ciphertext, nil
}

// Decrypt decrypts data using the private key.
func (c *RSACipher) Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	if c.privateKey == nil {
		return nil, fmt.Errorf("private key is uninitialized for decryption")
	}

	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		c.privateKey,
		ciphertext,
		c.label,
	)
	if err != nil {
		return nil, fmt.Errorf("rsa decrypt failed: %w", err)
	}

	return plaintext, nil
}
