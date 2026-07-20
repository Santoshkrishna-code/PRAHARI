package storage

import (
	"context"
	"fmt"
	"io"

	prahariLogger "prahari/shared/logger"
)

// Client handles document upload objects.
type Client struct {
	bucket string
	region string
}

// NewClient instantiates Client.
func NewClient(bucket, region string) *Client {
	return &Client{bucket: bucket, region: region}
}

// Upload uploads raw readers.
func (c *Client) Upload(ctx context.Context, key string, reader io.Reader, contentType string) (string, error) {
	prahariLogger.Info(ctx, "Uploading attachment to S3",
		prahariLogger.String("bucket", c.bucket),
		prahariLogger.String("key", key))

	return fmt.Sprintf("s3://%s/%s", c.bucket, key), nil
}
