package secretsmanager_test

import (
	"context"
	"testing"

	prahariSM "prahari/shared/aws/secretsmanager"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariSM.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil Secrets Manager client, got nil")
	}
}

func TestClient_GetSecret_Uninitialized(t *testing.T) {
	client := prahariSM.NewClient(nil)
	_, err := client.GetSecret(context.Background(), "secret-id")
	if err == nil {
		t.Error("expected secret retrieval to error on uninitialized client, got nil")
	}
}
