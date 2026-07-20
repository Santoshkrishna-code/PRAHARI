package ses_test

import (
	"context"
	"testing"

	prahariSES "prahari/shared/aws/ses"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariSES.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil SES client, got nil")
	}
}

func TestClient_SendEmail_Uninitialized(t *testing.T) {
	client := prahariSES.NewClient(nil)
	_, err := client.SendEmail(context.Background(), "from@prahari.com", "to@prahari.com", "subject", "body")
	if err == nil {
		t.Error("expected email dispatch to error on uninitialized client, got nil")
	}
}
