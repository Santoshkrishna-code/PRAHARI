package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3API declares interface for mockable S3 transactions.
type S3API interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

// Client wraps the S3 SDK client and presign capabilities.
type Client struct {
	s3Client      *s3.Client
	presignClient *s3.PresignClient
}

// NewClient constructs an S3 client wrapper.
func NewClient(s3Client *s3.Client) *Client {
	return &Client{
		s3Client:      s3Client,
		presignClient: s3.NewPresignClient(s3Client),
	}
}

// Ping implements aws.HealthChecker checking connectivity to S3 by listings buckets.
func (c *Client) Ping(ctx context.Context) error {
	if c.s3Client == nil {
		return fmt.Errorf("s3 client is uninitialized")
	}
	_, err := c.s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	return err
}

// PresignPutObject generates a presigned URL to upload objects to S3.
func (c *Client) PresignPutObject(ctx context.Context, bucket, key string, lifetime time.Duration) (string, error) {
	if c.presignClient == nil {
		return "", fmt.Errorf("presign client is uninitialized")
	}

	req, err := c.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, s3.WithPresignExpires(lifetime))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned PUT URL: %w", err)
	}

	return req.URL, nil
}

// PresignGetObject generates a presigned URL to retrieve files from S3.
func (c *Client) PresignGetObject(ctx context.Context, bucket, key string, lifetime time.Duration) (string, error) {
	if c.presignClient == nil {
		return "", fmt.Errorf("presign client is uninitialized")
	}

	req, err := c.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, s3.WithPresignExpires(lifetime))
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned GET URL: %w", err)
	}

	return req.URL, nil
}
