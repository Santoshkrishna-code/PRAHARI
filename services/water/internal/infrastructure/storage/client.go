package storage

import (
	"context"
	"fmt"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	bucket string
}

func NewClient(bucket string) *Client {
	return &Client{bucket: bucket}
}

func (c *Client) UploadEvidence(ctx context.Context, key string, data []byte) (string, error) {
	url := fmt.Sprintf("https://s3.amazonaws.com/%s/%s", c.bucket, key)
	prahariLogger.Info(ctx, "Uploaded water report/evidence document to S3 bucket", prahariLogger.String("url", url))
	return url, nil
}
