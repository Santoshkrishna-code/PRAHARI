package random_test

import (
	"regexp"
	"strings"
	"testing"

	"prahari/shared/security/random"
)

func TestBytes(t *testing.T) {
	b, err := random.Bytes(16)
	if err != nil {
		t.Fatalf("failed to generate random bytes: %v", err)
	}
	if len(b) != 16 {
		t.Errorf("expected 16 bytes, got %d", len(b))
	}
}

func TestGenerateHexToken(t *testing.T) {
	tok, err := random.GenerateHexToken(16)
	if err != nil {
		t.Fatalf("failed to generate hex token: %v", err)
	}
	if len(tok) != 32 { // 16 bytes = 32 hex chars
		t.Errorf("expected 32 characters, got %d", len(tok))
	}
}

func TestGenerateBase64Token(t *testing.T) {
	tok, err := random.GenerateBase64Token(32)
	if err != nil {
		t.Fatalf("failed to generate base64 token: %v", err)
	}
	// Verify it contains standard base64 raw URL characters
	if strings.Contains(tok, "+") || strings.Contains(tok, "/") || strings.Contains(tok, "=") {
		t.Errorf("token '%s' contains non-URL-safe characters", tok)
	}
}

func TestGenerateAlphanumericToken(t *testing.T) {
	tok, err := random.GenerateAlphanumericToken(24)
	if err != nil {
		t.Fatalf("failed to generate alphanumeric token: %v", err)
	}

	if len(tok) != 24 {
		t.Errorf("expected 24 characters, got %d", len(tok))
	}

	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", tok)
	if !match {
		t.Errorf("token '%s' contains invalid alphanumeric characters", tok)
	}
}

func TestGenerateUUIDv4(t *testing.T) {
	uuid, err := random.GenerateUUIDv4()
	if err != nil {
		t.Fatalf("failed to generate UUIDv4: %v", err)
	}

	// UUIDv4 format check: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(uuid) {
		t.Errorf("generated string '%s' is not a valid UUIDv4 format", uuid)
	}
}
