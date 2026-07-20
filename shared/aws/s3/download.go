package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Download retrieves a file stream from the specified S3 bucket and key.
// Caller is responsible for closing the returned stream.
func (c *Client) Download(ctx context.Context, bucket, key string) (io.ReadCloser, error) {
	if c.s3Client == nil {
		return nil, fmt.Errorf("s3 client is uninitialized")
	}

	output, err := c.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download object from s3://%s/%s: %w", bucket, key, err)
	}

	return output.Body, nil
}
