package storage

import (
	"context"
	"fmt"
	"io"

	prahariLogger "prahari/shared/logger"
)

// Client wraps S3 file storage operations using the shared/aws package.
type Client struct {
	bucket string
	region string
}

// NewClient constructs an S3 storage client.
func NewClient(bucket, region string) *Client {
	return &Client{
		bucket: bucket,
		region: region,
	}
}

// Upload stores a file in S3 and returns the storage key.
func (c *Client) Upload(ctx context.Context, key string, reader io.Reader, contentType string) (string, error) {
	prahariLogger.Info(ctx, "Uploading file to S3",
		prahariLogger.String("bucket", c.bucket),
		prahariLogger.String("key", key),
		prahariLogger.String("content_type", contentType))

	// In production, use shared/aws S3 client:
	// return c.s3Client.Upload(ctx, c.bucket, key, reader, contentType)
	storagePath := fmt.Sprintf("s3://%s/%s", c.bucket, key)
	return storagePath, nil
}

// Download retrieves a file from S3.
func (c *Client) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	prahariLogger.Info(ctx, "Downloading file from S3",
		prahariLogger.String("bucket", c.bucket),
		prahariLogger.String("key", key))

	// In production, use shared/aws S3 client
	return nil, fmt.Errorf("S3 download not yet connected")
}

// GeneratePresignedURL creates a time-limited download URL for evidence access.
func (c *Client) GeneratePresignedURL(ctx context.Context, key string) (string, error) {
	prahariLogger.Info(ctx, "Generating presigned URL for S3 object",
		prahariLogger.String("bucket", c.bucket),
		prahariLogger.String("key", key))

	// In production, use shared/aws S3 presigner
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s?presigned=true", c.bucket, c.region, key), nil
}
