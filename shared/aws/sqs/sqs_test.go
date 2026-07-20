package sqs_test

import (
	"context"
	"testing"

	prahariSQS "prahari/shared/aws/sqs"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariSQS.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil SQS client, got nil")
	}
}

func TestClient_SendMessage_Uninitialized(t *testing.T) {
	client := prahariSQS.NewClient(nil)
	ctx := context.Background()

	_, err := client.SendMessage(ctx, "https://sqs.us-east-1.amazonaws.com/123/Queue", "body")
	if err == nil {
		t.Error("expected message send to error on uninitialized client, got nil")
	}
}

func TestClient_Consumer_Uninitialized(t *testing.T) {
	client := prahariSQS.NewClient(nil)
	ctx := context.Background()

	err := client.StartConsumer(ctx, "https://sqs.us-east-1.amazonaws.com/123/Queue", func(ctx context.Context, body string) error {
		return nil
	})
	if err == nil {
		t.Error("expected consumer loop to error out on uninitialized client, got nil")
	}
}
