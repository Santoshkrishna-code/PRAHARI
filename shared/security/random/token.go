package random

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
)

// GenerateHexToken creates a cryptographically secure hex-encoded string of the given byte size.
func GenerateHexToken(size int) (string, error) {
	b, err := Bytes(size)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// GenerateBase64Token creates a cryptographically secure base64-encoded string.
func GenerateBase64Token(size int) (string, error) {
	b, err := Bytes(size)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// GenerateAlphanumericToken builds a high-entropy string using standard alphanumerics characters.
func GenerateAlphanumericToken(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("failed to generate secure alphanumeric index: %w", err)
		}
		result[i] = charset[num.Int64()]
	}
	
	return string(result), nil
}
