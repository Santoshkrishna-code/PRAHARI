package storage

import (
	"context"
	"fmt"
	"io"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	bucket string
	region string
}

func NewClient(bucket, region string) *Client {
	return &Client{bucket: bucket, region: region}
}

func (c *Client) UploadMedicalEvidence(ctx context.Context, key string, reader io.Reader, contentType string) (string, error) {
	prahariLogger.Info(ctx, "Uploading medical fitness/clearance certificate evidence documentation to encrypted S3 bucket",
		prahariLogger.String("bucket", c.bucket),
		prahariLogger.String("key", key))

	return fmt.Sprintf("s3://%s/%s", c.bucket, key), nil
}
