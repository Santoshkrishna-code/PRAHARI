package mocks

import (
	"context"
)

// MockS3Client exposes hook parameters to override S3 transactions.
type MockS3Client struct {
	UploadFunc   func(ctx context.Context, bucket, key string, data []byte) error
	DownloadFunc func(ctx context.Context, bucket, key string) ([]byte, error)
	DeleteFunc   func(ctx context.Context, bucket, key string) error
}

// UploadObject delegates transaction to UploadFunc.
func (m *MockS3Client) UploadObject(ctx context.Context, bucket, key string, data []byte) error {
	if m.UploadFunc != nil {
		return m.UploadFunc(ctx, bucket, key, data)
	}
	return nil
}

// DownloadObject delegates transaction to DownloadFunc.
func (m *MockS3Client) DownloadObject(ctx context.Context, bucket, key string) ([]byte, error) {
	if m.DownloadFunc != nil {
		return m.DownloadFunc(ctx, bucket, key)
	}
	return nil, nil
}

// DeleteObject delegates transaction to DeleteFunc.
func (m *MockS3Client) DeleteObject(ctx context.Context, bucket, key string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, bucket, key)
	}
	return nil
}
