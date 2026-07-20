package storage

import (
	"context"
	"fmt"

	prahariS3 "prahari/shared/aws/s3"
)

type Storage struct {
	client *prahariS3.Client
	bucket string
}

func NewStorage(client *prahariS3.Client, bucket string) *Storage {
	return &Storage{client: client, bucket: bucket}
}

func (s *Storage) UploadAttachment(ctx context.Context, key string, data []byte) (string, error) {
	if s.client == nil {
		return fmt.Sprintf("https://s3.amazonaws.com/%s/%s", s.bucket, key), nil
	}
	url, err := s.client.PresignPutObject(ctx, s.bucket, key, 0)
	if err != nil {
		return "", fmt.Errorf("failed to upload attachment: %w", err)
	}
	return url, nil
}
