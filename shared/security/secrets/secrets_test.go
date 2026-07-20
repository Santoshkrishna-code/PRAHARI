package secrets_test

import (
	"context"
	"testing"

	"prahari/shared/security/secrets"
)

func TestKMSClient_Uninitialized(t *testing.T) {
	// Constructing KMSClient with nil sdk client should return errors on execution
	kmsClient := secrets.NewKMSClient(nil)

	ctx := context.Background()
	_, err := kmsClient.EncryptKey(ctx, "alias/key", []byte("raw-key"))
	if err == nil {
		t.Error("expected encryption error on nil AWS SDK client, got nil")
	}

	_, err = kmsClient.DecryptKey(ctx, []byte("encrypted-key"))
	if err == nil {
		t.Error("expected decryption error on nil AWS SDK client, got nil")
	}
}
