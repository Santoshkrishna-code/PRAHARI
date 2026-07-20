package mocks_test

import (
	"context"
	"errors"
	"testing"

	"prahari/shared/testing/mocks"
)

func TestMocksBehavior(t *testing.T) {
	ctx := context.Background()

	// 1. S3 Mock
	s3Mock := &mocks.MockS3Client{
		UploadFunc: func(ctx context.Context, bucket, key string, data []byte) error {
			return errors.New("s3 upload error")
		},
	}
	err := s3Mock.UploadObject(ctx, "bucket", "key", nil)
	if err == nil || err.Error() != "s3 upload error" {
		t.Errorf("expected upload error, got: %v", err)
	}

	// 2. SQS Mock
	sqsMock := &mocks.MockSQSClient{
		SendFunc: func(ctx context.Context, queueURL, body string) (string, error) {
			return "msg-id-123", nil
		},
	}
	id, err := sqsMock.SendMessage(ctx, "queue", "body")
	if err != nil || id != "msg-id-123" {
		t.Errorf("expected msg-id-123, got: %v (id: %s)", err, id)
	}

	// 3. Redis Mock
	redisMock := &mocks.MockRedisClient{
		GetFunc: func(ctx context.Context, key string) (string, error) {
			return "cached-value", nil
		},
	}
	val, err := redisMock.Get(ctx, "key")
	if err != nil || val != "cached-value" {
		t.Errorf("expected cached-value, got: %v (val: %s)", err, val)
	}

	// 4. Kafka Mock
	kafkaMock := &mocks.MockKafkaPublisher{
		PublishFunc: func(ctx context.Context, topic, key string, payload []byte) error {
			return errors.New("kafka connection failure")
		},
	}
	err = kafkaMock.Publish(ctx, "topic", "key", nil)
	if err == nil || err.Error() != "kafka connection failure" {
		t.Errorf("expected kafka error, got: %v", err)
	}
}
