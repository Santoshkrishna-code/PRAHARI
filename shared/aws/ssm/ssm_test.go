package ssm_test

import (
	"context"
	"testing"

	prahariSSM "prahari/shared/aws/ssm"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariSSM.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil SSM client, got nil")
	}
}

func TestClient_GetParameter_Uninitialized(t *testing.T) {
	client := prahariSSM.NewClient(nil)
	_, err := client.GetParameter(context.Background(), "/prahari/config/value")
	if err == nil {
		t.Error("expected parameter retrieval to error on uninitialized client, got nil")
	}
}
