package mocks

import (
	"context"
)

// MockSQSClient exposes hook parameters to override SQS transactions.
type MockSQSClient struct {
	SendFunc    func(ctx context.Context, queueURL, body string) (string, error)
	ReceiveFunc func(ctx context.Context, queueURL string, maxMessages int) ([]string, error)
	DeleteFunc  func(ctx context.Context, queueURL, receiptHandle string) error
}

// SendMessage delegates transaction to SendFunc.
func (m *MockSQSClient) SendMessage(ctx context.Context, queueURL, body string) (string, error) {
	if m.SendFunc != nil {
		return m.SendFunc(ctx, queueURL, body)
	}
	return "mock-message-id", nil
}

// ReceiveMessages delegates transaction to ReceiveFunc.
func (m *MockSQSClient) ReceiveMessages(ctx context.Context, queueURL string, maxMessages int) ([]string, error) {
	if m.ReceiveFunc != nil {
		return m.ReceiveFunc(ctx, queueURL, maxMessages)
	}
	return nil, nil
}

// DeleteMessage delegates transaction to DeleteFunc.
func (m *MockSQSClient) DeleteMessage(ctx context.Context, queueURL, receiptHandle string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, queueURL, receiptHandle)
	}
	return nil
}
