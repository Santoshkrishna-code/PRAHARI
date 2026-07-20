package encryption_test

import (
	"context"
	"crypto/rand"
	"errors"
	"testing"

	"prahari/shared/security/encryption"
)

// MockKMS implements encryption.KMSAPI for envelope tests.
type MockKMS struct {
	EncryptFunc func(ctx context.Context, keyID string, plaintextKey []byte) ([]byte, error)
	DecryptFunc func(ctx context.Context, encryptedKey []byte) ([]byte, error)
}

func (m *MockKMS) EncryptKey(ctx context.Context, keyID string, plaintextKey []byte) ([]byte, error) {
	return m.EncryptFunc(ctx, keyID, plaintextKey)
}

func (m *MockKMS) DecryptKey(ctx context.Context, encryptedKey []byte) ([]byte, error) {
	return m.DecryptFunc(ctx, encryptedKey)
}

func TestAESCipher_EncryptDecrypt(t *testing.T) {
	key := make([]byte, 32)
	_, _ = rand.Read(key)

	cipher, err := encryption.NewAESCipher(key)
	if err != nil {
		t.Fatalf("failed to construct AES cipher: %v", err)
	}

	plaintext := []byte("prahari-industrial-safety-payload")
	ctx := context.Background()

	ciphertext, err := cipher.Encrypt(ctx, plaintext)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	decrypted, err := cipher.Decrypt(ctx, ciphertext)
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}

	if string(decrypted) != string(plaintext) {
		t.Errorf("expected '%s', got '%s'", string(plaintext), string(decrypted))
	}
}

func TestRSACipher_PEMAndEncryptDecrypt(t *testing.T) {
	// Generate keys
	privKey, pubKey, err := encryption.GenerateRSAKeys(2048)
	if err != nil {
		t.Fatalf("failed to generate RSA keys: %v", err)
	}

	// PEM serialization tests
	privPEM := encryption.ExportPrivateKeyPEM(privKey)
	pubPEM, err := encryption.ExportPublicKeyPEM(pubKey)
	if err != nil {
		t.Fatalf("failed to export public PEM: %v", err)
	}

	// PEM parsing tests
	parsedPriv, err := encryption.ParsePrivateKeyPEM(privPEM)
	if err != nil {
		t.Fatalf("failed to parse private PEM: %v", err)
	}
	parsedPub, err := encryption.ParsePublicKeyPEM(pubPEM)
	if err != nil {
		t.Fatalf("failed to parse public PEM: %v", err)
	}

	cipher := encryption.NewRSACipher(parsedPriv, parsedPub)
	plaintext := []byte("rsa-test-message")
	ctx := context.Background()

	ciphertext, err := cipher.Encrypt(ctx, plaintext)
	if err != nil {
		t.Fatalf("RSA encryption failed: %v", err)
	}

	decrypted, err := cipher.Decrypt(ctx, ciphertext)
	if err != nil {
		t.Fatalf("RSA decryption failed: %v", err)
	}

	if string(decrypted) != string(plaintext) {
		t.Errorf("expected '%s', got '%s'", string(plaintext), string(decrypted))
	}
}

func TestEnvelopeCipher_EncryptDecrypt(t *testing.T) {
	mockKMS := &MockKMS{
		// Mock KMS Encrypt wraps key in fake prefix
		EncryptFunc: func(ctx context.Context, keyID string, plaintextKey []byte) ([]byte, error) {
			res := append([]byte("kms-wrapped-"), plaintextKey...)
			return res, nil
		},
		// Mock KMS Decrypt strips fake prefix
		DecryptFunc: func(ctx context.Context, encryptedKey []byte) ([]byte, error) {
			prefix := []byte("kms-wrapped-")
			if len(encryptedKey) < len(prefix) {
				return nil, errors.New("malformed key")
			}
			return encryptedKey[len(prefix):], nil
		},
	}

	envelope := encryption.NewEnvelopeCipher(mockKMS, "alias/prahari-master-key")
	plaintext := []byte("highly-confidential-sensor-telemetry")
	ctx := context.Background()

	ciphertext, err := envelope.Encrypt(ctx, plaintext)
	if err != nil {
		t.Fatalf("envelope encryption failed: %v", err)
	}

	decrypted, err := envelope.Decrypt(ctx, ciphertext)
	if err != nil {
		t.Fatalf("envelope decryption failed: %v", err)
	}

	if string(decrypted) != string(plaintext) {
		t.Errorf("expected '%s', got '%s'", string(plaintext), string(decrypted))
	}
}
