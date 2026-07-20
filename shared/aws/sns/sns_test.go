package sns_test

import (
	"context"
	"testing"

	prahariSNS "prahari/shared/aws/sns"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariSNS.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil SNS client, got nil")
	}
}

func TestClient_Publish_Uninitialized(t *testing.T) {
	client := prahariSNS.NewClient(nil)
	ctx := context.Background()

	_, err := client.PublishMessage(ctx, "arn:aws:sns:us-east-1:123456789012:Topic", "message")
	if err == nil {
		t.Error("expected topic publish to error on uninitialized client, got nil")
	}

	_, err = client.PublishSMS(ctx, "+15555555555", "message")
	if err == nil {
		t.Error("expected SMS dispatch to error on uninitialized client, got nil")
	}
}
