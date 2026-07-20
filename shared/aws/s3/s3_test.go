package s3_test

import (
	"context"
	"strings"
	"testing"

	prahariS3 "prahari/shared/aws/s3"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariS3.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil S3 client, got nil")
	}
}

func TestClient_UploadDownload_Uninitialized(t *testing.T) {
	client := prahariS3.NewClient(nil)
	ctx := context.Background()

	err := client.Upload(ctx, "bucket", "key", strings.NewReader("data"))
	if err == nil {
		t.Error("expected upload to error on uninitialized client, got nil")
	}

	_, err = client.Download(ctx, "bucket", "key")
	if err == nil {
		t.Error("expected download to error on uninitialized client, got nil")
	}
}
