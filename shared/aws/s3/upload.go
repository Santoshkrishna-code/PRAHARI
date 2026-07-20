package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Upload puts an object input stream into the target S3 bucket.
func (c *Client) Upload(ctx context.Context, bucket, key string, body io.Reader) error {
	if c.s3Client == nil {
		return fmt.Errorf("s3 client is uninitialized")
	}

	// Seek size if possible, otherwise pass reader
	var seekable io.ReadSeeker
	if rs, ok := body.(io.ReadSeeker); ok {
		seekable = rs
	} else {
		// Buffer in memory to allow seeking if SDK requires it,
		// or pass directly if simple put handles it
		tempBytes, err := io.ReadAll(body)
		if err != nil {
			return fmt.Errorf("failed to read upload stream: %w", err)
		}
		// Wrap in seekable reader
		seekable = io.NewSectionReader(bytesReader{tempBytes}, 0, int64(len(tempBytes)))
	}

	_, err := c.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   seekable,
	})
	if err != nil {
		return fmt.Errorf("failed to upload object to s3://%s/%s: %w", bucket, key, err)
	}

	return nil
}

type bytesReader struct {
	b []byte
}

func (r bytesReader) ReadAt(b []byte, off int64) (n int, err error) {
	if off >= int64(len(r.b)) {
		return 0, io.EOF
	}
	n = copy(b, r.b[off:])
	if n < len(b) {
		err = io.EOF
	}
	return
}
