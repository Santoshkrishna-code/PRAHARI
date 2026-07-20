package encryption

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

// KMSAPI defines the wrapper capabilities needed to encrypt and decrypt data keys via AWS KMS.
type KMSAPI interface {
	EncryptKey(ctx context.Context, keyID string, plaintextKey []byte) ([]byte, error)
	DecryptKey(ctx context.Context, encryptedKey []byte) ([]byte, error)
}

// EnvelopeCipher orchestrates envelope encryption using KMS and local AES ciphers.
type EnvelopeCipher struct {
	kms   KMSAPI
	keyID string
}

// NewEnvelopeCipher constructs the Envelope cipher wrapper.
func NewEnvelopeCipher(kms KMSAPI, keyID string) *EnvelopeCipher {
	return &EnvelopeCipher{
		kms:   kms,
		keyID: keyID,
	}
}

// Encrypt encrypts standard data by generating a random data key and sealing it with KMS.
func (c *EnvelopeCipher) Encrypt(ctx context.Context, plaintext []byte) ([]byte, error) {
	// 1. Generate local 32-byte data key
	dataKey := make([]byte, 32)
	if _, err := rand.Read(dataKey); err != nil {
		return nil, fmt.Errorf("failed to generate raw data key: %w", err)
	}

	// 2. Encrypt plaintext locally via AES-GCM using the data key
	localAES, err := NewAESCipher(dataKey)
	if err != nil {
		return nil, err
	}
	ciphertext, err := localAES.Encrypt(ctx, plaintext)
	if err != nil {
		return nil, fmt.Errorf("local envelope AES encrypt failed: %w", err)
	}

	// 3. Encrypt the data key using KMS
	encryptedKey, err := c.kms.EncryptKey(ctx, c.keyID, dataKey)
	if err != nil {
		return nil, fmt.Errorf("kms key seal failed: %w", err)
	}

	// 4. Build output envelope payload: [keyLen (4 bytes)] + [encryptedKey] + [ciphertext]
	keyLen := uint32(len(encryptedKey))
	result := make([]byte, 4+keyLen+uint32(len(ciphertext)))
	
	binary.BigEndian.PutUint32(result[0:4], keyLen)
	copy(result[4:4+keyLen], encryptedKey)
	copy(result[4+keyLen:], ciphertext)

	return result, nil
}

// Decrypt decodes the envelope key size, pulls the cipher payload, unseals the key via KMS, and decrypts the data.
func (c *EnvelopeCipher) Decrypt(ctx context.Context, envelopeCiphertext []byte) ([]byte, error) {
	if len(envelopeCiphertext) < 4 {
		return nil, fmt.Errorf("malformed envelope payload: smaller than size prefix")
	}

	// 1. Parse encrypted data key length
	keyLen := binary.BigEndian.Uint32(envelopeCiphertext[0:4])
	if uint32(len(envelopeCiphertext)) < 4+keyLen {
		return nil, fmt.Errorf("malformed envelope payload: key bounds mismatch")
	}

	encryptedKey := envelopeCiphertext[4 : 4+keyLen]
	ciphertext := envelopeCiphertext[4+keyLen:]

	// 2. Decrypt data key using KMS
	dataKey, err := c.kms.DecryptKey(ctx, encryptedKey)
	if err != nil {
		return nil, fmt.Errorf("kms key unseal failed: %w", err)
	}

	if len(dataKey) != 32 {
		return nil, fmt.Errorf("kms unseal returned invalid key length: %d", len(dataKey))
	}

	// 3. Decrypt ciphertext using local AES cipher
	localAES, err := NewAESCipher(dataKey)
	if err != nil {
		return nil, err
	}

	return localAES.Decrypt(ctx, ciphertext)
}
