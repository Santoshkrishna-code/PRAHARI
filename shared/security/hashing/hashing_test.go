package hashing_test

import (
	"context"
	"strings"
	"testing"

	"prahari/shared/security/hashing"
)

func TestBcryptHasher_HashVerify(t *testing.T) {
	hasher := hashing.NewBcryptHasher(10)
	ctx := context.Background()

	password := "SecurePass1!"
	hash, err := hasher.Hash(ctx, password)
	if err != nil {
		t.Fatalf("failed to generate bcrypt hash: %v", err)
	}

	// Verify that bcrypt output has standard format
	if !strings.HasPrefix(hash, "$2a$10$") {
		t.Errorf("expected bcrypt prefix, got '%s'", hash)
	}

	match, err := hasher.Verify(ctx, password, hash)
	if err != nil {
		t.Fatalf("verification execution error: %v", err)
	}
	if !match {
		t.Error("expected matching passwords to return true")
	}

	mismatch, _ := hasher.Verify(ctx, "WrongPassword1!", hash)
	if mismatch {
		t.Error("expected mismatched passwords to return false")
	}
}

func TestArgon2Hasher_HashVerify(t *testing.T) {
	hasher := hashing.NewArgon2Hasher(16384, 2, 2) // Small limits for testing speed
	ctx := context.Background()

	password := "ArgonSecure2!"
	hash, err := hasher.Hash(ctx, password)
	if err != nil {
		t.Fatalf("failed to generate Argon2id hash: %v", err)
	}

	if !strings.HasPrefix(hash, "$argon2id$") {
		t.Errorf("expected Argon2id prefix, got '%s'", hash)
	}

	match, err := hasher.Verify(ctx, password, hash)
	if err != nil {
		t.Fatalf("verification execution error: %v", err)
	}
	if !match {
		t.Error("expected matching passwords to return true")
	}

	mismatch, _ := hasher.Verify(ctx, "WrongPassword1!", hash)
	if mismatch {
		t.Error("expected mismatched passwords to return false")
	}
}

func TestVerifyPasswordStrength(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"Pass1!", false},       // Too short
		{"passwords1!", false},  // No uppercase
		{"PASSWORDS1!", false},  // No lowercase
		{"Passwordss!", false},  // No number
		{"Passwords11", false},  // No special char
		{"Passwords1!", true},   // Valid
	}

	for _, tt := range tests {
		err := hashing.VerifyPasswordStrength(tt.password)
		valid := err == nil
		if valid != tt.expected {
			t.Errorf("password '%s': expected valid=%t, got valid=%t (error: %v)", tt.password, tt.expected, valid, err)
		}
	}
}

func TestDetectBcryptUpgrade(t *testing.T) {
	// Bcrypt hash with cost 10
	lowCostHash := "$2a$10$vI8aWBnW3fID.i10Bjsye.13vNd5r4F2n.yKzC0T23Z0h0o22q12q"
	
	if !hashing.DetectBcryptUpgrade(lowCostHash, 12) {
		t.Error("expected cost upgrade detection to be true for cost 10 < target 12")
	}

	if hashing.DetectBcryptUpgrade(lowCostHash, 10) {
		t.Error("expected cost upgrade detection to be false for cost 10 == target 10")
	}
}
