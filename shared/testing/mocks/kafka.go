package mocks

import (
	"context"
)

// MockKafkaPublisher exposes hook parameters to override message dispatches.
type MockKafkaPublisher struct {
	PublishFunc      func(ctx context.Context, topic string, key string, payload []byte) error
	PublishBatchFunc func(ctx context.Context, topic string, messages map[string][]byte) error
}

// Publish delegates transaction to PublishFunc.
func (m *MockKafkaPublisher) Publish(ctx context.Context, topic string, key string, payload []byte) error {
	if m.PublishFunc != nil {
		return m.PublishFunc(ctx, topic, key, payload)
	}
	return nil
}

// PublishBatch delegates transaction to PublishBatchFunc.
func (m *MockKafkaPublisher) PublishBatch(ctx context.Context, topic string, messages map[string][]byte) error {
	if m.PublishBatchFunc != nil {
		return m.PublishBatchFunc(ctx, topic, messages)
	}
	return nil
}
