package kms_test

import (
	"context"
	"testing"

	prahariKMS "prahari/shared/aws/kms"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariKMS.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil KMS client, got nil")
	}
}

func TestClient_GenerateDataKey_Uninitialized(t *testing.T) {
	client := prahariKMS.NewClient(nil)
	_, _, err := client.GenerateDataKey(context.Background(), "alias/key")
	if err == nil {
		t.Error("expected key generation to error on uninitialized client, got nil")
	}
}
