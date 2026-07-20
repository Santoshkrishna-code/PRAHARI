package eventbridge_test

import (
	"context"
	"testing"

	prahariEB "prahari/shared/aws/eventbridge"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariEB.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil EventBridge client, got nil")
	}
}

func TestClient_PublishEvent_Uninitialized(t *testing.T) {
	client := prahariEB.NewClient(nil)
	_, err := client.PublishEvent(context.Background(), "bus", "source", "type", "{}")
	if err == nil {
		t.Error("expected event publish to error on uninitialized client, got nil")
	}
}
